package main

import (
	"errors"
	"fmt"
	"os"
)

var stack = make(map[string]string)

func main() {
	firstArg, err := getArg(1)
	if err != nil {
		fmt.Println("No arguments passed!")
		return
	}

	if firstArg == "run" {
		secondArg, err := getArg(2)
		if err != nil {
			return
		}

		var isValid, errType = verifyFile(secondArg)
		if !isValid {
			switch errType {
			case FileTypeError:
				Message(ERROR, "Invalid file type!",
					"Hint: Files must end in .abt")

			case FileDoesNotExistError:
				Message(ERROR, "Invalid file!",
					"Hint: Are you sure this file exists?")

			}
			return
		}

		var lines = loadScript(secondArg)
		fmt.Println(lines)
	}
}

func getArg(index int) (string, error) {
	args := os.Args
	if len(args) > index {
		return args[index], nil
	}

	return args[0], errors.New("Not enough arguments!")
}
