package main

const (
	TYPE_INT       = -1
	TYPE_STRING    = -2
	TYPE_VAR       = -2
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
	COMMAND_MOVE    = iota
	COMMAND_IMPORT  = iota
	COMMAND_BUF     = iota
	COMMAND_WHILE   = iota
	COMMAND_SKIP    = iota
	COMMAND_BREAK   = iota
	COMMAND_GEN     = iota
	COMMAND_MODULE  = iota
	COMMAND_USEMOD  = iota
	COMMAND_POP     = iota
	COMMAND_INPUT   = iota

	COMMAND_ADD = iota
	COMMAND_SUB = iota
	COMMAND_MUL = iota
	COMMAND_DIV = iota

	TOTAL_COMMAND_COUNT = iota
)

type Token struct {
	Key, Line int
	Value     string
}
