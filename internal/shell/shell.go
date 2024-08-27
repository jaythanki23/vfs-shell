package shell

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"vfs-shell/internal/vfs"
)

func StartShell(fs *vfs.FileSystem) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(fs.CurrentDir.Name + "/> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		err := executeCommand(fs, input)
		if err != nil {
			log.Print("error: ", err)
		}
	}
}

func executeCommand(fs *vfs.FileSystem, command string) error {
	args := strings.Fields(command)
	if len(args) == 0 {
		return errors.New("invalid command")
	}

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("cd: missing argument")
		} else {
			err := fs.ChangeDir(args[1])
			if err != nil {
				return err
			}
		}

	case "mkdir":
		if len(args) < 2 {
			return errors.New("mkdir: missing argument")
		} else {
			err := fs.CreateDir(args[1])
			if err != nil {
				return err
			}
		}

	case "touch":
		if len(args) < 2 {
			return errors.New("touch: missing argument")
		} else {
			err := fs.CreateFile(args[1], []byte("Hello Jay"))
			if err != nil {
				return err
			}
		}

	case "ls":
		return listDirectoryContents(fs)

	default:
		return errors.New("Unknown command")
	}

	return nil
}

func listDirectoryContents(fs *vfs.FileSystem) error {
	if !fs.CurrentDir.IsDir {
		return errors.New("current node is not an directory")
	}

	for _, file := range fs.CurrentDir.Files {
		if file.IsDir {
			fmt.Println(file.Name + "/")
		} else {
			fmt.Println(file.Name)
		}
	}

	return nil
}
