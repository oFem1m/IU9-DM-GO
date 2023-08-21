package main

import "fmt"

func bfs(graph [][]int, pivot int) []int {
	queue := []int{pivot}
	paths := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		paths[i] = 0
	}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, u := range graph[v] {
			if paths[u] == 0 {
				paths[u] = paths[v] + 1
				queue = append(queue, u)
			}
		}
	}
	return paths
}
func contains(slice []int, element int) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

func EqualDistances(graph [][]int, referenceVertices []int) []int {
	n := len(graph)
	ResultPaths := make(map[int]int)
	for _, pivot := range referenceVertices {
		paths := bfs(graph, pivot)
		for i := 0; i < n; i++ {
			value, exists := ResultPaths[i]
			if !exists {
				ResultPaths[i] = paths[i]
			} else if value != -1 && value != paths[i] {
				ResultPaths[i] = -1
			}
		}
	}
	var result []int
	for i := 0; i < n; i++ {
		if ResultPaths[i] != -1 && ResultPaths[i] != 0 && !contains(referenceVertices, i) {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	var n, m, k int
	fmt.Scan(&n, &m)
	graph := make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	fmt.Scan(&k)
	referenceVertices := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&referenceVertices[i])
	}

	result := EqualDistances(graph, referenceVertices)
	if len(result) != 0 {
		for _, i := range result {
			fmt.Print(i, " ")
		}
	} else {
		fmt.Print("-")
	}
}
