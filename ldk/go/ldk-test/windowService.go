package ldktest

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/utils"
)

type WindowService struct {
	ActiveWindowf       func(ctx context.Context) (utils.WindowInfo, error)
	ListenActiveWindowf func(ctx context.Context, handler service.ListenActiveWindowHandler) error
	Statef              func(ctx context.Context) ([]utils.WindowInfo, error)
	ListenStatef        func(ctx context.Context, handler service.ListenWindowStateHandler) error
}

func (w *WindowService) ActiveWindow(ctx context.Context) (utils.WindowInfo, error) {
	return w.ActiveWindowf(ctx)
}

func (w *WindowService) ListenActiveWindow(ctx context.Context, handler service.ListenActiveWindowHandler) error {
	return w.ListenActiveWindowf(ctx, handler)
}

func (w *WindowService) State(ctx context.Context) ([]utils.WindowInfo, error) {
	return w.Statef(ctx)
}

func (w *WindowService) ListenState(ctx context.Context, handler service.ListenWindowStateHandler) error {
	return w.ListenStatef(ctx, handler)
}
