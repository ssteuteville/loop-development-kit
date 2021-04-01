package ldktest

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/client"
)

type NetworkService struct {
	HTTPRequestf func(context.Context, *client.HTTPRequest) (*client.HTTPResponse, error)
}

func (f *NetworkService) HTTPRequest(ctx context.Context, req *client.HTTPRequest) (*client.HTTPResponse, error) {
	return f.HTTPRequestf(ctx, req)
}
