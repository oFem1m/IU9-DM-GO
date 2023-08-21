package main

import (
	"fmt"
)

type Graph struct {
	Vertices  int
	AdjList   map[int][]int
	Visited   []bool
	Discovery []int
	Low       []int
	Parent    []int
	Bridges   [][]int
	Time      int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices:  vertices,
		AdjList:   make(map[int][]int),
		Visited:   make([]bool, vertices),
		Discovery: make([]int, vertices),
		Low:       make([]int, vertices),
		Parent:    make([]int, vertices),
		Bridges:   [][]int{},
		Time:      0,
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.AdjList[u] = append(g.AdjList[u], v)
	g.AdjList[v] = append(g.AdjList[v], u)
}

func (g *Graph) FindBridges() {
	for i := 0; i < g.Vertices; i++ {
		if !g.Visited[i] {
			g.DFS(i)
		}
	}
}

func (g *Graph) DFS(u int) {
	g.Visited[u] = true
	g.Time++
	g.Discovery[u] = g.Time
	g.Low[u] = g.Time

	for _, v := range g.AdjList[u] {
		if !g.Visited[v] {
			g.Parent[v] = u
			g.DFS(v)

			g.Low[u] = min(g.Low[u], g.Low[v])

			if g.Low[v] > g.Discovery[u] {
				bridge := []int{u, v}
				g.Bridges = append(g.Bridges, bridge)
			}
		} else if v != g.Parent[u] {
			g.Low[u] = min(g.Low[u], g.Discovery[v])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var vertices, edges int
	fmt.Scan(&vertices)
	fmt.Scan(&edges)

	graph := NewGraph(vertices)

	for i := 0; i < edges; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		graph.AddEdge(u, v)
	}

	graph.FindBridges()

	fmt.Println(len(graph.Bridges))
}
