package algorithms

/* COPY FROM DATA STRUCTURE */

type BinarySearchNode struct {
	Value int               `json:"value"`
	Left  *BinarySearchNode `json:"left"`
	Right *BinarySearchNode `json:"right"`
}

func NewBinarySearchNode(value int) *BinarySearchNode {
	return &BinarySearchNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

// BinarySearchTree is a collection of nodes, starting at the root. It has a
// check method that based of the boolean output inserts a node to the left
// (false) or to the right (true).
type BinarySearchTree struct {
	Root *BinarySearchNode `json:"root"`
}

// NewBinarySearchTree create a new binary sear
func NewBinarySearchTree(value int) BinarySearchTree {
	root := NewBinarySearchNode(value)
	return BinarySearchTree{
		Root: root,
	}
}

// traverse is used to traverse through the tree, it either returns the node
// with the same value as the search for the traverse or the nearest parent node
// is returned.
func (b *BinarySearchTree) traverse(value int) *BinarySearchNode {
	node := b.Root
	for {
		// if the node is found
		if node.Value == value {
			break
		}
		if value < node.Value {
			// traverse left
			if node.Left == nil {
				break
			}
			node = node.Left
		}
		if value > node.Value {
			// traverse right
			if node.Right == nil {
				break
			}
			node = node.Right
		}
	}
	return node
}

// traverseToParent finds the parent node for the node with value given or nil.
//func (b *BinarySearchTree) traverseToParent(value int) *BinarySearchNode {
//	node := b.Root
//
//	return nil
//}

// Insert will insert the value into the BinarySearchTree as a new node.
func (b *BinarySearchTree) Insert(value int) {
	newNode := NewBinarySearchNode(value)
	node := b.traverse(value)
	if value < node.Value {
		node.Left = newNode
	} else {
		node.Right = newNode
	}
}

// Lookup returns the node pointed to for a value if it exists, otherwise nil.
func (b *BinarySearchTree) Lookup(value int) *BinarySearchNode {
	node := b.traverse(value)
	if node.Value == value {
		return node
	}
	return nil
}

func (b *BinarySearchTree) Remove() *BinarySearchNode {
	return nil
}

// BreadthFirstSearch is used to traverse on each level first
func (b *BinarySearchTree) BreadthFirstSearch(v int) []int {
	currentNode := b.Root
	// list of how we traverse, in order of BFS
	list := make([]int, 0)
	// use slice to keep track of level we are at, so that we can access
	// the children.
	// add all the children to the queue, then if we do not find an answer in
	// the list, we return the queue and continue to search from there.
	queue := make([]*BinarySearchNode, 0)
	// the queue can get pretty large if the queue is very wide and use
	// quite a bit of memory.

	queue = append(queue, currentNode)

	for len(queue) > 0 {
		currentNode = queue[0]
		//fmt.Println(currentNode.Value)
		queue = queue[1:] // remove the first node from the queue
		list = append(list, currentNode.Value)
		if currentNode.Value == v {
			break
		}
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}
	return list
}

// BreadthFirstSearchRecursive is used to traverse on each level first using
// a recursive method.
func (b *BinarySearchTree) BreadthFirstSearchRecursive(v int, queue []*BinarySearchNode, list []int) []int {
	// check for the base case
	if len(queue) == 0 {
		return list
	}
	currentNode := queue[0]
	queue = queue[1:]
	list = append(list, currentNode.Value)
	if currentNode.Value == v {
		return list
	}
	if currentNode.Left != nil {
		queue = append(queue, currentNode.Left)
	}
	if currentNode.Right != nil {
		queue = append(queue, currentNode.Right)
	}
	return b.BreadthFirstSearchRecursive(v, queue, list)
}

func (b *BinarySearchTree) TraversInOrder(node *BinarySearchNode, list []int) []int {
	if node.Left != nil {
		list = b.TraversInOrder(node.Left, list)
	}
	// if we cannot go any further down the left node, append the node
	// to the list.
	list = append(list, node.Value)
	if node.Right != nil {
		list = b.TraversInOrder(node.Right, list)
	}
	return list
}

// DepthFirstSearchInOrder the depth first search using the in-order
// implementation.
func (b *BinarySearchTree) DepthFirstSearchInOrder(v int) []int {
	// just starts at the root node.
	return b.TraversInOrder(b.Root, make([]int, 0))
}

func (b *BinarySearchTree) TraversePreOrder(node *BinarySearchNode, list []int) []int {
	// if we cannot go any further down the left node, append the node
	// to the list.
	list = append(list, node.Value)
	if node.Left != nil {
		list = b.TraversePreOrder(node.Left, list)
	}
	if node.Right != nil {
		list = b.TraversePreOrder(node.Right, list)
	}
	return list
}

// DepthFirstSearchPreOrder the depth first search using the pre-order
// implementation.
func (b *BinarySearchTree) DepthFirstSearchPreOrder(v int) []int {
	return b.TraversePreOrder(b.Root, make([]int, 0))
}

func (b *BinarySearchTree) TraversePostOrder(node *BinarySearchNode, list []int) []int {
	if node.Left != nil {
		list = b.TraversePostOrder(node.Left, list)
	}
	if node.Right != nil {
		list = b.TraversePostOrder(node.Right, list)
	}
	// if we cannot go any further down the left node, append the node
	// to the list.
	list = append(list, node.Value)
	return list
}

// DepthFirstSearchPostOrder the depth first search using the post-order
// implementation.
func (b *BinarySearchTree) DepthFirstSearchPostOrder(v int) []int {
	return b.TraversePostOrder(b.Root, make([]int, 0))
}
