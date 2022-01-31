package main

import (
	"fmt"
	"strings"
)

type Compiler struct {
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
	c.Source = "std::stack<int> stack;\nint _rounded, _gec_one, _gec_two = 0;\n"

	for index, token := range c.Tokens {
		if c.ShouldIgnore(index) {
			continue
		}

		if token.Key == TYPE_FUNCTION {
			funcVariables := []string{}
			variableFormat := ""

			for i, t := range c.Tokens[index+1:] {
				if t.Key == TYPE_STRING {
					funcVariables = append(funcVariables, t.Value)
					c.Ignore = append(c.Ignore, index+i)
				} else if t.Key == TYPE_DOUBLEDOT {
					break
				} else {
					panic("Function argument names must be string")
				}
			}

			for _, vari := range funcVariables {
				variableFormat += "int " + vari + ","
			}

			variableFormat = strings.TrimRight(variableFormat, ",")
			c.Source += fmt.Sprintf("int %s(%s){\n", token.Value, variableFormat)
		} else if token.Key == COMMAND_PUSH {
			if c.Tokens[index+1].Key == TYPE_INT || c.Tokens[index+1].Key == TYPE_STRING {
				c.Source += fmt.Sprintf("stack.push(%v);\n", c.Tokens[index+1].Value)
			} else {
				panic("Push command only accepts integer or variable")
			}
		} else if token.Key == COMMAND_END {
			c.Source += "}\n"
		} else if token.Key == COMMAND_HALT {
			if c.Tokens[index+1].Key == TYPE_INT || c.Tokens[index+1].Key == TYPE_STRING {
				c.Source += fmt.Sprintf("return %v;\n", c.Tokens[index+1].Value)
			} else {
				panic("Halt command only accepts integer or variable")
			}
		} else if token.Key == COMMAND_DUMP {
			if c.Tokens[index+1].Key == TYPE_INT || c.Tokens[index+1].Key == TYPE_STRING {
				c.Source += "std::cout << " + c.Tokens[index+1].Value + ";\n"
			} else {
				c.Source += "_gec_one = stack.top();\nstack.pop();\nstd::cout << _gec_one;\n"
			}
		} else if token.Key == COMMAND_CALL {
			if c.Tokens[index+1].Key == TYPE_FUNCTION {
				c.Ignore = append(c.Ignore, index+1)
				funcVariables := []string{}

				for i, t := range c.Tokens[index+2:] {
					if t.Key == TYPE_STRING || t.Key == TYPE_INT {
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
				panic("You only can call functions")
			}
		} else if token.Key == COMMAND_DUP {
			if c.Tokens[index+1].Key == TYPE_STRING {
				c.Ignore = append(c.Ignore, index+1)
				c.Source += "int " + c.Tokens[index+1].Value + " = stack.top();\nstack.pop();\n"
			} else {
				c.Source += "stack.push(stack.top());\n"
			}
		} else if token.Key == COMMAND_ADD {
			c.Source += "_gec_one = stack.top();\nstack.pop();\n_gec_two = stack.top();\nstack.pop();\nstack.push(_gec_one+_gec_two);\n"
		} else if token.Key == COMMAND_REM {
			c.Source += "_gec_one = stack.top();\nstack.pop();\n_gec_two = stack.top();\nstack.pop();\nstack.push(_gec_two-_gec_one);\n"
		} else if token.Key == COMMAND_MUL {
			c.Source += "_gec_one = stack.top();\nstack.pop();\n_gec_two = stack.top();\nstack.pop();\nstack.push(_gec_one*_gec_two);\n"
		} else if token.Key == COMMAND_DIV {
			c.Source += "_gec_one = stack.top();\nstack.pop();\n_gec_two = stack.top();\nstack.pop();\nif(_gec_two%_gec_one==0){\nstack.push(_gec_two/_gec_one);\n_rounded = 0;}else{\nstack.push((int)(_gec_two/_gec_one));\n_rounded = 1;}\n"
		} else if token.Key == COMMAND_ROUNDED {
			c.Source += "stack.push(_rounded);\n"
		} else if token.Key == COMMAND_DUMPC {
			if c.Tokens[index+1].Key == TYPE_INT || c.Tokens[index+1].Key == TYPE_STRING {
				c.Source += "std::cout << (char)" + c.Tokens[index+1].Value + ";\n"
			} else {
				c.Source += "_gec_one = stack.top();\nstack.pop();\nstd::cout << (char)_gec_one;\n"
			}
		}
	}

	includeCompile := ""

	for _, include := range c.Includes {
		includeCompile += "#include <" + include + ">\n"
	}

	c.Source = includeCompile + c.Source
}
