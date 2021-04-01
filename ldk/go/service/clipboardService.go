package service

import "context"

// ClipboardService is an interface that defines what methods are made available to the loops for interacting with the ClipboardService
type ClipboardService interface {
	Read(context.Context) (string, error)
	Listen(context.Context, ClipboardListenConfiguration) error
	Write(context.Context, string) error
}

type LoopClipboardService interface {
	Listen(context.Context, ReadListenHandler) error
}

// ReadListenHandler is the signature for a handler than handles changes to the ClipboardService text
type ReadListenHandler func(string, error)

type ClipboardListenConfiguration struct {
	Handler                 ReadListenHandler
	IncludeOliveHelpTraffic bool
}
