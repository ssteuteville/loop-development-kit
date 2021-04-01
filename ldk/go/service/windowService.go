package service

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/utils"
)

// WindowService is an interface that defines what methods plugins can expect from the host
type WindowService interface {
	ActiveWindow(context.Context) (utils.WindowInfo, error)
	ListenActiveWindow(context.Context, ListenActiveWindowHandler) error
	State(context.Context) ([]utils.WindowInfo, error)
	ListenState(context.Context, ListenWindowStateHandler) error
}

// ListenActiveWindowHandler is the signature for a handler than handles changes to the active WindowService
type ListenActiveWindowHandler func(utils.WindowInfo, error)

// ListenWindowStateHandler is the signature for a handler than handles changes to WindowService state
type ListenWindowStateHandler func(utils.WindowEvent, error)
