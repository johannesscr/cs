package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := &Stack{}
	stack.Push(1)
	stack.Push("2")
	fmt.Println(stack)
	v := stack.PopValue()
	fmt.Printf("%v of type %T\n", v, v)
	v = stack.PopValue()
	fmt.Printf("%v of type %T\n", v, v)
}
