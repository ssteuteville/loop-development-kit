package ldk

import (
	"context"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-plugin"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// LoopServer is used by the controller plugin to receive host initiated communication
type LoopServer struct {
	Impl Loop

	Broker *plugin.GRPCBroker
	Conn   *grpc.ClientConn
	Logger *Logger
}

// LoopStart is called by the host when the plugin is started to provide access to the host process
func (m *LoopServer) LoopStart(_ context.Context, req *proto.LoopStartRequest) (*emptypb.Empty, error) {
	session, err := server.NewSessionFromProto(req.Session)
	if err != nil {
		return nil, err
	}

	hosts := req.ServiceHosts

	m.Conn, err = m.Broker.Dial(hosts.HostBrokerId)
	if err != nil {
		println("[ERROR] loopServer.go - conn Error")
		println("[ERROR]" + err.Error())
		return &emptypb.Empty{}, err
	}

	eli := &ExceptionLoggingInterceptor{
		InterceptedConn: m.Conn,
		Logger:          m.Logger,
	}

	sidekickClient := &SidekickClient{
		VaultService: &VaultClient{
			VaultClient:  proto.NewVaultClient(eli),
			Session: session,
		},
		WhisperService: &WhisperClient{
			WhisperClient:  proto.NewWhisperClient(eli),
			Session: session,
		},
		ClipboardService: &ClipboardClient{
			ClipboardClient:  proto.NewClipboardClient(eli),
			Session: session,
		},
		KeyboardService: &KeyboardClient{
			KeyboardClient:  proto.NewKeyboardClient(eli),
			Session: session,
		},
		ProcessService: &ProcessClient{
			ProcessClient:  proto.NewProcessClient(eli),
			Session: session,
		},
		CursorService: &CursorClient{
			CursorClient:  proto.NewCursorClient(eli),
			Session: session,
		},
		FilesystemService: &FilesystemClient{
			FilesystemClient:  proto.NewFilesystemClient(eli),
			Session: session,
		},
		UiService: &UIClient{
			UiClient:  proto.NewUIClient(eli),
			Session: session,
		},
		NetworkService: &NetworkClient{
			NetworlClient:  proto.NewNetworkClient(eli),
			Session: session,
		},
		WindowService: &WindowClient{
			WindowClient:  proto.NewWindowClient(eli),
			Session: session,
		},
	}
	println("[INFO] loopServer.go - LoopStart complete")
	return &emptypb.Empty{}, m.Impl.LoopStart(
		sidekickClient,
	)
}

// LoopStop is called by the host when the plugin is stopped
func (m *LoopServer) LoopStop(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	var multiErr error

	// stop loop
	if err := m.Impl.LoopStop(); err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	// close service connection
	if err := m.Conn.Close(); err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	return &emptypb.Empty{}, multiErr
}
