package utils

import (
	"os"
	"time"
)

type FileInfo struct {
	name    string
	mode    int
	size    int
	updated time.Time
	isDir   bool
}

func (f *FileInfo) IsDir() bool {
	return f.isDir
}

func (f *FileInfo) Name() string {
	return f.name
}
func (f *FileInfo) Size() int64 {
	return int64(f.size)
}

func (f *FileInfo) Mode() os.FileMode {
	return os.FileMode(f.mode)
}

func (f *FileInfo) ModTime() time.Time {
	return f.updated
}

func (f *FileInfo) Sys() interface{} {
	return nil
}

func NewFileInfo(name string, mode, size int, updated time.Time, isDir bool) FileInfo {
	return FileInfo{
		name:    name,
		mode:    mode,
		size:    size,
		updated: updated,
		isDir:   isDir,
	}
}