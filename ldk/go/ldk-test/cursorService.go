package ldktest

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
)

type CursorService struct {
	PositionF       func(context.Context) (service.CursorPosition, error)
	ListenPositionF func(context.Context, service.ListenPositionHandler) error
}

func (c *CursorService) Position(ctx context.Context) (service.CursorPosition, error) {
	return c.PositionF(ctx)
}
func (c *CursorService) ListenPosition(ctx context.Context, handler service.ListenPositionHandler) error {
	return c.ListenPositionF(ctx, handler)
}
