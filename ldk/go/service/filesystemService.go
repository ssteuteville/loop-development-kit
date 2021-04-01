package service

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/utils"
	"os"
)

// FilesystemService is an interface that defines what methods plugins can expect from the host
type FilesystemService interface {
	Dir(context.Context, string) ([]os.FileInfo, error)
	ListenDir(context.Context, string, ListenDirHandler) error
	ListenFile(context.Context, string, ListenFileHandler) error
	Open(context.Context, string) (utils.File, error)
	Create(context.Context, string) (utils.File, error)
	MakeDir(context.Context, string, uint32) error
	Copy(context.Context, string, string) error
	Move(context.Context, string, string) error
	Remove(context.Context, string, bool) error
}


type ListenDirHandler func(utils.FileEvent, error)
type ListenFileHandler func(utils.FileEvent, error)

