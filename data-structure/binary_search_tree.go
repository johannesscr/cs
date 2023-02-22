package data_structure

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
