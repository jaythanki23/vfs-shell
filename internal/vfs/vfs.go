package vfs

import (
	"errors"
	"time"
)

// node of a tree
type File struct {
	Name    string
	Content []byte
	IsDir   bool
	Mode    string
	ModTime time.Time
	Parent  *File
	Files   []*File
}

type FileSystem struct {
	Root       *File
	CurrentDir *File
}

func NewFileSystem() *FileSystem {
	root := &File{
		Name:    "/",
		IsDir:   true,
		Mode:    "drwxr-xr-x",
		ModTime: time.Now(),
		Parent:  nil,
	}

	return &FileSystem{Root: root}
}

func (fs *FileSystem) createFile(name string, content []byte) error {
	if !fs.CurrentDir.IsDir {
		return errors.New("not a directory")
	}

	newFile := &File{
		Name:    name,
		Content: content,
		IsDir:   false,
		Mode:    "-rw-r--r--",
		ModTime: time.Now(),
		Parent:  fs.CurrentDir,
	}

	fs.CurrentDir.Files = append(fs.CurrentDir.Files, newFile)
	return nil
}

func (fs *FileSystem) createDir(name string) error {
	if !fs.CurrentDir.IsDir {
		return errors.New("not a directory")
	}

	newDir := &File{
		Name:    name,
		IsDir:   true,
		Mode:    "drwxr-xr-x",
		ModTime: time.Now(),
		Parent:  fs.CurrentDir,
	}

	fs.CurrentDir.Files = append(fs.CurrentDir.Files, newDir)
	return nil
}
