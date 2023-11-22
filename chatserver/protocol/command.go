package protocol

import "errors"

var (
	UnknownCommand = errors.New("Unknown command")
)

type SendCmd struct {
	Message string
}

type NameCmd struct {
	Name string
}

type MessageCmd struct {
	Name    string
	Message string
}
