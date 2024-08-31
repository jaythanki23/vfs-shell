package vfs

import (
	"errors"
	// "log"
	"strings"
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

	return &FileSystem{Root: root, CurrentDir: root}
}

func (fs *FileSystem) CreateFile(name string, content []byte) error {

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

func (fs *FileSystem) CreateDir(name string) error {

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

func (fs *FileSystem) ChangeDir(inputText string, currentPath string) (string, error) {
	if inputText == ".." {
		if fs.CurrentDir.Parent != nil {
			indexOfLastSlash := strings.LastIndex(currentPath, "/")

			if indexOfLastSlash == 0 {
				currentPath = "/"
			} else {
				currentPath = currentPath[:indexOfLastSlash]
			}

			fs.CurrentDir = fs.CurrentDir.Parent
		}

		return currentPath, nil
	}

	// determine whether inputText is a name of a directory or a path
	if strings.Contains(inputText, "/") { // could use strings.ContainsRune()?
		// it is a path
		path := strings.Split(inputText, "/")
		curr := fs.Root

		for _, dirName := range path {
			var directoryFound bool = false
			for _, dir := range curr.Files {
				if dir.IsDir && dir.Name == dirName {
					directoryFound = true
					curr = dir
				}
			}

			if !directoryFound {
				errors.New("Path invalid")
			}
		}

		fs.CurrentDir = curr
		return inputText, nil

	} else {
		// it is the name of a sub-directory
		for _, file := range fs.CurrentDir.Files {
			if file.IsDir && file.Name == inputText {
				fs.CurrentDir = file
				if len(currentPath) > 1 {
					currentPath = currentPath + "/" + fs.CurrentDir.Name
				} else {
					currentPath = currentPath + fs.CurrentDir.Name
				}
				return currentPath, nil
			}
		}

		return currentPath, errors.New("Directory not found")
	}

}

func (fs *FileSystem) ReadFile(name string) (string, error) {

	for _, file := range fs.CurrentDir.Files {
		if !file.IsDir && file.Name == name {
			return string(file.Content), nil
		}
	}

	return "", errors.New("file not found")
}

func (fs *FileSystem) GetCurrentDirectory() string {
	var path string
	current := fs.CurrentDir

	for current != nil {
		if current.Name == "/" {
			path = current.Name + path
		} else {
			path = current.Name + "/" + path
		}

		current = current.Parent
	}

	return path
}
