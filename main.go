package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"os/exec"
	"path"
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

	fmt.Println("[INFO]: Starting lexical analysis")
	lexer := NewLexer(string(fileBody))
	lexer.Run()
	fmt.Printf("[INFO]: Lexical analysis completed and found %d tokens\n", len(lexer.CollectedToken))

	fmt.Println("[INFO]: Transpiling into C++")
	compiler := NewCompiler(lexer.Tokens)
	compiler.Run()

	cPlusPlusFile := fmt.Sprintf("%x.cpp", sha1.Sum(fileBody))

	fmt.Printf("[INFO]: Writing compiled source into %s\n", cPlusPlusFile)
	err = os.WriteFile(cPlusPlusFile, []byte(compiler.Source), 0666)

	if err != nil {
		panic(err)
	}

	noExtFileName := fileName[:len(fileName)-len(path.Ext(fileName))]

	fmt.Printf("[INFO]: Compiling the C++ source code with \"g++\" to: \"%s\"\n", noExtFileName)
	_, err = exec.Command("g++", cPlusPlusFile, "-o", noExtFileName).CombinedOutput()

	if err != nil {
		panic(err)
	}
}
