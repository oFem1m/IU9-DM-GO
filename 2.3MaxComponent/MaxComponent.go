package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph struct {
	vertices      int
	adjacencyList map[int][]int
	component     []int
	visited       []bool
	componentID   int
	maxSize       int
	maxEdges      int
	minVertex     int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices:      vertices,
		adjacencyList: make(map[int][]int),
		component:     make([]int, vertices),
		visited:       make([]bool, vertices),
		componentID:   0,
		maxSize:       0,
		maxEdges:      0,
		minVertex:     vertices - 1,
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u)
}

func (g *Graph) DFS(v int) int {
	g.visited[v] = true
	g.component[v] = g.componentID
	size := 1

	for _, neighbor := range g.adjacencyList[v] {
		if !g.visited[neighbor] {
			size += g.DFS(neighbor)
		}
	}

	return size
}

func (g *Graph) FindLargestComponent() {
	for v := 0; v < g.vertices; v++ {
		if !g.visited[v] {
			g.componentID++
			size := g.DFS(v)
			edges := 0

			for i := 0; i < g.vertices; i++ {
				if g.component[i] == g.component[v] {
					edges += len(g.adjacencyList[i])
				}
			}

			if size > g.maxSize || (size == g.maxSize && edges > g.maxEdges) {
				g.maxSize = size
				g.maxEdges = edges
				g.minVertex = v
			}
		}
	}
}

func (g *Graph) PrintGraph() {
	fmt.Println("graph {")

	for v := 0; v < g.vertices; v++ {
		if g.component[v] == g.component[g.minVertex] {
			fmt.Printf("  %d [color = red]\n", v)
		} else {
			fmt.Printf("  %d\n", v)
		}
	}

	for v := 0; v < g.vertices; v++ {
		for _, neighbor := range g.adjacencyList[v] {
			if neighbor >= v {
				if g.component[v] == g.component[g.minVertex] && g.component[neighbor] == g.component[g.minVertex] {
					fmt.Printf("  %d -- %d [color = red]\n", v, neighbor)
				} else {
					fmt.Printf("  %d -- %d\n", v, neighbor)
				}
			}
		}
	}

	fmt.Println("}")
}

func parseInput(input string) int {
	var num int
	fmt.Sscanf(input, "%d", &num)
	return num
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	vertices := parseInput(scanner.Text())
	scanner.Scan()
	edges := parseInput(scanner.Text())

	graph := NewGraph(vertices)

	for i := 0; i < edges; i++ {
		scanner.Scan()
		edgeData := strings.Split(scanner.Text(), " ")
		u := parseInput(edgeData[0])
		v := parseInput(edgeData[1])
		graph.AddEdge(u, v)
	}
	graph.FindLargestComponent()

	graph.PrintGraph()
}
