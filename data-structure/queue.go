package data_structure

import "fmt"

// QueueNode represents each element within the Queue.
type QueueNode struct {
	value interface{}
	next  *QueueNode
}

// NewQueueNode returns the value of a new node pointed to, with value as passed
// as the argument.
func NewQueueNode(value interface{}) *QueueNode {
	return &QueueNode{
		value: value,
		next:  nil,
	}
}

// Queue is the data structure that represents a queue, with the start being
// the beginning of the queue and end being the newest element.
// First In: goes to end.
// Last out: start of the queue.
type Queue struct {
	start  *QueueNode
	end    *QueueNode
	length int
}

// NewQueue creates a new Queue, with an initial node being from the value
// passed as the argument.
func NewQueue(value interface{}) Queue {
	node := NewQueueNode(value)
	return Queue{
		start:  node,
		end:    node,
		length: 1,
	}
}

// Enqueue adds a new QueueNode to the end of the Queue.
func (q *Queue) Enqueue(value interface{}) {
	node := NewQueueNode(value)
	q.end.next = node
	q.end = node
	q.length += 1
}

// Dequeue removes a QueueNode to the start of the Queue and returns that
// QueueNode.
func (q *Queue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.start
	q.start = n.next
	q.length -= 1
	return n.value
}

// Peek returns the value at the start of the Queue.
func (q *Queue) Peek() interface{} {
	return q.start.value
}

// String converts and returns the Queue into a string representation of
// the queue.
func (q *Queue) String() string {
	node := q.start
	s := fmt.Sprintf("in -> %v -> ", node.value)
	for {
		node = node.next
		if node == nil {
			break
		}
		s += fmt.Sprintf("%v -> ", node.value)
	}
	s += "out"
	return s
}
