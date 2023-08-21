package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type Vertex struct {
	adjacentVertices []int
	partition        int
}

type Solution struct {
	leftPartition  []int
	rightPartition []int
}

type MarsGraph struct {
	isConnected bool
	vertices    []Vertex
}

func createMarsGraph(vertexCount int) *MarsGraph {
	graph := &MarsGraph{
		isConnected: false,
		vertices:    make([]Vertex, vertexCount),
	}
	return graph
}

func (graph *MarsGraph) addEdge(from, to int) {
	graph.vertices[from].adjacentVertices = append(graph.vertices[from].adjacentVertices, to)
	graph.vertices[to].adjacentVertices = append(graph.vertices[to].adjacentVertices, from)
	graph.isConnected = true
}

func (graph *MarsGraph) dfs(v, partition int, leftPartition, rightPartition *[]int) bool {
	graph.vertices[v].partition = partition
	if partition == 1 {
		*leftPartition = append(*leftPartition, v)
	} else {
		*rightPartition = append(*rightPartition, v)
	}

	for _, i := range graph.vertices[v].adjacentVertices {
		if graph.vertices[i].partition == partition {
			return false
		}
		if graph.vertices[i].partition == 0 && !graph.dfs(i, -partition, leftPartition, rightPartition) {
			return false
		}
	}

	return true
}

func (graph *MarsGraph) exploreConnectedComponents() []Solution {
	solutions := make([]Solution, 0)
	for v := range graph.vertices {
		if graph.vertices[v].partition != 0 {
			continue
		}

		leftPartition := make([]int, 0)
		rightPartition := make([]int, 0)
		if !graph.dfs(v, 1, &leftPartition, &rightPartition) {
			fmt.Println("No solution")
			os.Exit(0)
		}
		solution := Solution{leftPartition: leftPartition, rightPartition: rightPartition}
		solutions = append(solutions, solution)
	}
	return solutions
}

func getGroups(solutions *[]Solution, groups *[][]int, depth, currentDepth int, currentGroup []int) {
	if currentDepth == depth+1 {
		*groups = append(*groups, append([]int{}, currentGroup...))
		return
	}

	getGroups(solutions, groups, depth, currentDepth+1, append(append([]int{}, currentGroup...), (*solutions)[currentDepth].leftPartition...))
	getGroups(solutions, groups, depth, currentDepth+1, append(append([]int{}, currentGroup...), (*solutions)[currentDepth].rightPartition...))
}

func lexicalComparator(A []int, B []int) bool {
	for i := 0; i < len(A); i++ {
		if A[i] < B[i] {
			return true
		} else if A[i] > B[i] {
			return false
		}
	}
	return false
}

func findBestGroup(groups [][]int, vertexCount int) []int {
	minDifference := math.MaxInt
	bestGroup := groups[0]
	for i := 0; i < len(groups); i++ {
		difference := int(math.Abs(float64(len(groups[i]) - (vertexCount - len(groups[i])))))
		switch {
		case minDifference > difference:
			bestGroup = groups[i]
			minDifference = difference
		case minDifference == difference:
			if len(bestGroup) > len(groups[i]) || len(bestGroup) == len(groups[i]) && lexicalComparator(groups[i], bestGroup) {
				bestGroup = groups[i]
			}
		}
	}
	return bestGroup
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

	if !graph.isConnected {
		for i := 0; i < vertexCount/2; i++ {
			fmt.Print(i+1, " ")
		}
		return
	}

	solutions := graph.exploreConnectedComponents()

	groups := make([][]int, 0)
	getGroups(&solutions, &groups, len(solutions)-1, 0, make([]int, 0))
	for i := 0; i < len(groups); i++ {
		sort.Ints(groups[i])
	}

	bestGroup := findBestGroup(groups, vertexCount)

	for i := 0; i < len(bestGroup); i++ {
		fmt.Print(bestGroup[i]+1, " ")
	}
}
