package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var stack = make(map[string]string)

type Message struct {
	Title       string
	Description string
	Type        string // info, success, warning, error
}

func (m *Message) Write() {

}

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

		var isValid = verifyFile(secondArg)

	}
}

func loadScript(path string) {

}

func getArg(index int) (string, error) {
	args := os.Args
	if len(args) > index {
		return args[index], nil
	}

	return args[0], errors.New("Not enough arguments!")
}

func verifyFile(path string) bool {
	if strings.HasSuffix(path, ".amb") {
		_, err := os.ReadFile(path)
		if err != nil {
			return false
		}
		return true
	}

	return false
}
