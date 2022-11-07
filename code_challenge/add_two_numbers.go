package code_challenge

/*
You are given two non-empty linked lists representing two non-negative
integers. The digits are stored in reverse order, and each of their nodes
contains a single digit. Add the two numbers and return the sum as a linked
list.

You may assume the two numbers do not contain any leading zero, except the
number 0 itself.
*/

/*
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode is a node for the single linked list.
// The start of the linked list is simply the first node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// EXAMPLE
//              +1        +1
//   234        234        234        234
// +  87  ->  +  87  ->  +  87  ->  +  87
// -----      -----      -----      -----
//                1         21        321

// addTwoNumbers is how we would do classic addition when adding a list of
// numbers on top of one another.
//
// where the column of digits is added up, all the increments of 10 are
// carried forward to the next column
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return &ListNode{}
}

/*
PLANNING
the linked list if given in reverse order, therefore as we would naturally do
addition. The function should:
- step through each node adding all the two numbers.
- any multiple of 10 should be carried forward and done again.
- we don't know how long the list is, so it should recurse on itself.
*/

func addNode(l1 *ListNode, l2 *ListNode, carriedForward int) *ListNode {
	// the total of the two nodes and carried forward
	t := l1.Val + l2.Val + carriedForward
	// nothing to add
	// base condition
	// the single digits
	v := t % 10
	// the number of 10's carried forward
	c := (t - v) / 10
	n := &ListNode{
		Val: v,
	}
	if t == 0 || (l1 == nil && l2 == nil) {
		return n
	}

}
