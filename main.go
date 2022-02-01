package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		panic("File name must be given")
	}

	err := os.Chdir(path.Dir(os.Args[1]))

	if err != nil {
		panic(err)
	}

	fileName := path.Base(os.Args[1])
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

	noExtFileName := fileName[:len(fileName)-len(path.Ext(fileName))]
	cPlusPlusFile := noExtFileName + ".cpp"

	fmt.Printf("[INFO]: Writing compiled source into %s\n", cPlusPlusFile)
	err = os.WriteFile(cPlusPlusFile, []byte(compiler.Source), 0666)

	if err != nil {
		panic(err)
	}

	fmt.Printf("[INFO]: Compiling the C++ source code with \"g++\" to: \"%s\"\n", noExtFileName)
	_, err = exec.Command("g++", cPlusPlusFile, "-o", noExtFileName).CombinedOutput()

	if err != nil {
		panic(err)
	}
}
