package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Status int

const (
	SUCCESS Status = iota
	INFO
	WARNING
	ERROR
)

func Message(mtype Status, title string, description string) {
	if mtype == INFO {
		color.Blue("%v\n", title)
		fmt.Println(description)

	} else if mtype == SUCCESS {
		color.Green("%v\n", title)
		fmt.Println(description)
	} else if mtype == WARNING {
		color.Yellow("%v\n", title)
		fmt.Println(description)
	} else if mtype == ERROR {
		color.Red("%v\n", title)
		fmt.Println(description)
	}

}
