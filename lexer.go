package main

import "strconv"

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
			case '>':
				l.DetermineToken()

				l.Tokens = append(l.Tokens, &Token{
					Key:   TYPE_COMPARE,
					Line:  l.CurrentLine,
					Value: ">",
				})
			case '<':
				l.DetermineToken()

				l.Tokens = append(l.Tokens, &Token{
					Key:   TYPE_COMPARE,
					Line:  l.CurrentLine,
					Value: "<",
				})
			case '=':
				l.DetermineToken()

				l.Tokens = append(l.Tokens, &Token{
					Key:   TYPE_COMPARE,
					Line:  l.CurrentLine,
					Value: "==",
				})
			case '!':
				l.DetermineToken()

				l.Tokens = append(l.Tokens, &Token{
					Key:   TYPE_COMPARE,
					Line:  l.CurrentLine,
					Value: "!=",
				})
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

	if TOTAL_COMMAND_COUNT != 14 {
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
	case "int":
		l.Tokens = append(l.Tokens, &Token{
			Key:   TYPE_INT,
			Line:  l.CurrentLine,
			Value: "",
		})
	case "str":
		l.Tokens = append(l.Tokens, &Token{
			Key:   TYPE_STRING,
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
	case "rem":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_REM,
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
	default:
		if l.CollectedToken[0] == '#' {
			l.Tokens = append(l.Tokens, &Token{
				Key:   TYPE_FUNCTION,
				Line:  l.CurrentLine,
				Value: l.CollectedToken[1:],
			})
		} else {
			_, err := strconv.Atoi(l.CollectedToken)

			if err != nil {
				panic("Undetermined token error")
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
