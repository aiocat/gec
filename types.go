package main

const (
	TYPE_INT       = -1
	TYPE_STRING    = -2
	TYPE_COMPARE   = -3
	TYPE_FUNCTION  = -4
	TYPE_DOUBLEDOT = -5
)

const (
	COMMAND_PUSH    = iota
	COMMAND_HALT    = iota
	COMMAND_END     = iota
	COMMAND_DUMP    = iota
	COMMAND_DUMPC   = iota
	COMMAND_CALL    = iota
	COMMAND_DUP     = iota
	COMMAND_ROUNDED = iota
	COMMAND_IF      = iota
	COMMAND_ELSE    = iota

	COMMAND_ADD = iota
	COMMAND_REM = iota
	COMMAND_MUL = iota
	COMMAND_DIV = iota

	TOTAL_COMMAND_COUNT = iota
)

type Token struct {
	Key, Line int
	Value     string
}
