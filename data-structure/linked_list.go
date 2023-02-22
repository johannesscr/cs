package data_structure

import (
	"fmt"
	"strings"
)

/*
	SINGLE LINKED LIST
*/

type SingleNode struct {
	Value    int
	NextNode *SingleNode
}

func NewSingleNode(value int) SingleNode {
	return SingleNode{
		Value: value,
	}
}

type SingleLinkedList struct {
	head   *SingleNode
	tail   *SingleNode
	length int
}

func NewSingleLinkedList(value int) SingleLinkedList {
	head := NewSingleNode(value)
	return SingleLinkedList{
		head:   &head,
		tail:   &head,
		length: 1,
	}
}

// Push adds a new value to the end of the SingleLinkedList.
//
// wil be O(1)
func (l *SingleLinkedList) Push(value int) {
	newNode := NewSingleNode(value)
	// here the current tail points to the new node
	l.tail.NextNode = &newNode

	// but now set the tail node to the new node, which is also pointed
	// to by the previous tail node
	l.tail = &newNode
	l.length += 1
}

// Prepend adds a new value to the beginning of the SingleLinkedList.
//
// O(1)
func (l *SingleLinkedList) Prepend(value int) {
	newNode := NewSingleNode(value)
	newNode.NextNode = l.head
	l.head = &newNode
	l.length += 1
}

// Pop removes the last item at the end of the SingleLinkedList.
//
// wil be O(n)
func (l *SingleLinkedList) Pop() int {
	v := l.tail.Value
	if l.length == 1 {
		return v
	}
	// remove the last node
	node := l.head
	for i := 1; i < l.length; i++ {
		// node before the tail node
		if i == l.length-1 {
			node.NextNode = nil // remove the pointer
			l.tail = node
			break
		}
		node = node.NextNode
	}
	// decrement the length
	l.length -= 1
	return v
}

func (l *SingleLinkedList) String() string {
	node := l.head
	s := fmt.Sprintf("%d --> ", node.Value)
	for i := 1; i < l.length; i++ {
		node = node.NextNode
		s += fmt.Sprintf("%d --> ", node.Value)
	}
	return strings.TrimSpace(s)
}

func (l *SingleLinkedList) Print() {
	fmt.Println(l.String())
}

// traverseToIndex loops through the linked list until it is it reaches the
// index, then returns the value pointed to at that index.
func (l *SingleLinkedList) traverseToIndex(index int) *SingleNode {
	node := l.head
	for i := 0; i < l.length; i++ {
		// this is the node at index i
		if i == index {
			return node
		}
		node = node.NextNode
	}
	return nil
}

// Insert wil insert an element at an index with the value provided.
//
// wil be O(n)
func (l *SingleLinkedList) Insert(index, value int) {
	// if in the beginning
	if index == 0 {
		l.Prepend(value)
		return
	}
	// if in the last position or further
	if index >= l.length-1 {
		l.Push(value)
		return
	}

	// here the node is at index i - 1, because
	// (i - 1) * --- *  (i)
	//            + if we insert here, then
	// (i - 1) * -- + (i) -- * (i + 1)

	// if somewhere in between
	newNode := NewSingleNode(value)
	node := l.traverseToIndex(index - 1) // O(n)
	newNode.NextNode = node.NextNode
	node.NextNode = &newNode
	l.length += 1
}

// Remove deletes a SingleNode from the SingleLinkedList at the index provided
// and returns the value of that node at that index.
//
// O(n)
func (l *SingleLinkedList) Remove(index int) int {
	// edge case at the beginning or end
	// index at the beginning
	if index == 0 {
		node := l.head
		l.head = node.NextNode
		l.length -= 1
		return node.Value
	}
	// index at the end
	if index >= l.length-1 {
		return l.Pop()
	}

	node := l.traverseToIndex(index - 1) // O(n)
	if node == nil {
		return 0
	}
	removeNode := node.NextNode
	node.NextNode = removeNode.NextNode
	l.length -= 1
	return removeNode.Value
}

//func (l *SingleLinkedList) Reverse() SingleLinkedList {
//	r := NewSingleLinkedList(l.head.Value)
//	node := l.head
//	for i := 1; i < l.length; i++ {
//		node = node.NextNode
//		r.Prepend(node.Value)
//	}
//	return r
//}

func (l *SingleLinkedList) Reverse() SingleLinkedList {
	first := l.head
	l.tail = first
	second := first.NextNode
	for second != nil {
		third := second.NextNode
		second.NextNode = first
		first = second
		second = third
	}
	l.head.NextNode = nil
	l.head = first
	return *l
}

/*
	DOUBLE LINKED LIST
*/
