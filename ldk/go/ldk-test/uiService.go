package ldktest

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
)

type UIService struct {
	ListenSearchbarf    func(context.Context, service.ListenSearchHandler) error
	ListenGlobalSearchf func(context.Context, service.ListenSearchHandler) error
}

func (u *UIService) ListenSearchbar(ctx context.Context, cb service.ListenSearchHandler) error {
	return u.ListenSearchbarf(ctx, cb)
}

func (u *UIService) ListenGlobalSearch(ctx context.Context, cb service.ListenSearchHandler) error {
	return u.ListenGlobalSearchf(ctx, cb)
}
