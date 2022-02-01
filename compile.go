package main

import (
	"fmt"
	"os"
	"strings"
)

type Compiler struct {
	Blank    bool
	Source   string
	Ignore   []int
	Tokens   []*Token
	Includes []string
}

func NewCompiler(tokens []*Token) *Compiler {
	return &Compiler{
		Tokens:   tokens,
		Includes: []string{"stack", "iostream"},
	}
}

func (c *Compiler) ShouldIgnore(index int) bool {
	for _, i := range c.Ignore {
		if i == index {
			return true
		}
	}

	return false
}

func (c *Compiler) Run() {
	if !c.Blank {
		c.Source = "std::stack<int> stack;\nint _rounded, _gec_one, _gec_two = 0;\n"
	}

	for index, token := range c.Tokens {
		if c.ShouldIgnore(index) {
			continue
		}

		if token.Key == TYPE_FUNCTION {
			funcVariables := []string{}
			variableFormat := ""

			for i, t := range c.Tokens[index+1:] {
				if t.Key == TYPE_VAR {
					funcVariables = append(funcVariables, t.Value)
					c.Ignore = append(c.Ignore, index+i)
				} else if t.Key == TYPE_DOUBLEDOT {
					break
				} else {
					panic(fmt.Sprintf("[L%d]: Function argument names must be variable", token.Line))
				}
			}

			for _, vari := range funcVariables {
				variableFormat += "int " + vari + ","
			}

			variableFormat = strings.TrimRight(variableFormat, ",")
			c.Source += fmt.Sprintf("int %s(%s){\n", token.Value, variableFormat)
		} else if token.Key == COMMAND_PUSH {
			if index+1 < len(c.Tokens) && (c.Tokens[index+1].Key == TYPE_INT || c.Tokens[index+1].Key == TYPE_VAR) {
				c.Source += fmt.Sprintf("stack.push(%v);\n", c.Tokens[index+1].Value)
			} else {
				panic(fmt.Sprintf("[L%d]: Push command only accepts integer or variable", token.Line))
			}
		} else if token.Key == COMMAND_END {
			c.Source += "}\n"
		} else if token.Key == COMMAND_HALT {
			if index+1 < len(c.Tokens) && (c.Tokens[index+1].Key == TYPE_INT || c.Tokens[index+1].Key == TYPE_VAR) {
				c.Source += fmt.Sprintf("return %v;\n", c.Tokens[index+1].Value)
			} else {
				panic(fmt.Sprintf("[L%d]: Halt command only accepts integer or variable", token.Line))
			}
		} else if token.Key == COMMAND_DUMP {
			if index+1 < len(c.Tokens) && (c.Tokens[index+1].Key == TYPE_INT || c.Tokens[index+1].Key == TYPE_VAR) {
				c.Source += "std::cout << " + c.Tokens[index+1].Value + ";\n"
			} else {
				c.Source += "_gec_one = stack.top();\nstack.pop();\nstd::cout << _gec_one;\n"
			}
		} else if token.Key == COMMAND_CALL {
			if index+1 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_FUNCTION {
				c.Ignore = append(c.Ignore, index+1)
				funcVariables := []string{}

				for i, t := range c.Tokens[index+2:] {
					if t.Key == TYPE_VAR || t.Key == TYPE_INT {
						funcVariables = append(funcVariables, t.Value)
						c.Ignore = append(c.Ignore, index+i)
					} else if t.Key == TYPE_DOUBLEDOT {
						break
					} else {
						panic("Function argument names must be string")
					}
				}

				c.Source += fmt.Sprintf("_gec_one=%s(%s);\nif(_gec_one!=0){return _gec_one;}\n", c.Tokens[index+1].Value, strings.Join(funcVariables, ","))
			} else {
				panic(fmt.Sprintf("[L%d]: You only can call functions", token.Line))
			}
		} else if token.Key == COMMAND_DUP {
			if index+1 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_VAR {
				c.Ignore = append(c.Ignore, index+1)
				c.Source += "int " + c.Tokens[index+1].Value + " = stack.top();\n"
			} else {
				c.Source += "stack.push(stack.top());\n"
			}
		} else if token.Key == COMMAND_ADD {
			c.Source += "_gec_one = stack.top();\nstack.pop();\n_gec_two = stack.top();\nstack.pop();\nstack.push(_gec_one+_gec_two);\n"
		} else if token.Key == COMMAND_REM {
			c.Source += "_gec_one = stack.top();\nstack.pop();\n_gec_two = stack.top();\nstack.pop();\nstack.push(_gec_one-gec_two);\n"
		} else if token.Key == COMMAND_MUL {
			c.Source += "_gec_one = stack.top();\nstack.pop();\n_gec_two = stack.top();\nstack.pop();\nstack.push(_gec_one*_gec_two);\n"
		} else if token.Key == COMMAND_DIV {
			c.Source += "_gec_one = stack.top();\nstack.pop();\n_gec_two = stack.top();\nstack.pop();\nif(_gec_one%_gec_two==0){\nstack.push(_gec_one/_gec_two);\n_rounded = 0;}else{\nstack.push((int)(_gec_one/_gec_two));\n_rounded = 1;}\n"
		} else if token.Key == COMMAND_ROUNDED {
			c.Source += "stack.push(_rounded);\n"
		} else if token.Key == COMMAND_DUMPC {
			if index+1 < len(c.Tokens) && (c.Tokens[index+1].Key == TYPE_INT || c.Tokens[index+1].Key == TYPE_VAR) {
				c.Source += "std::cout << (char)" + c.Tokens[index+1].Value + ";\n"
			} else {
				c.Source += "_gec_one = stack.top();\nstack.pop();\nstd::cout << (char)_gec_one;\n"
			}
		} else if token.Key == COMMAND_IF {
			if index+3 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_COMPARE && (c.Tokens[index+2].Key == TYPE_INT || c.Tokens[index+2].Key == TYPE_VAR) && (c.Tokens[index+3].Key == TYPE_INT || c.Tokens[index+3].Key == TYPE_VAR) {
				c.Source += "if(" + c.Tokens[index+2].Value + c.Tokens[index+1].Value + c.Tokens[index+3].Value + "){\n"
			} else if index+2 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_COMPARE && (c.Tokens[index+2].Key == TYPE_INT || c.Tokens[index+2].Key == TYPE_VAR) {
				c.Source += "_gec_one = stack.top();\nstack.pop();\nif(_gec_one" + c.Tokens[index+1].Value + c.Tokens[index+2].Value + "){\n"
			} else if index+1 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_COMPARE {
				c.Source += "_gec_one = stack.top();\nstack.pop();\n_gec_two = stack.top();\nstack.pop();if(_gec_one" + c.Tokens[index+1].Value + "_gec_two){\n"
			} else {
				panic(fmt.Sprintf("[L%d]: Wrong usage for if statement. Please check the docs", token.Line))
			}
		} else if token.Key == COMMAND_ELSE {
			c.Source += "}else{"
		} else if token.Key == COMMAND_MOVE {
			if index+1 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_VAR {
				c.Ignore = append(c.Ignore, index+1)
				c.Source += "int " + c.Tokens[index+1].Value + " = stack.top();\nstack.pop();\n"
			} else {
				panic(fmt.Sprintf("[L%d]: Move command only accepts variable", token.Line))
			}
		} else if token.Key == COMMAND_IMPORT {
			if index+1 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_STRING {
				fileBody, err := os.ReadFile(c.Tokens[index+1].Value)

				if err != nil {
					panic(err)
				}

				lexer := NewLexer(string(fileBody))
				lexer.Run()

				compiler := &Compiler{
					Tokens: lexer.Tokens,
					Blank:  true,
				}
				compiler.Run()
				c.Source += compiler.Source
			}
		} else if token.Key == COMMAND_BUF {
			if index+1 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_STRING {
				if len(c.Tokens[index+1].Value) == 0 {
					panic("String can't be blank")
				}

				for _, char := range c.Tokens[index+1].Value {
					c.Source += fmt.Sprintf("stack.push(%d);\n", int(char))
				}
			}
		}
	}

	includeCompile := ""

	for _, include := range c.Includes {
		includeCompile += "#include <" + include + ">\n"
	}

	c.Source = includeCompile + c.Source
}
