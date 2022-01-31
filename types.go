package main

const (
	TYPE_INT    = -1
	TYPE_STRING = -2
)

const (
	COMMAND_PUSH        = iota
	COMMAND_ADD         = iota
	COMMAND_HALT        = iota
	TOTAL_COMMAND_COUNT = iota
)

type Token struct {
	Key, Line int
	Value     interface{}
}
