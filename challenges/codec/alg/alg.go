package alg

import (
	"fmt"
)

// robotDirectionAfter
func robotDirectionAfter(distance float32) string {
	directions := []string{"N", "E", "S", "W"}
	totalDistance := (distance * 10 / 2) / 10
	var directionChanges float32 = 0
	var travelDistance float32 = 0
	for i := float32(0); i <= totalDistance; i++ {
		fmt.Println(i, travelDistance, totalDistance, directions[int(directionChanges)%4])
		if travelDistance >= totalDistance {
			break
		}
		directionChanges = i
		travelDistance += i
	}
	//for travelDistance != totalDistance {
	//	travelDistance += directionChanges
	//	directionChanges += 1
	//}

	return directions[int(directionChanges)%4]
}

func isMine(x string) int {
	if x == "X" {
		return 1
	}
	return 0
}

//func mineIndex(M, i, j int) int {
//	if i
//	return 0
//}

//x 0 0
//0 x 0
//x x 0

func minesweeper(M []string) string {
	n := len(M)    // number of row
	m := len(M[0]) // number of columns
	matrix := ""
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			rowTop := M[i-1]
			p := isMine(rowTop[j-1 : j])
			p += isMine(rowTop[j : j+1])
			p += isMine(rowTop[j+1 : j+2])
			row := M[i]
			p += isMine(row[j-1 : j])
			p += isMine(row[j+1 : j+2])
			rowBottom := M[i+1]
			p += isMine(rowBottom[j-1 : j])
			p += isMine(rowBottom[j : j+1])
			p += isMine(rowBottom[j+1 : j+2])

			ps := fmt.Sprintf("%d", p)
			matrix += ps
		}
	}
	return matrix
}
