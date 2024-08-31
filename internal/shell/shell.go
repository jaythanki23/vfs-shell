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

var path string = "/"

func StartShell(fs *vfs.FileSystem) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(path + "> ")

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
			var err error
			path, err = fs.ChangeDir(args[1], path)
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
		for _, file := range fs.CurrentDir.Files {
			if file.IsDir {
				fmt.Println(file.Name + "/")
			} else {
				fmt.Println(file.Name)
			}
		}

		return nil

	case "cat":
		if len(args) < 2 {
			return errors.New("cat: missing argument")
		} else {
			str, err := fs.ReadFile(args[1])
			if err != nil {
				return err
			} else {
				fmt.Println(str)
			}
		}

	case "pwd":
		fmt.Println(fs.GetCurrentDirectory())

	default:
		return errors.New("unknown command")
	}

	return nil
}
