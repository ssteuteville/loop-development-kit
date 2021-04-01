package ldk

import (
	"bytes"
	"context"
	"errors"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/utils"
	"io"
	"os"
	"sync"

	"github.com/golang/protobuf/ptypes"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
)

// FilesystemClient is used by the controller plugin to facilitate plugin initiated communication with the host
type FilesystemClient struct {
	FilesystemClient proto.FilesystemClient
	Session          *server.Session
}

// Dir list the contents of a directory
func (f *FilesystemClient) Dir(ctx context.Context, dir string) ([]os.FileInfo, error) {
	resp, err := f.FilesystemClient.FilesystemDir(ctx, &proto.FilesystemDirRequest{
		Directory: dir,
		Session:   f.Session.ToProto(),
	})
	if err != nil {
		return nil, err
	}

	files := resp.GetFiles()

	lfiles := make([]os.FileInfo, 0, len(files))
	for _, f := range files {
		t, err := ptypes.Timestamp(f.Updated)
		if err != nil {
			return nil, err
		}
		file := utils.NewFileInfo(f.GetName(), int(f.GetMode()), int(f.GetSize()), t, f.GetIsDir())
		lfiles = append(lfiles, &file)

	}

	return lfiles, nil
}

// ListenDir stream any updates to the contents of a directory
func (f *FilesystemClient) ListenDir(ctx context.Context, dir string, handler service.ListenDirHandler) error {
	client, err := f.FilesystemClient.FilesystemDirStream(ctx, &proto.FilesystemDirStreamRequest{
		Directory: dir,
		Session:   f.Session.ToProto(),
	})
	if err != nil {
		return err
	}

	go func() {
		for {
			resp, err := client.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				handler(utils.FileEvent{}, err)
				return
			}

			if resp.GetError() != "" {
				err = errors.New(resp.GetError())
				handler(utils.FileEvent{}, err)
				continue
			}

			file := resp.GetFile()
			t, err := ptypes.Timestamp(file.Updated)
			if err != nil {
				handler(utils.FileEvent{}, err)
				continue
			}

			fi := utils.NewFileInfo(file.GetName(), int(file.GetMode()), int(file.GetSize()), t, file.GetIsDir())
			handler(utils.FileEvent{Info: &fi, Action: utils.ProtoActionToAction(resp.GetAction())}, err)
		}
	}()

	return nil
}

// ListenFile stream any updates to a file
func (f *FilesystemClient) ListenFile(ctx context.Context, path string, handler service.ListenFileHandler) error {
	client, err := f.FilesystemClient.FilesystemFileInfoStream(ctx, &proto.FilesystemFileInfoStreamRequest{
		Path:    path,
		Session: f.Session.ToProto(),
	})
	if err != nil {
		return err
	}

	go func() {
		for {
			resp, err := client.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				handler(utils.FileEvent{}, err)
				return
			}

			if resp.GetError() != "" {
				err = errors.New(resp.GetError())
				handler(utils.FileEvent{}, err)
				continue
			}

			file := resp.GetFile()
			t, err := ptypes.Timestamp(file.Updated)
			if err != nil {
				handler(utils.FileEvent{}, err)
				continue
			}

			fi := utils.NewFileInfo(file.GetName(), int(file.GetMode()), int(file.GetSize()), t, file.GetIsDir())
			handler(utils.FileEvent{Info: &fi, Action: utils.ProtoActionToAction(resp.GetAction())}, err)
		}
	}()

	return nil
}

type GRPCFile struct {
	ctx       context.Context
	stream    proto.Filesystem_FilesystemFileStreamClient
	fileMutex sync.Mutex
	buffer    bytes.Buffer
}

func newGRPCFile(ctx context.Context, stream proto.Filesystem_FilesystemFileStreamClient) *GRPCFile {
	GRPCFile := &GRPCFile{
		ctx:    ctx,
		stream: stream,
	}

	return GRPCFile
}

func (f *GRPCFile) Sync() error {
	return nil
}

func (f *GRPCFile) makeRequest() (int, error) {
	err := f.stream.Send(&proto.FilesystemFileStreamRequest{
		RequestOneOf: &proto.FilesystemFileStreamRequest_Read_{
			Read: &proto.FilesystemFileStreamRequest_Read{},
		},
	})
	if err != nil {
		return 0, err
	}

	for {
		select {
		case <-f.ctx.Done():
			return 0, errors.New("context closed")
		default:
			resp, err := f.stream.Recv()
			if err != nil {
				return 0, err
			}
			switch resp.ResponseOneOf.(type) {
			case *proto.FilesystemFileStreamResponse_Read_:

				if resp.GetRead().GetError() == "EOF" {
					return 0, io.EOF
				}

				if resp.GetRead().GetError() == "" {
					recievedBytes := resp.GetRead().GetData()
					n := len(recievedBytes)
					f.buffer.Write(recievedBytes)
					return n, nil
				}

				return 0, errors.New(resp.GetRead().GetError())
			default:
				return 0, errors.New("unexpected response from server")
			}

		}
	}

}
func (f *GRPCFile) Read(p []byte) (int, error) {
	f.fileMutex.Lock()
	defer f.fileMutex.Unlock()
	pl := len(p)

	for {
		b := f.buffer.Bytes()
		bl := len(b)

		if bl >= pl {
			n, err := f.buffer.Read(p)
			if err != nil && err != io.EOF {
				f.buffer.Reset()
				return 0, err
			}
			return n, nil
		}

		_, err := f.makeRequest()
		if err != nil && err != io.EOF {
			f.buffer.Reset()
			return 0, err
		}

		if err == io.EOF {
			readBytes := f.buffer.Bytes()
			_, err = f.buffer.Read(p)
			if err != nil {
				f.buffer.Reset()
				return 0, err
			}

			return len(readBytes), io.EOF
		}

	}
}

func (f *GRPCFile) Write(b []byte) (int, error) {
	f.fileMutex.Lock()
	defer f.fileMutex.Unlock()

	err := f.stream.Send(&proto.FilesystemFileStreamRequest{
		RequestOneOf: &proto.FilesystemFileStreamRequest_Write_{
			Write: &proto.FilesystemFileStreamRequest_Write{
				Data: b,
			},
		},
	})
	if err != nil {
		return 0, err
	}
	for {
		select {
		case <-f.ctx.Done():
			return 0, errors.New("context closed")
		default:
			resp, err := f.stream.Recv()
			if err != nil {
				return 0, err
			}
			switch resp.ResponseOneOf.(type) {
			case *proto.FilesystemFileStreamResponse_Write_:
				if resp.GetWrite().GetError() == "" {
					return int(resp.GetWrite().GetNumOfBytes()), nil
				}
				return int(resp.GetWrite().GetNumOfBytes()), errors.New(resp.GetWrite().GetError())
			default:
				return 0, errors.New("unexpected response from server")
			}

		}
	}
}

// Close closes file
func (f *GRPCFile) Close() error {
	f.fileMutex.Lock()
	defer f.fileMutex.Unlock()
	err := f.stream.Send(&proto.FilesystemFileStreamRequest{
		RequestOneOf: &proto.FilesystemFileStreamRequest_Close_{
			Close: &proto.FilesystemFileStreamRequest_Close{},
		},
	})
	if err != nil {
		println("[ERROR] GRPCFile.Close - Received error sending Close", err.Error())
		return err
	}
	return nil
}

// Chown changes owner of file
func (f *GRPCFile) Chown(uid, gid int) error {
	f.fileMutex.Lock()
	defer f.fileMutex.Unlock()
	err := f.stream.Send(&proto.FilesystemFileStreamRequest{
		RequestOneOf: &proto.FilesystemFileStreamRequest_Chown_{
			Chown: &proto.FilesystemFileStreamRequest_Chown{
				Uid: int32(uid),
				Gid: int32(gid),
			},
		},
	})
	if err != nil {
		return err
	}
	for {
		select {
		case <-f.ctx.Done():
			return errors.New("context closed")
		default:
			resp, err := f.stream.Recv()
			if err != nil {
				return err
			}
			switch resp.ResponseOneOf.(type) {
			case *proto.FilesystemFileStreamResponse_Chown_:
				if resp.GetChown().GetError() == "" {
					return nil
				}
				return errors.New(resp.GetChown().GetError())
			default:
				return errors.New("unexpected response from server")
			}

		}
	}
}

// Chmod changes permissions of file
func (f *GRPCFile) Chmod(mode os.FileMode) error {
	f.fileMutex.Lock()
	defer f.fileMutex.Unlock()
	err := f.stream.Send(&proto.FilesystemFileStreamRequest{
		RequestOneOf: &proto.FilesystemFileStreamRequest_Chmod_{
			Chmod: &proto.FilesystemFileStreamRequest_Chmod{
				Mode: uint32(mode),
			},
		},
	})
	if err != nil {
		return err
	}
	for {
		select {
		case <-f.ctx.Done():
			return errors.New("context closed")
		default:
			resp, err := f.stream.Recv()
			if err != nil {
				return err
			}
			switch resp.ResponseOneOf.(type) {
			case *proto.FilesystemFileStreamResponse_Chmod_:
				if resp.GetChmod().GetError() == "" {
					return nil
				}
				return errors.New(resp.GetChmod().GetError())
			default:
				return errors.New("unexpected response from server")
			}

		}
	}

}

// Stat get info of file
func (f *GRPCFile) Stat() (os.FileInfo, error) {
	f.fileMutex.Lock()
	defer f.fileMutex.Unlock()
	err := f.stream.Send(&proto.FilesystemFileStreamRequest{
		RequestOneOf: &proto.FilesystemFileStreamRequest_Stat_{
			Stat: &proto.FilesystemFileStreamRequest_Stat{},
		},
	})
	if err != nil {
		return nil, err
	}
	for {
		select {
		case <-f.ctx.Done():
			return nil, errors.New("context closed")
		default:
			resp, err := f.stream.Recv()
			if err != nil {
				return nil, err
			}
			switch resp.ResponseOneOf.(type) {
			case *proto.FilesystemFileStreamResponse_Stat_:
				info := resp.GetStat().GetInfo()
				t, err := ptypes.Timestamp(info.Updated)
				if err != nil {
					return nil, err
				}
				fi := utils.NewFileInfo(info.GetName(), int(info.GetMode()), int(info.GetSize()), t, info.GetIsDir())
				if resp.GetStat().GetError() == "" {
					return &fi, nil
				}
				return &fi, errors.New(resp.GetStat().GetError())
			}
		}
	}
}

// Open something
func (f *FilesystemClient) Open(ctx context.Context, path string) (utils.File, error) {

	fileStreamClient, err := f.FilesystemClient.FilesystemFileStream(ctx)
	if err != nil {
		return nil, err
	}

	GRPCFile := newGRPCFile(ctx, fileStreamClient)

	err = fileStreamClient.Send(&proto.FilesystemFileStreamRequest{
		RequestOneOf: &proto.FilesystemFileStreamRequest_Open_{
			Open: &proto.FilesystemFileStreamRequest_Open{
				Session: f.Session.ToProto(),
				Path:    path,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return GRPCFile, nil
}

func (f *FilesystemClient) Create(ctx context.Context, path string) (utils.File, error) {

	fileStreamClient, err := f.FilesystemClient.FilesystemFileStream(ctx)
	if err != nil {
		return nil, err
	}

	GRPCFile := newGRPCFile(ctx, fileStreamClient)

	err = fileStreamClient.Send(&proto.FilesystemFileStreamRequest{
		RequestOneOf: &proto.FilesystemFileStreamRequest_Create_{
			Create: &proto.FilesystemFileStreamRequest_Create{
				Session: f.Session.ToProto(),
				Path:    path,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return GRPCFile, nil
}

// MakeDir create new directory
func (f *FilesystemClient) MakeDir(ctx context.Context, path string, perm uint32) error {
	_, err := f.FilesystemClient.FilesystemMakeDir(ctx, &proto.FilesystemMakeDirRequest{
		Session: f.Session.ToProto(),
		Path:    path,
		Perm:    perm,
	})
	if err != nil {
		return err
	}

	return nil
}

// Copy file or directory
func (f *FilesystemClient) Copy(ctx context.Context, source, dest string) error {
	_, err := f.FilesystemClient.FilesystemCopy(ctx, &proto.FilesystemCopyRequest{
		Session: f.Session.ToProto(),
		Source:  source,
		Dest:    dest,
	})
	if err != nil {
		return err
	}

	return nil
}

// Move file or directory
func (f *FilesystemClient) Move(ctx context.Context, source, dest string) error {
	_, err := f.FilesystemClient.FilesystemMove(ctx, &proto.FilesystemMoveRequest{
		Session: f.Session.ToProto(),
		Source:  source,
		Dest:    dest,
	})
	if err != nil {
		return err
	}

	return nil
}

// Remove file or directory
func (f *FilesystemClient) Remove(ctx context.Context, path string, recursive bool) error {
	_, err := f.FilesystemClient.FilesystemRemove(ctx, &proto.FilesystemRemoveRequest{
		Session:   f.Session.ToProto(),
		Path:      path,
		Recursive: recursive,
	})
	if err != nil {
		return err
	}

	return nil
}
