package main

import (
	"errors"
	"fmt"
	"os"
)

type Token struct {
	Token string
	Value string
}

var stack []Token

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
				Message(ERROR, "Invalid file type! Unable to start interpreter.",
					"Hint: Files must end in .abt")

			case FileDoesNotExistError:
				Message(ERROR, "Invalid file! Unable to start interpreter.",
					"Hint: Are you sure this file exists?")

			}
			return
		}

		var lines = loadScript(secondArg)
		var tokenErrorMessage, tokenError = createTokens(lines)

		if tokenError != NoError {
			Message(ERROR, "An error occured!", tokenErrorMessage)
			return
		}
	}
}

func getArg(index int) (string, error) {
	args := os.Args
	if len(args) > index {
		return args[index], nil
	}

	return args[0], errors.New("Not enough arguments!")
}
