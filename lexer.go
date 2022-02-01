package main

import (
	"fmt"
	"strconv"
)

type Lexer struct {
	CollectingString bool
	CurrentLine      int

	Tokens              []*Token
	Raw, CollectedToken string
}

func NewLexer(data string) *Lexer {
	return &Lexer{
		Raw:         data,
		CurrentLine: 1,
	}
}

func (l *Lexer) Run() {
	for index, char := range l.Raw {
		if !l.CollectingString {
			switch char {
			case '\r':
				continue
			case '\n':
				l.DetermineToken()
				l.CurrentLine++
				continue
			case ' ':
				l.DetermineToken()
			case ':':
				l.DetermineToken()

				l.Tokens = append(l.Tokens, &Token{
					Key:   TYPE_DOUBLEDOT,
					Line:  l.CurrentLine,
					Value: "",
				})
			case '"':
				l.CollectingString = true
				continue
			default:
				l.CollectedToken += string(char)

				if index+1 == len(l.Raw) {
					l.DetermineToken()
				}
			}
		} else {
			if char == '"' && l.Raw[index-1] != '\\' {
				l.Tokens = append(l.Tokens, &Token{
					Key:   TYPE_STRING,
					Line:  l.CurrentLine,
					Value: l.CollectedToken,
				})
				l.CollectedToken = ""
				l.CollectingString = false

				continue
			}

			l.CollectedToken += string(char)
		}
	}
}

func (l *Lexer) DetermineToken() {
	if len(l.CollectedToken) == 0 {
		return
	}

	if TOTAL_COMMAND_COUNT != 23 {
		panic("Mismatched number of commands")
	}

	switch l.CollectedToken {
	case "push":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_PUSH,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "add":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_ADD,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "halt":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_HALT,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "end":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_END,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "dump":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_DUMP,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "call":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_CALL,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "dup":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_DUP,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "move":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_MOVE,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "sub":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_SUB,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "mul":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_MUL,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "div":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_DIV,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "rounded":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_ROUNDED,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "dumpc":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_DUMPC,
			Line:  l.CurrentLine,
			Value: "",
		})
	case ">":
		l.Tokens = append(l.Tokens, &Token{
			Key:   TYPE_COMPARE,
			Line:  l.CurrentLine,
			Value: ">",
		})
	case "<":
		l.Tokens = append(l.Tokens, &Token{
			Key:   TYPE_COMPARE,
			Line:  l.CurrentLine,
			Value: "<",
		})
	case "=":
		l.Tokens = append(l.Tokens, &Token{
			Key:   TYPE_COMPARE,
			Line:  l.CurrentLine,
			Value: "=",
		})
	case "!":
		l.Tokens = append(l.Tokens, &Token{
			Key:   TYPE_COMPARE,
			Line:  l.CurrentLine,
			Value: "!",
		})
	case ">=":
		l.Tokens = append(l.Tokens, &Token{
			Key:   TYPE_COMPARE,
			Line:  l.CurrentLine,
			Value: ">=",
		})
	case "<=":
		l.Tokens = append(l.Tokens, &Token{
			Key:   TYPE_COMPARE,
			Line:  l.CurrentLine,
			Value: "<=",
		})
	case "if":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_IF,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "else":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_ELSE,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "import":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_IMPORT,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "buf":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_BUF,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "while":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_WHILE,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "skip":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_SKIP,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "break":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_BREAK,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "gen":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_GEN,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "module":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_MODULE,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "usemod":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_USEMOD,
			Line:  l.CurrentLine,
			Value: "",
		})
	default:
		if l.CollectedToken[0] == '#' {
			l.Tokens = append(l.Tokens, &Token{
				Key:   TYPE_FUNCTION,
				Line:  l.CurrentLine,
				Value: l.CollectedToken[1:],
			})
		} else if l.CollectedToken[0] == '$' {
			if len(l.CollectedToken) == 1 {
				l.Tokens = append(l.Tokens, &Token{
					Key:   TYPE_VAR,
					Line:  l.CurrentLine,
					Value: "stack.top()",
				})
			} else {
				l.Tokens = append(l.Tokens, &Token{
					Key:   TYPE_VAR,
					Line:  l.CurrentLine,
					Value: l.CollectedToken[1:],
				})
			}
		} else {
			_, err := strconv.Atoi(l.CollectedToken)

			if err != nil {
				panic(fmt.Sprintf("[L%d]: Undetermined token error", l.CurrentLine))
			}

			l.Tokens = append(l.Tokens, &Token{
				Key:   TYPE_INT,
				Line:  l.CurrentLine,
				Value: l.CollectedToken,
			})
		}
	}

	l.CollectedToken = ""
}
