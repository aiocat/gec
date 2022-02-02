package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	STACK_COUNT   = 1
	CURRENT_STACK = 0
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
		Includes: []string{"stack", "iostream", "vector"},
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
		c.Source = "std::stack<int> stack_0;\nint _rounded, _gec_one, _gec_two = 0;\n"
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
				c.Source += fmt.Sprintf("stack_%d.push(%v);\n", CURRENT_STACK, c.Tokens[index+1].Value)
			} else {
				panic(fmt.Sprintf("[L%d]: Push command only accepts integer or variable", token.Line))
			}
		} else if token.Key == COMMAND_END {
			c.Source += "};\n"
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
				c.Source += fmt.Sprintf("_gec_one = stack_%d.top();\nstack_%d.pop();\nstd::cout << _gec_one;\n", CURRENT_STACK, CURRENT_STACK)
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
				c.Source += c.Tokens[index+1].Value + fmt.Sprintf(" = stack_%d.top();\n", CURRENT_STACK)
			} else {
				c.Source += fmt.Sprintf("stack_%d.push(stack_%d.top());\n", CURRENT_STACK, CURRENT_STACK)
			}
		} else if token.Key == COMMAND_ADD {
			c.Source += fmt.Sprintf("_gec_one = stack_%d.top();\nstack_%d.pop();\n_gec_two = stack_%d.top();\nstack_%d.pop();\nstack_%d.push(_gec_one+_gec_two);\n", CURRENT_STACK, CURRENT_STACK, CURRENT_STACK, CURRENT_STACK, CURRENT_STACK)
		} else if token.Key == COMMAND_SUB {
			c.Source += fmt.Sprintf("_gec_one = stack_%d.top();\nstack_%d.pop();\n_gec_two = stack_%d.top();\nstack_%d.pop();\nstack_%d.push(_gec_two-_gec_one);\n", CURRENT_STACK, CURRENT_STACK, CURRENT_STACK, CURRENT_STACK, CURRENT_STACK)
		} else if token.Key == COMMAND_MUL {
			c.Source += fmt.Sprintf("_gec_one = stack_%d.top();\nstack_%d.pop();\n_gec_two = stack_%d.top();\nstack_%d.pop();\nstack_%d.push(_gec_one*_gec_two);\n", CURRENT_STACK, CURRENT_STACK, CURRENT_STACK, CURRENT_STACK, CURRENT_STACK)
		} else if token.Key == COMMAND_DIV {
			c.Source += fmt.Sprintf("_gec_one = stack_%d.top();\nstack_%d.pop();\n_gec_two = stack_%d.top();\nstack_%d.pop();\nif(_gec_two%%_gec_one==0){\nstack_%d.push(_gec_two/_gec_one);\n_rounded = 0;}else{\nstack_%d.push((int)(_gec_two/_gec_one));\n_rounded = 1;}\n", CURRENT_STACK, CURRENT_STACK, CURRENT_STACK, CURRENT_STACK, CURRENT_STACK, CURRENT_STACK)
		} else if token.Key == COMMAND_ROUNDED {
			c.Source += fmt.Sprintf("stack_%d.push(_rounded);\n", CURRENT_STACK)
		} else if token.Key == COMMAND_DUMPC {
			if index+1 < len(c.Tokens) && (c.Tokens[index+1].Key == TYPE_INT || c.Tokens[index+1].Key == TYPE_VAR) {
				c.Source += "std::cout << (char)" + c.Tokens[index+1].Value + ";\n"
			} else {
				c.Source += fmt.Sprintf("_gec_one = stack_%d.top();\nstack_%d.pop();\nstd::cout << (char)_gec_one;\n", CURRENT_STACK, CURRENT_STACK)
			}
		} else if token.Key == COMMAND_IF {
			if index+3 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_COMPARE && (c.Tokens[index+2].Key == TYPE_INT || c.Tokens[index+2].Key == TYPE_VAR) && (c.Tokens[index+3].Key == TYPE_INT || c.Tokens[index+3].Key == TYPE_VAR) {
				c.Source += "if(" + c.Tokens[index+2].Value + c.Tokens[index+1].Value + c.Tokens[index+3].Value + "){\n"
			} else if index+2 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_COMPARE && (c.Tokens[index+2].Key == TYPE_INT || c.Tokens[index+2].Key == TYPE_VAR) {
				c.Source += fmt.Sprintf("_gec_one = stack_%d.top();\nstack_%d.pop();\nif(_gec_one"+c.Tokens[index+1].Value+c.Tokens[index+2].Value+"){\n", CURRENT_STACK, CURRENT_STACK)
			} else if index+1 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_COMPARE {
				c.Source += fmt.Sprintf("_gec_one = stack_%d.top();\nstack_%d.pop();\n_gec_two = stack_%d.top();\nstack_%d.pop();\nif(_gec_one"+c.Tokens[index+1].Value+"_gec_two){\n", CURRENT_STACK, CURRENT_STACK, CURRENT_STACK, CURRENT_STACK)
			} else {
				panic(fmt.Sprintf("[L%d]: Wrong usage for if statement. Please check the docs", token.Line))
			}
		} else if token.Key == COMMAND_ELSE {
			c.Source += "}else{"
		} else if token.Key == COMMAND_MOVE {
			if index+1 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_VAR {
				c.Ignore = append(c.Ignore, index+1)
				c.Source += c.Tokens[index+1].Value + fmt.Sprintf(" = stack_%d.top();\nstack_%d.pop();\n", CURRENT_STACK, CURRENT_STACK)
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

				c.Source += fmt.Sprintf("stack_%d.push(0);\n", CURRENT_STACK)
				for _, char := range reverse(c.Tokens[index+1].Value) {
					c.Source += fmt.Sprintf("stack_%d.push(%d);\n", CURRENT_STACK, int(char))
				}
			}
		} else if token.Key == COMMAND_WHILE {
			if index+3 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_COMPARE && (c.Tokens[index+2].Key == TYPE_INT || c.Tokens[index+2].Key == TYPE_VAR) && (c.Tokens[index+3].Key == TYPE_INT || c.Tokens[index+3].Key == TYPE_VAR) {
				c.Source += "while(" + c.Tokens[index+2].Value + c.Tokens[index+1].Value + c.Tokens[index+3].Value + "){\n"
			} else if index+2 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_COMPARE && (c.Tokens[index+2].Key == TYPE_INT || c.Tokens[index+2].Key == TYPE_VAR) {
				c.Source += fmt.Sprintf("while(stack_%d.top() %s %s){\n", CURRENT_STACK, c.Tokens[index+1].Value, c.Tokens[index+2].Value)
			} else {
				panic(fmt.Sprintf("[L%d]: Wrong usage for while statement. Please check the docs", token.Line))
			}
		} else if token.Key == COMMAND_SKIP {
			c.Source += "continue;\n"
		} else if token.Key == COMMAND_BREAK {
			c.Source += "break;\n"
		} else if token.Key == COMMAND_GEN {
			if index+1 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_VAR {
				c.Ignore = append(c.Ignore, index+1)
				c.Source += "int " + c.Tokens[index+1].Value + ";\n"
			} else {
				panic(fmt.Sprintf("[L%d]: Gen command only accepts variable", token.Line))
			}
		} else if token.Key == COMMAND_MODULE {
			if index+1 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_VAR {
				c.Ignore = append(c.Ignore, index+1)
				c.Source += "class " + c.Tokens[index+1].Value + "{\n  public:\n"
			} else {
				panic(fmt.Sprintf("[L%d]: Module command only accepts variable", token.Line))
			}
		} else if token.Key == COMMAND_USEMOD {
			if index+2 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_VAR && c.Tokens[index+2].Key == TYPE_VAR {
				c.Ignore = append(c.Ignore, index+1)
				c.Ignore = append(c.Ignore, index+2)

				c.Source += c.Tokens[index+1].Value + " " + c.Tokens[index+2].Value + ";\n"
			} else {
				panic(fmt.Sprintf("[L%d]: Usemod command only accepts variable", token.Line))
			}
		} else if token.Key == COMMAND_POP {
			c.Source += fmt.Sprintf("stack_%d.pop();\n", CURRENT_STACK)
		} else if token.Key == COMMAND_INPUT {
			c.Source += fmt.Sprintf("std::cin >> _gec_one;\nstack_%d.push(_gec_one);\n", CURRENT_STACK)
		} else if token.Key == COMMAND_NST {
			STACK_COUNT++
			c.Source += fmt.Sprintf("std::stack<int> stack_%d;\n", STACK_COUNT-1)
		} else if token.Key == COMMAND_SWITCH {
			if index+1 < len(c.Tokens) && c.Tokens[index+1].Key == TYPE_INT {
				out, err := strconv.Atoi(c.Tokens[index+1].Value)

				if err != nil {
					panic(fmt.Sprintf("[L%d]: Not a valid integer", c.Tokens[index+1].Line))
				}

				CURRENT_STACK = out
			} else {
				panic(fmt.Sprintf("[L%d]: Switch command only accepts integer", token.Line))
			}
		} else if token.Key == COMMAND_DST {
			STACK_COUNT--
			c.Source += fmt.Sprintf("delete &stack_%d;\n", CURRENT_STACK)
		} else if token.Key == COMMAND_REP {
			if index+1 < len(c.Tokens) && (c.Tokens[index+1].Key == TYPE_VAR || c.Tokens[index+1].Key == TYPE_INT) {
				c.Source += fmt.Sprintf("for(int _rep=0;_rep<%s;_rep++){\n", c.Tokens[index+1].Value)
			} else {
				c.Source += "_gen_one=stack.top();\nstack.pop();\nfor(int _rep=0;_rep<_gen_one;_rep++){\n"
			}
		}

	}

	includeCompile := ""

	for _, include := range c.Includes {
		includeCompile += "#include <" + include + ">\n"
	}

	c.Source = includeCompile + c.Source
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
