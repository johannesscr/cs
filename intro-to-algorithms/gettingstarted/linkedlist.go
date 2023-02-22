package gettingstarted

/* LINKED LISTS */

type ListNode struct {
	root bool
	NextNode *ListNode
	Data interface{}
}

func NewLinkedList(data interface{}) *ListNode {
	return &ListNode{
		root: true,
		Data: data,
	}
}

func (ln *ListNode) InsertAfter(data interface{}) {
	if ln.NextNode != nil {

	}
	ln.NextNode = &ListNode{
		Data: data,
	}
}