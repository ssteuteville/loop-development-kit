package client

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
)

// NetworkClient is the VaultClient used by the NetworkService
type NetworkClient struct {
	NetworlClient proto.NetworkClient
	Session       *server.Session
}

func (n *NetworkClient) HTTPRequest(ctx context.Context, req *service.HTTPRequest) (*service.HTTPResponse, error) {
	reqHeaders := make(map[string]*proto.HTTPHeader)

	for name, values := range req.Headers {
		reqHeaders[name] = &proto.HTTPHeader{
			Values: values,
		}
	}

	resp, err := n.NetworlClient.HTTPRequest(ctx, &proto.HTTPRequestMsg{
		Session: n.Session.ToProto(),
		Url:     req.URL,
		Method:  req.Method,
		Body:    req.Body,
		Headers: reqHeaders,
	})

	if err != nil {
		return nil, err
	}

	respHeaders := make(map[string][]string)

	for name, values := range resp.Headers {
		respHeaders[name] = values.Values
	}

	return &service.HTTPResponse{
		ResponseCode: int(resp.ResponseCode),
		Data:         resp.Data,
		Headers:      respHeaders,
	}, err
}
