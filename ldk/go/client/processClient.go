package client

import (
	"context"
	"errors"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
	"io"

	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
)

type ProcessClient struct {
	ProcessClient proto.ProcessClient
	Session       *server.Session
}

func convertToProcessInfo(pi *proto.ProcessInfo) service.ProcessInfo {
	return service.ProcessInfo{
		Arguments: pi.Arguments,
		Command:   pi.Command,
		PID:       int(pi.Pid),
	}
}

func convertToProcessEvent(pi *proto.ProcessStateStreamResponse) service.ProcessEvent {
	return service.ProcessEvent{
		Process: convertToProcessInfo(pi.Process),
		Action:  service.ProcessAction(pi.Action),
	}
}

func (p *ProcessClient) State(ctx context.Context) ([]service.ProcessInfo, error) {
	msg := &proto.ProcessStateRequest{Session: p.Session.ToProto()}
	resp, err := p.ProcessClient.ProcessState(ctx, msg)
	if err != nil {
		return nil, err
	}
	processes := resp.GetProcesses()
	var pis []service.ProcessInfo
	for _, pi := range processes {
		pis = append(pis, convertToProcessInfo(pi))
	}
	return pis, err
}

func (p *ProcessClient) ListenState(ctx context.Context, handler service.ListenProcessStateHandler) error {
	msg := &proto.ProcessStateStreamRequest{Session: p.Session.ToProto()}
	client, err := p.ProcessClient.ProcessStateStream(ctx, msg)
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
				handler(convertToProcessEvent(resp), err)
				return
			}
			if resp.GetError() != "" {
				err = errors.New(resp.GetError())
			}
			handler(convertToProcessEvent(resp), err)
		}
	}()

	return nil
}
