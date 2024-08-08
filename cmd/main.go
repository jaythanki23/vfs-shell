package main

import (
	"fmt"
	"log"
	"vfs-shell/internal/vfs"
)

func main() {
	fmt.Println("Hello World!!")
	fs := vfs.NewFileSystem()
	log.Println(*(fs.Root))
}
