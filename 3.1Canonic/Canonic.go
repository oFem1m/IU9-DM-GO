package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	used  = make(map[int]bool)
	pos   = 0
	order = make(map[int]int)
)

func DFS(matrix [][]int, v int) {
	used[v] = true
	order[pos] = v
	pos++
	for _, u := range matrix[v] {
		if !used[u] {
			DFS(matrix, u)
		}
	}
}

func reverse(list map[int]int) map[int]int {
	reversedList := make(map[int]int)
	for i := 0; i < len(list); i++ {
		reversedList[list[i]] = i
	}
	return reversedList
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n, m, q0 int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	q0, _ = strconv.Atoi(scanner.Text())

	transition := make([][]int, n)
	output := make([][]string, n)
	for i := 0; i < n; i++ {
		transition[i] = make([]int, m)
		scanner.Scan()
		line := scanner.Text()
		numbers := strings.Fields(line)
		for j := 0; j < m; j++ {
			transition[i][j], _ = strconv.Atoi(numbers[j])
		}
	}

	for i := 0; i < n; i++ {
		output[i] = make([]string, m)
		scanner.Scan()
		line := scanner.Text()
		lineList := strings.Fields(line)
		for j := 0; j < m && j < len(lineList); j++ {
			if len(lineList[j]) != 0 {
				output[i][j] = lineList[j]
			}
		}
	}
	DFS(transition, q0)
	orderReversed := reverse(order)

	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(0)
	for i := 0; i < len(order); i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", orderReversed[transition[order[i]][j]])
		}
		fmt.Println()
	}

	for i := 0; i < len(order); i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%s ", output[order[i]][j])
		}
		fmt.Println()
	}
}
