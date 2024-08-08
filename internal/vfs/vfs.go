package vfs

import (
	"time"
)

// node of a tree
type File struct {
	Name    string
	Content []byte
	IsDir   bool
	Mode    string
	ModTime time.Time
	Files   []*File
}

type FileSystem struct {
	Root *File
}

func NewFileSystem() *FileSystem {
	root := &File{
		Name:    "/",
		IsDir:   true,
		Mode:    "drwxr-xr-x",
		ModTime: time.Now(),
	}

	return &FileSystem{Root: root}
}
