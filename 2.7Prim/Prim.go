package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	u, v, len int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numHouses, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	numRoads, _ := strconv.Atoi(scanner.Text())

	adjacencyList := make([][]Edge, numHouses)

	for i := 0; i < numRoads; i++ {
		scanner.Scan()
		line := scanner.Text()
		u, v, l := parseRoadInfo(line)

		adjacencyList[u] = append(adjacencyList[u], Edge{u, v, l})
		adjacencyList[v] = append(adjacencyList[v], Edge{v, u, l})
	}

	minTotalLength := prim(adjacencyList)

	fmt.Println(minTotalLength)
}

func parseRoadInfo(line string) (int, int, int) {
	fields := strings.Fields(line)
	u, _ := strconv.Atoi(fields[0])
	v, _ := strconv.Atoi(fields[1])
	l, _ := strconv.Atoi(fields[2])
	return u, v, l
}

func prim(adjacencyList [][]Edge) int {
	numHouses := len(adjacencyList)

	visited := make([]bool, numHouses)
	distances := make([]int, numHouses)
	for i := 0; i < numHouses; i++ {
		distances[i] = 1<<31 - 1
	}
	distances[0] = 0
	totalLength := 0
	for i := 0; i < numHouses; i++ {
		u := -1
		for j := 0; j < numHouses; j++ {
			if !visited[j] && (u == -1 || distances[j] < distances[u]) {
				u = j
			}
		}
		visited[u] = true

		for _, edge := range adjacencyList[u] {
			v := edge.v
			l := edge.len
			if !visited[v] && l < distances[v] {
				distances[v] = l
			}
		}

		totalLength += distances[u]
	}

	return totalLength
}
