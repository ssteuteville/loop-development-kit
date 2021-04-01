package client

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-plugin"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// LoopClient is used by the controller plugin host to facilitate host initiated communication with controller plugins
type LoopClient struct {
	Broker *plugin.GRPCBroker
	Client proto.LoopClient
	S      *grpc.Server
}

// LoopStart is called by the host when the plugin is started to provide access to the host process
func (m *LoopClient) LoopStart(ctx context.Context, host Sidekick, session *server.Session) error {
	// setup service servers
	clipboardHostServer := &server.ClipboardServer{Impl: host.Clipboard()}
	cursorHostServer := &server.CursorServer{Impl: host.Cursor()}
	filesystemHostServer := &server.FilesystemServer{Impl: host.Filesystem()}
	keyboardHostServer := &server.KeyboardServer{Impl: host.Keyboard()}
	networkHostServer := &server.NetworkServer{Impl: host.Network()}
	processHostServer := &server.ProcessServer{Impl: host.Process()}
	vaultHostServer := &server.VaultServer{Impl: host.Vault()}
	uiHostServer := &server.UIServer{Impl: host.UI()}
	whisperHostServer := &server.WhisperServer{Impl: host.Whisper()}
	windowHostServer := &server.WindowServer{Impl: host.Window()}

	brokerID := m.Broker.NextId()

	readyChan := make(chan bool)

	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		m.S = grpc.NewServer(opts...)
		proto.RegisterClipboardServer(m.S, clipboardHostServer)
		proto.RegisterCursorServer(m.S, cursorHostServer)
		proto.RegisterFilesystemServer(m.S, filesystemHostServer)
		proto.RegisterKeyboardServer(m.S, keyboardHostServer)
		proto.RegisterNetworkServer(m.S, networkHostServer)
		proto.RegisterProcessServer(m.S, processHostServer)
		proto.RegisterVaultServer(m.S, vaultHostServer)
		proto.RegisterUIServer(m.S, uiHostServer)
		proto.RegisterWhisperServer(m.S, whisperHostServer)
		proto.RegisterWindowServer(m.S, windowHostServer)
		readyChan <- true
		return m.S
	}

	go m.Broker.AcceptAndServe(brokerID, serverFunc)

	select {
	case <-ctx.Done():
		return nil
	case <-readyChan:
	}

	serviceHosts := &proto.ServiceHosts{
		HostBrokerId: brokerID,
	}
	_, err := m.Client.LoopStart(ctx, &proto.LoopStartRequest{
		ServiceHosts: serviceHosts,
		Session:      session.ToProto(),
	})
	if err != nil {
		fmt.Println("loopClient.go ERROR FOLLOWS")
		fmt.Println(err)
	}
	return err
}

// LoopStop is called by the host when the plugin is stopped
func (m *LoopClient) LoopStop(ctx context.Context) error {
	var multiErr error

	_, err := m.Client.LoopStop(ctx, &emptypb.Empty{})
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	m.S.Stop()

	return multiErr
}
