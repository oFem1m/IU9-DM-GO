package main

import (
	"fmt"
	"sort"
)

func addEdge(graph map[int][]int, u, v int) {
	graph[u] = append(graph[u], v)
}

func dfs(graph map[int][]int, v int, visited []bool, stack *[]int) {
	visited[v] = true

	for _, i := range graph[v] {
		if !visited[i] {
			dfs(graph, i, visited, stack)
		}
	}

	*stack = append(*stack, v)
}

func dfsStrong(graph map[int][]int, v int, visited []bool, component *[]int) {
	visited[v] = true
	*component = append(*component, v)

	for _, i := range graph[v] {
		if !visited[i] {
			dfsStrong(graph, i, visited, component)
		}
	}
}

func transpose(graph map[int][]int) map[int][]int {
	transposedGraph := make(map[int][]int)

	for u, neighbors := range graph {
		for _, v := range neighbors {
			addEdge(transposedGraph, v, u)
		}
	}

	return transposedGraph
}

func findStrongComponents(graph map[int][]int, n int) [][]int {
	visited := make([]bool, n)
	stack := make([]int, 0)
	components := make([][]int, 0)

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(graph, i, visited, &stack)
		}
	}

	transposedGraph := transpose(graph)

	for i := 0; i < n; i++ {
		visited[i] = false
	}

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[v] {
			component := make([]int, 0)
			dfsStrong(transposedGraph, v, visited, &component)
			sort.Ints(component)
			components = append(components, component)
		}
	}

	return components
}

func buildCondensedGraph(graph map[int][]int, strongComponents [][]int) map[int][]int {
	condensedGraph := make(map[int][]int)
	condensedEdges := make(map[int]map[int]bool)

	componentsMapping := make(map[int]int)
	for index, component := range strongComponents {
		for _, node := range component {
			componentsMapping[node] = index
		}
	}

	for node, neighbors := range graph {
		component := componentsMapping[node]
		for _, neighbor := range neighbors {
			neighborComponent := componentsMapping[neighbor]
			if component != neighborComponent {
				if _, exists := condensedEdges[component]; !exists {
					condensedEdges[component] = make(map[int]bool)
				}
				if !condensedEdges[component][neighborComponent] {
					condensedGraph[component] = append(condensedGraph[component], neighborComponent)
					condensedEdges[component][neighborComponent] = true
				}
			}
		}
	}

	for _, component := range strongComponents {
		componentIndex := componentsMapping[component[0]]
		if _, exists := condensedGraph[componentIndex]; !exists {
			condensedGraph[componentIndex] = []int{}
		}
	}

	return condensedGraph
}

func findCondensationBase(condensedGraph map[int][]int) []int {
	base := make([]int, 0)
	incomingEdges := make(map[int]bool)

	for _, neighbors := range condensedGraph {
		for _, neighbor := range neighbors {
			incomingEdges[neighbor] = true
		}
	}

	for v := range condensedGraph {
		if !incomingEdges[v] {
			base = append(base, v)
		}
	}

	return base
}

func main() {
	graph := make(map[int][]int)
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		addEdge(graph, u, v)
	}

	strongComponents := findStrongComponents(graph, n)
	condensedGraph := buildCondensedGraph(graph, strongComponents)
	base := findCondensationBase(condensedGraph)

	var result []int
	for i := 0; i < len(base); i++ {
		result = append(result, strongComponents[base[i]][0])
	}
	sort.Ints(result)
	for _, i := range result {
		fmt.Print(i, " ")
	}
}
