package data_structure

import (
	"testing"
)

func TestNewSingleLinkedList(t *testing.T) {
	ll := NewSingleLinkedList(5)
	if ll.length != 1 {
		t.Errorf("expected a new single linked list to have an initial length of 1 got %d", ll.length)
	}
	// now get the first node at the head
	head := ll.head
	if head.Value != 5 {
		t.Errorf("expected new single linked list to have a head with value 5 got %d", head.Value)
	}
	if head.NextNode != nil {
		t.Errorf("expected new single linked list to have a head with next pointer to be nil got %v", head.NextNode)
	}
}

func TestSingleLinkedList_Push(t *testing.T) {
	ll := NewSingleLinkedList(5)
	ll.Push(10)
	if ll.length != 2 {
		t.Errorf("expected single linked list to have length 2 got %d", ll.length)
	}
	// here we are not testing if it is the same pointer/address, but we are
	// testing that the value of the pointer is the same
	if ll.head.Value != 5 {
		t.Errorf("expected head node value to be %v got %v", 5, ll.head.Value)
	}
	if ll.head.NextNode.Value != 10 {
		t.Errorf("expected the head node to point to the tail node %v got %v", 10, ll.head.NextNode.Value)
	}
	if ll.tail.Value != 10 {
		t.Errorf("expected tail node to be %v got %v", 10, ll.tail)
	}
}

func TestSingleLinkedList_Pop(t *testing.T) {
	ll := NewSingleLinkedList(5)
	ll.Push(10)
	ll.Push(7)
	ll.Push(3)

	if ll.length != 4 {
		t.Errorf("expected linked list to have length 4 got %d", ll.length)
	}
	// Pop 1
	v := ll.Pop()
	if ll.length != 3 {
		t.Errorf("expected linked list to have length 3 got %d", ll.length)
	}
	if v != 3 {
		t.Errorf("expected poped element to have value 3 got %d", v)
	}
	if ll.tail.Value != 7 {
		t.Errorf("expected new tail to have value 7 got %d", ll.tail.Value)
	}
	// Pop 2
	v = ll.Pop()
	if ll.length != 2 {
		t.Errorf("expected linked list to have length 2 got %d", ll.length)
	}
	if v != 7 {
		t.Errorf("expected poped element to have value 7 got %d", v)
	}
	if ll.tail.Value != 10 {
		t.Errorf("expected new tail to have value 10 got %d", ll.tail.Value)
	}
	// Pop 3
	v = ll.Pop()
	if ll.length != 1 {
		t.Errorf("expected linked list to have length 1 got %d", ll.length)
	}
	if v != 10 {
		t.Errorf("expected poped element to have value 10 got %d", v)
	}
	if ll.tail.Value != 5 {
		t.Errorf("expected new tail to have value 5 got %d", ll.tail.Value)
	}
	// Pop 4
	v = ll.Pop()
	if ll.length != 1 {
		t.Errorf("expected linked list to have length 1 got %d", ll.length)
	}
	if v != 5 {
		t.Errorf("expected poped element to have value 5 got %d", v)
	}
	if ll.tail.Value != 5 {
		t.Errorf("expected new tail to have value 5 got %d", ll.tail.Value)
	}
}

func TestSingleLinkedList_Prepend(t *testing.T) {
	ll := NewSingleLinkedList(5)
	ll.Prepend(7)
	if ll.length != 2 {
		t.Errorf("expected linked list to have length of 2 got %d", ll.length)
	}
	if ll.head.NextNode != ll.tail {
		t.Errorf("expected the prepended head to point to the now tail")
	}
	if ll.head.Value != 7 {
		t.Errorf("expected the head to have value 7 got %d", ll.head.Value)
	}
}

func TestSingleLinkedList_String(t *testing.T) {
	ll := NewSingleLinkedList(5)
	ll.Push(7)
	ll.Prepend(8)

	s := "8 --> 5 --> 7 -->"
	if s != ll.String() {
		t.Errorf("expected linked list string '%s' got '%s'", s, ll.String())
	}
}

func TestSingleLinkedList_Insert(t *testing.T) {
	ll := NewSingleLinkedList(5)
	ll.Push(7)
	ll.Prepend(8)
	ll.Push(12)
	ll.Insert(2, 16)
	s := "8 --> 5 --> 16 --> 7 --> 12 -->"
	if s != ll.String() {
		t.Errorf("expected linked list '%s' got '%s'", s, ll.String())
	}
}

func TestSingleLinkedList_Remove(t *testing.T) {
	ll := NewSingleLinkedList(5)
	ll.Push(7)
	ll.Push(8)
	ll.Push(12)

	ll.Remove(2)
	if ll.length != 3 {
		t.Errorf("expected the linked list to have a length of 3 got %d", ll.length)
	}
	s := "5 --> 7 --> 12 -->"
	if ll.String() != s {
		t.Errorf("expected linked list '%s' got '%s'", s, ll.String())
	}
}

func TestSingleLinkedList_Reverse(t *testing.T) {
	l := NewSingleLinkedList(5)
	l.Push(8)
	l.Push(7)
	l.Push(12)
	l.Push(3)
	l.Reverse()

	s := "3 --> 12 --> 7 --> 8 --> 5 -->"
	if l.String() != s {
		t.Errorf("expected a linked list %s got %s", s, l.String())
	}
}
