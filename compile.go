package main

import (
	"fmt"
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
	c.Source = "std::stack<int> stack;\n"

	for index, token := range c.Tokens {
		if c.ShouldIgnore(index) {
			continue
		}

		if token.Key == TYPE_FUNCTION {
			funcVariables := []string{}
			variableFormat := ""

			for i, t := range c.Tokens[index+1:] {
				if t.Key == TYPE_STRING {
					funcVariables = append(funcVariables, t.Value.(string))
					c.Ignore = append(c.Ignore, index+i)
				} else if t.Key == TYPE_DOUBLEDOT {
					break
				} else {
					panic("Function argument names must be string")
				}
			}

			for _, vari := range funcVariables {
				variableFormat += "int " + vari
			}

			c.Source += fmt.Sprintf("int %s(%s){\n", token.Value, variableFormat)
		} else if token.Key == COMMAND_PUSH {
			if c.Tokens[index+1].Key == TYPE_INT || c.Tokens[index+1].Key == TYPE_STRING {
				c.Source += fmt.Sprintf("stack.push(%v);\n", c.Tokens[index+1].Value)
			} else {
				panic("Push command only accepts integer or variable")
			}
		} else if token.Key == COMMAND_ADD {
			c.Source += "int _gec_one = stack.top();\nstack.pop();\nint _gec_two = stack.top();\nstack.pop();\nstack.push(_gec_one+_gec_two);\n"
		} else if token.Key == COMMAND_END {
			c.Source += "}\n"
		} else if token.Key == COMMAND_HALT {
			if c.Tokens[index+1].Key == TYPE_INT || c.Tokens[index+1].Key == TYPE_STRING {
				c.Source += fmt.Sprintf("return %v;\n", c.Tokens[index+1].Value)
			} else {
				panic("Halt command only accepts integer or variable")
			}
		} else if token.Key == COMMAND_DUMP {
			c.Source += "int _gec_dump = stack.top();\nstack.pop();\nstd::cout << _gec_dump;\n"
		}
	}

	includeCompile := ""

	for _, include := range c.Includes {
		includeCompile += "#include <" + include + ">\n"
	}

	c.Source = includeCompile + c.Source
}
