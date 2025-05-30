package main

type Error int

const (
	NoError Error = iota
	BaseError
	FileSystemError
	FileDoesNotExistError
	FileTypeError
	SyntaxError
)
