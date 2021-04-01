package ldktest

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
)

type ClipboardService struct {
	Readf   func(context.Context) (string, error)
	Listenf func(context.Context, service.ClipboardListenConfiguration) error
	Writef  func(context.Context, string) error
}

func (c *ClipboardService) Read(ctx context.Context) (string, error) {
	return c.Readf(ctx)
}

func (c *ClipboardService) Listen(ctx context.Context, cb service.ClipboardListenConfiguration) error {
	return c.Listenf(ctx, cb)
}

func (c *ClipboardService) Write(ctx context.Context, s string) error {
	return c.Writef(ctx, s)
}
