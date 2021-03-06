package ast

import (
	"fmt"
	"log"

	"tango/src/token"
)

// Stack represents an array of Attribs which can also be used as a stack
type Stack []Attrib

// Empty checks if the stack is empty
func (s Stack) Empty() bool {
	return len(s) == 0
}

// Push pushes an element onto the stack
func (s Stack) Push(v Attrib) Stack {
	return append(s, v)
}

// Pop removes an element from the stack
func (s Stack) Pop() (Stack, Attrib) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s Stack) String() string {
	str := ""
	for i, value := range s {
		newline := false
		switch v := value.(type) {
		case *token.Token:
			str += fmt.Sprintf("%q", v)
			if token.TokMap.Type("stmt_end") == v.Type {
				str += "\n"
				newline = true
			}
		case *Node:
			str += fmt.Sprintf("%s", v)
		default:
			log.Fatalf("Unknown type: %T", v)
		}
		if i != len(s)-1 && !newline {
			str += " "
		}
	}
	return str
}
