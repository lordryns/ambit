package main

import (
	"os"
	"strings"
)

func verifyFile(path string) (bool, Error) {
	if strings.HasSuffix(path, ".abt") {
		_, err := os.ReadFile(path)
		if err != nil {
			return false, FileDoesNotExistError
		}
		return true, NoError
	}

	return false, FileTypeError
}

func loadScript(path string) []string {
	byteContent, _ := os.ReadFile(path)
	return strings.Split(string(byteContent), "\n")
}
