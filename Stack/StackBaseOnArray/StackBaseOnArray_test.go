package StackBaseOnArray

import (
	"fmt"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := NewStack()

	for i := 0; i < 10; i++{
		stack.Push(i)
	}

	stack.Print()

	fmt.Println()

	for i := 0; i < 10; i++{
		stack.Pop()
	}

	stack.Print()
}

