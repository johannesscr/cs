package data_structure

import (
	"fmt"
	"log"
)

const (
	FROM          = iota
	BIDIRECTIONAL = iota
	TO            = iota
)

type Vertex struct {
	Name string
}

// Edge is a struct that represents the edge of a graph.
// Direction has three possible options FROM = 1, BIDIRECTIONAL = 2 and TO = 3.
type Edge struct {
	Weight    int
	Direction int
	V1        *Vertex
	V2        *Vertex
}

func NewEdge(v1, v2 *Vertex, weight int, direction int) *Edge {
	d := 2
	switch direction {
	case FROM:
		d = 1
	case BIDIRECTIONAL:
		d = 2
	case TO:
		d = 3
	}
	return &Edge{
		Weight:    weight,
		Direction: d,
		V1:        v1,
		V2:        v2,
	}
}

type Graph struct {
	NumberOfNodes int
	Vertices      map[string]*Vertex
	AdjacencyList map[string][]*Edge
}

func NewGraph() Graph {
	return Graph{
		NumberOfNodes: 0,
		Vertices:      make(map[string]*Vertex),
		AdjacencyList: make(map[string][]*Edge),
	}
}

func (g *Graph) AddVertex(name string) {
	v := &Vertex{Name: name}
	g.NumberOfNodes += 1
	g.Vertices[name] = v
	g.AdjacencyList[name] = make([]*Edge, 0)
	fmt.Println(g.Vertices)
}

func (g *Graph) AddEdge(v1Name, v2Name string, weight, direction int) {
	v1, ok := g.Vertices[v1Name]
	if !ok {
		log.Fatalln("vertex/node not found", v1Name)
	}
	v2, ok := g.Vertices[v2Name]
	if !ok {
		log.Fatalln("vertex/node not found", v2Name)
	}
	edge := NewEdge(v1, v2, weight, direction)
	g.AdjacencyList[v1Name] = append(g.AdjacencyList[v1Name], edge)
	g.AdjacencyList[v2Name] = append(g.AdjacencyList[v2Name], edge)
}

func (g *Graph) String() string {
	s := ""
	for name := range g.Vertices {
		edges := g.AdjacencyList[name]
		si := ""
		for _, edge := range edges {
			if edge.V1.Name == name {
				si += fmt.Sprintf("[%s %d] ", edge.V2.Name, edge.Weight)
			} else {
				si += fmt.Sprintf("[%s %d] ", edge.V1.Name, edge.Weight)
			}
		}
		s += fmt.Sprintf("%-5s --> %s\n", name, si)
	}
	return s
}
