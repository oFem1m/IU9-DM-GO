package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x, y int
}

type Edge struct {
	u, v   int
	weight float64
}

type ByWeight []Edge

func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Less(i, j int) bool { return a[i].weight < a[j].weight }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
	}
	return parent[i]
}

func union(parent, rank []int, x, y int) {
	xroot := find(parent, x)
	yroot := find(parent, y)
	if rank[xroot] < rank[yroot] {
		parent[xroot] = yroot
	} else if rank[xroot] > rank[yroot] {
		parent[yroot] = xroot
	} else {
		parent[yroot] = xroot
		rank[xroot]++
	}
}

func calculateDistance(p1, p2 Point) float64 {
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	var n int
	fmt.Scanln(&n)

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&points[i].x, &points[i].y)
	}

	edges := make([]Edge, 0)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			weight := calculateDistance(points[i], points[j])
			edges = append(edges, Edge{i, j, weight})
		}
	}

	sort.Sort(ByWeight(edges))

	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	minCost := 0.0
	for _, edge := range edges {
		if find(parent, edge.u) != find(parent, edge.v) {
			minCost += edge.weight
			union(parent, rank, edge.u, edge.v)
		}
	}

	fmt.Printf("%.2f\n", minCost)
}
