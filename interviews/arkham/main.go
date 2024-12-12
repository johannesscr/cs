package main

import (
	"fmt"
	"strings"
)

// To execute Go code, please declare a func main() in a package "main"

/*
# Designer
# -UX Designer
# Engineer
# -Mechanical Engineer
# -Software Engineer
# --Front-End Engineer


positions = [{"id": 1, "name": "Engineer", "parent_id":None},
              {"id": 2, "name": "Front-End Engineer", "parent_id":3},
              {"id": 3, "name": "Software Engineer", "parent_id":1},
              {"id": 4, "name": "Mechanical Engineer", "parent_id":1},
              {"id": 5, "name": "UX Designer", "parent_id":6},
              {"id": 6, "name": "Designer", "parent_id":None}]

print(positions)
*/

type Position struct {
	id        int
	name      string
	parent_id int
}

var positions = []Position{
	{
		id:        1,
		name:      "Engineer",
		parent_id: 0,
	},
	{
		id:        2,
		name:      "Front-End Engineer",
		parent_id: 3,
	}, {
		id:        3,
		name:      "Software Engineer",
		parent_id: 1,
	}, {
		id:        4,
		name:      "Mechanical Engineer",
		parent_id: 1,
	}, {
		id:        5,
		name:      "UX Designer",
		parent_id: 6,
	}, {
		id:        6,
		name:      "Designer",
		parent_id: 0,
	},
}

var positionMap = map[int][]Position{}

type Children struct {
	ParentId int
	Children []Position
}

func FindChildren(position Position) []Position {
	children := make([]Position, 0)
	for _, child := range positions {
		if child.parent_id == position.id {
			children = append(children, child)
		}
	}
	return children
}

// sort.Slice(planets, func(i, j int) bool {
//   return planets[i].Axis < planets[j].Axis
// })

func PrintPositions(position Position, level int) {
	fmt.Printf("%s%s\n", strings.Repeat("-", level), position.name)
	// children := FindChildren(position)
	if children, ok := positionMap[position.id]; ok {

		for _, subPosition := range children {
			PrintPositions(subPosition, level+1)
		}
	}
}

func main() {
	for _, position := range positions {
		if v, ok := positionMap[position.parent_id]; ok {
			positionMap[position.parent_id] = append(v, position)
		} else {
			positionMap[position.parent_id] = []Position{position}
		}
	}
	startingPositions := positionMap[0]
	for _, position := range startingPositions {
		PrintPositions(position, 0)
	}
}
