package main

import (
	"fmt"
	"strings"
)

/*
	The issue here stems from the len(split_line)

this causes a syntax error whenever the line is equal
*/
func createTokens(lines []string) (string, Error) {
	for index, line := range lines {
		split_line := strings.Split(line, "")

		if len(split_line) < 2 && split_line[0] != "\n" {
			return fmt.Sprintf("Breaking syntax rules, use <keyword> <value> instead.\nerror in line %v", index+1), SyntaxError
		}

		// fmt.Printf("Array: %v, Size: %v", split_line, len(split_line))

		if len(split_line) > 1 {
			// token should be parsed strictly!
			token := Token{Token: split_line[0], Value: split_line[1]}
			stack = append(stack, token)
		}

	}

	return "", NoError
}
