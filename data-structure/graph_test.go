package data_structure

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	g.AddVertex("0")
	g.AddVertex("1")
	g.AddVertex("2")
	g.AddVertex("3")
	g.AddVertex("4")
	g.AddVertex("5")
	g.AddVertex("6")
	g.AddEdge("3", "1", 0, 2)
	g.AddEdge("3", "4", 0, 2)
	g.AddEdge("4", "2", 0, 2)
	g.AddEdge("4", "5", 0, 2)
	g.AddEdge("1", "2", 0, 2)
	g.AddEdge("1", "0", 0, 2)
	g.AddEdge("0", "2", 0, 2)
	g.AddEdge("6", "5", 0, 2)
	fmt.Println(g.String())
}
