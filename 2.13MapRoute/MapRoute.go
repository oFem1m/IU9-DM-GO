package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	x, y int
	dist int
}

type PQueue []*Node

func (pq PQueue) Len() int {
	return len(pq)
}

func (pq PQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[:n-1]
	return node
}

func minPathLength(N int, matrix [][]int) int {
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	distance := make([][]int, N)
	for i := range distance {
		distance[i] = make([]int, N)
		for j := range distance[i] {
			distance[i][j] = 1e9
		}
	}

	distance[0][0] = matrix[0][0]

	pq := make(PQueue, 0)
	heap.Push(&pq, &Node{0, 0, matrix[0][0]})

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		x, y, dist := node.x, node.y, node.dist

		if dist > distance[x][y] {
			continue
		}

		for k := 0; k < 4; k++ {
			nx, ny := x+dx[k], y+dy[k]
			if nx >= 0 && nx < N && ny >= 0 && ny < N {
				newDist := dist + matrix[nx][ny]
				if newDist < distance[nx][ny] {
					distance[nx][ny] = newDist
					heap.Push(&pq, &Node{nx, ny, newDist})
				}
			}
		}
	}

	return distance[N-1][N-1]
}

func main() {
	var N int
	fmt.Scan(&N)

	matrix := make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, N)
		for j := range matrix[i] {
			fmt.Scan(&matrix[i][j])
		}
	}

	minLength := minPathLength(N, matrix)
	fmt.Println(minLength)
}
