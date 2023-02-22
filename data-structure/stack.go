package data_structure

import "fmt"

type StackNode struct {
	value interface{}
	next  *StackNode
}

// NewStackNode takes a value and returns a stack node pointed to with field
// value as the argument.
func NewStackNode(value interface{}) *StackNode {
	return &StackNode{
		value: value,
		next:  nil,
	}
}

type Stack struct {
	top    *StackNode
	length int
}

// NewStack creates a new Stack with the first value in the stack being the
// value passed.
func NewStack(value interface{}) Stack {
	s := NewStackNode(value)
	return Stack{
		top:    s,
		length: 1,
	}
}

// Lookup
// Peek
// Push
// Pop

// Push adds the new element to the top/beginning of the stack.
//
// O(1)
func (s *Stack) Push(value interface{}) {
	// create a new node
	// point the new node next to the top
	// point the top to the new node
	// increment the length
	n := NewStackNode(value)
	n.next = s.top
	s.top = n
	s.length += 1
}

// Pop removes the last element added / the top of the stack and returns the
// value that the element contained.
//
// O(1)
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	// get the top node and the next node
	// point the top to next node
	// decrement the length
	// return the top node's value
	n := s.top
	s.top = n.next
	s.length -= 1
	return n.value
}

// Peek returns the value for the next value in if the Pop method were to be
// used.
//
// O(1)
func (s *Stack) Peek() interface{} {
	return s.top.value
}

func (s *Stack) String() string {
	node := s.top
	str := fmt.Sprintf("in --+  +-> out\n     %v\n", node.value)
	for {
		node = node.next
		if node == nil {
			break
		}
		str += fmt.Sprintf("     %v\n", node.value)
	}
	return str
}
