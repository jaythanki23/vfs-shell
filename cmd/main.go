package main

import (
	"log"
	"vfs-shell/internal/vfs"
)

func main() {
	fs := vfs.NewFileSystem()

	if err := fs.CreateFile("Hello", []byte("Hi Jay")); err != nil {
		log.Printf("Error creating file: %v", err.Error())
	}

	// log.Println(fs.CurrentDir.Files[0])

	if err := fs.CreateDir("project"); err != nil {
		log.Printf("Error creating directory: %v", err.Error())
	}

	// log.Println(fs.CurrentDir.Files[1])

	if err := fs.ChangeDir("/project"); err != nil {
		log.Printf("Error changing directory: %v", err.Error())
	}

	log.Println("Current dir: ", fs.CurrentDir)
}
