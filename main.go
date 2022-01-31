package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("File name must be given")
	}

	fileName := os.Args[1]

	fileBody, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	lexer := NewLexer(string(fileBody))
	lexer.Run()

	for i, token := range lexer.Tokens {
		fmt.Println(i, token)
	}
}
