package utils

import "os"

type FileEvent struct {
	Info   os.FileInfo
	Action FileAction
}