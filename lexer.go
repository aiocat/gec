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
		Raw: data,
	}
}

func (l *Lexer) Run() {
	for index, char := range l.Raw {
		if !l.CollectingString {
			switch char {
			case '\r':
				continue
			case '\n':
				l.CurrentLine++
				l.DetermineToken()
				continue
			case ' ':
				l.DetermineToken()
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
					Key:   TYPE_INT,
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

	if TOTAL_COMMAND_COUNT != 3 {
		panic("Mismatched number of commands")
	}

	switch l.CollectedToken {
	case "push":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_PUSH,
			Line:  l.CurrentLine,
			Value: nil,
		})
	case "add":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_ADD,
			Line:  l.CurrentLine,
			Value: nil,
		})
	case "halt":
		l.Tokens = append(l.Tokens, &Token{
			Key:   COMMAND_HALT,
			Line:  l.CurrentLine,
			Value: nil,
		})
	default:
		output, err := strconv.Atoi(l.CollectedToken)

		if err != nil {
			panic("Undetermined token error")
		}

		l.Tokens = append(l.Tokens, &Token{
			Key:   TYPE_INT,
			Line:  l.CurrentLine,
			Value: output,
		})
	}

	l.CollectedToken = ""
}
