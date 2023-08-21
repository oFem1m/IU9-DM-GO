package main

import (
	"fmt"
	"sort"
)

type Vertex struct {
	IngoingEdges, Edges, Bucket                       []*Vertex
	Ancestor, Label, SemiDominator, Dominator, Parent *Vertex
	Visited, HasOperand                               bool
	Depth, OperandValue                               int
}

type Graph struct {
	Vertices []*Vertex
}

func NewVertex() *Vertex {
	vertex := &Vertex{}
	vertex.SemiDominator = vertex
	vertex.Label = vertex
	return vertex
}

type VerticesAreSorted []*Vertex

func (graph VerticesAreSorted) Len() int {
	return len(graph)
}

func (graph VerticesAreSorted) Less(i, j int) bool {
	a, b := graph[i].Depth, graph[j].Depth
	return a < b
}

func (graph VerticesAreSorted) Swap(i, j int) {
	graph[i], graph[j] = graph[j], graph[i]
}

func visitVerticesDFS(start *Vertex) {
	depth := 0

	var DFSInner func(*Vertex)
	DFSInner = func(vertex *Vertex) {
		vertex.Depth = depth
		vertex.Visited = true
		depth++

		for _, u := range vertex.Edges {
			if !u.Visited {
				u.Parent = vertex
				DFSInner(u)
			}
		}
	}
	DFSInner(start)
}

func filterVertices(graph *Graph) {
	visitVerticesDFS(graph.Vertices[0])
	var filteredGraph []*Vertex
	for _, vertex := range graph.Vertices {
		if vertex.Visited {
			filteredGraph = append(filteredGraph, vertex)
		}
	}
	graph.Vertices = filteredGraph

	for _, vertex := range graph.Vertices {
		var filteredIngoingEdges []*Vertex
		for _, u := range vertex.IngoingEdges {
			if u.Visited {
				filteredIngoingEdges = append(filteredIngoingEdges, u)
			}
		}
		vertex.IngoingEdges = filteredIngoingEdges
	}
}

func dominators(graph *Graph) {
	sort.Sort(VerticesAreSorted(graph.Vertices))
	for i := len(graph.Vertices) - 1; i > 0; i-- {
		vertex := graph.Vertices[i]
		for _, u := range vertex.IngoingEdges {
			minVertex := findMin(u)
			if minVertex.SemiDominator.Depth < vertex.SemiDominator.Depth {
				vertex.SemiDominator = minVertex.SemiDominator
			}
		}
		vertex.Ancestor = vertex.Parent
		vertex.SemiDominator.Bucket = append(vertex.SemiDominator.Bucket, vertex)
		for _, u := range vertex.Parent.Bucket {
			minVertex := findMin(u)
			if minVertex.SemiDominator == u.SemiDominator {
				u.Dominator = vertex.Parent
			} else {
				u.Dominator = minVertex
			}
		}
		vertex.Parent.Bucket = nil
	}

	for i := 1; i < len(graph.Vertices); i++ {
		if graph.Vertices[i].Dominator != graph.Vertices[i].SemiDominator {
			graph.Vertices[i].Dominator = graph.Vertices[i].Dominator.Dominator
		}
	}

	graph.Vertices[0].Dominator = nil
}

func findMin(vertex *Vertex) *Vertex {
	searchAndCut(vertex)
	return vertex.Label
}

func searchAndCut(vertex *Vertex) *Vertex {
	if vertex.Ancestor == nil {
		return vertex
	}
	root := searchAndCut(vertex.Ancestor)
	if vertex.Ancestor.Label.SemiDominator.Depth < vertex.Label.SemiDominator.Depth {
		vertex.Label = vertex.Ancestor.Label
	}
	vertex.Ancestor = root
	return root
}

func getLoopsCount(graph *Graph) int {
	loopsCount := 0
	for _, vertex := range graph.Vertices {
		for _, u := range vertex.IngoingEdges {
			for u != vertex && u != nil {
				u = u.Dominator
			}
			if u == vertex {
				loopsCount++
				break
			}
		}
	}
	return loopsCount
}

func linkVertex(graph *Graph, index int) {
	if index < len(graph.Vertices)-1 {
		graph.Vertices[index].Edges = append(graph.Vertices[index].Edges, graph.Vertices[index+1])
		graph.Vertices[index+1].IngoingEdges = append(graph.Vertices[index+1].IngoingEdges, graph.Vertices[index])
	}
}

func linkOperands(graph *Graph, indexToVertex map[int]*Vertex) {
	for _, vertex := range graph.Vertices {
		if !vertex.HasOperand {
			continue
		}
		vertex.Edges = append(vertex.Edges, indexToVertex[vertex.OperandValue])
		indexToVertex[vertex.OperandValue].IngoingEdges = append(indexToVertex[vertex.OperandValue].IngoingEdges, vertex)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	graph := &Graph{
		Vertices: make([]*Vertex, n),
	}
	for i := 0; i < n; i++ {
		graph.Vertices[i] = NewVertex()
	}

	indexToVertex := map[int]*Vertex{}
	for i := 0; i < n; i++ {
		var index int
		var command string
		fmt.Scan(&index, &command)
		indexToVertex[index] = graph.Vertices[i]
		if command == "ACTION" {
			linkVertex(graph, i)
		} else {
			var operand int
			fmt.Scan(&operand)
			graph.Vertices[i].HasOperand = true
			graph.Vertices[i].OperandValue = operand
			if command == "BRANCH" {
				linkVertex(graph, i)
			}
		}
	}

	linkOperands(graph, indexToVertex)
	filterVertices(graph)
	dominators(graph)
	fmt.Println(getLoopsCount(graph))
}
