package main

import (
	"fmt"
	"sort"
)

type MarsGraph struct {
	adjMatrix [][]bool
}

func createMarsGraph(vertexCount int) *MarsGraph {
	graph := &MarsGraph{
		adjMatrix: make([][]bool, vertexCount),
	}

	for i := 0; i < vertexCount; i++ {
		graph.adjMatrix[i] = make([]bool, vertexCount)
	}

	return graph
}

func (graph *MarsGraph) addEdge(from, to int) {
	graph.adjMatrix[from][to] = true
	graph.adjMatrix[to][from] = true
}

func (graph *MarsGraph) exploreConnectedComponents() ([]int, []int) {
	vertexCount := len(graph.adjMatrix)
	visited := make([]bool, vertexCount)
	group := make([]int, vertexCount)
	for i := range group {
		group[i] = -1
	}

	var dfs func(v, g int) bool
	dfs = func(v, g int) bool {
		visited[v] = true
		group[v] = g

		for i := 0; i < vertexCount; i++ {
			if graph.adjMatrix[v][i] {
				if group[i] == g {
					return false
				}
				if group[i] == -1 && !dfs(i, 1-g) {
					return false
				}
			}
		}

		return true
	}

	for i := 0; i < vertexCount; i++ {
		if !visited[i] && !dfs(i, 0) {
			return nil, nil
		}
	}

	leftGroup := make([]int, 0)
	rightGroup := make([]int, 0)
	for i, g := range group {
		if g == 0 {
			leftGroup = append(leftGroup, i)
		} else {
			rightGroup = append(rightGroup, i)
		}
	}

	return leftGroup, rightGroup
}

func main() {
	var vertexCount int
	fmt.Scan(&vertexCount)

	graph := createMarsGraph(vertexCount)

	for i := 0; i < vertexCount; i++ {
		for j := 0; j < vertexCount; j++ {
			var symbol string
			fmt.Scan(&symbol)
			if symbol == "+" {
				graph.addEdge(i, j)
			}
		}
	}

	leftGroup, rightGroup := graph.exploreConnectedComponents()

	if leftGroup == nil || rightGroup == nil {
		fmt.Println("No solution")
		return
	}

	if len(leftGroup) < len(rightGroup) {
		sort.Ints(leftGroup)
		for _, v := range leftGroup {
			fmt.Print(v+1, " ")
		}
	} else {
		sort.Ints(rightGroup)
		for _, v := range rightGroup {
			fmt.Print(v+1, " ")
		}
	}
}
