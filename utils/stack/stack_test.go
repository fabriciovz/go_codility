package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("", func(t *testing.T) {
		myStack := &stack{}
		fmt.Println(myStack)
		fmt.Println(myStack.IsEmpty())
		myStack.Push(100)
		myStack.Push(200)
		myStack.Push(300)
		fmt.Println(myStack)
		fmt.Println(myStack.Pop())
		fmt.Println(myStack)
		fmt.Println(myStack.Top())
		fmt.Println(myStack)
		fmt.Println(myStack.IsEmpty())
	})
}
