package algorithms

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree_BreadthFirstSearch(t *testing.T) {
	b := NewBinarySearchTree(9)
	b.Insert(4)
	b.Insert(6)
	b.Insert(20)
	b.Insert(170)
	b.Insert(15)
	b.Insert(1)

	//       9
	//   4       20
	// 1   6   15   170

	breadthTree := b.BreadthFirstSearch(6)
	fmt.Println(breadthTree)
}

func TestBinarySearchTree_BreadthFirstSearchRecursive(t *testing.T) {
	b := NewBinarySearchTree(9)
	b.Insert(4)
	b.Insert(6)
	b.Insert(20)
	b.Insert(170)
	b.Insert(15)
	b.Insert(1)

	//       9
	//   4       20
	// 1   6   15   170

	list := make([]int, 0)
	queue := []*BinarySearchNode{b.Root}
	breadthTree := b.BreadthFirstSearchRecursive(6, queue, list)
	fmt.Println(breadthTree)
}

func TestBinarySearchTree_DepthFirstSearchInOrder(t *testing.T) {
	b := NewBinarySearchTree(9)
	b.Insert(4)
	b.Insert(6)
	b.Insert(20)
	b.Insert(170)
	b.Insert(15)
	b.Insert(1)

	//       9
	//   4       20
	// 1   6   15   170

	depthTree := b.DepthFirstSearchInOrder(6)
	fmt.Println(depthTree)
}

func TestBinarySearchTree_DepthFirstSearchPreOrder(t *testing.T) {
	b := NewBinarySearchTree(9)
	b.Insert(4)
	b.Insert(6)
	b.Insert(20)
	b.Insert(170)
	b.Insert(15)
	b.Insert(1)

	//       9
	//   4       20
	// 1   6   15   170

	depthTree := b.DepthFirstSearchPreOrder(6)
	fmt.Println(depthTree)
}

func TestBinarySearchTree_DepthFirstSearchPostOrder(t *testing.T) {
	b := NewBinarySearchTree(9)
	b.Insert(4)
	b.Insert(6)
	b.Insert(20)
	b.Insert(170)
	b.Insert(15)
	b.Insert(1)

	//       9
	//   4       20
	// 1   6   15   170

	depthTree := b.DepthFirstSearchPostOrder(6)
	fmt.Println(depthTree)
}
