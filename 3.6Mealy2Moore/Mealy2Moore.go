package main

import (
	"fmt"
)

func readStrings(size int) []string {
	strings := make([]string, size)
	for i := range strings {
		fmt.Scan(&strings[i])
	}
	return strings
}

func readIntMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			fmt.Scan(&matrix[i][j])
		}
	}
	return matrix
}

func generateVertices(outputSize, n int, alphabetOutput []string) map[string]int {
	vertices := make(map[string]int)
	counter := 0

	for i := 0; i < n; i++ {
		for j := 0; j < outputSize; j++ {
			tmpVertex := fmt.Sprintf("(%d,%s)", i, alphabetOutput[j])
			vertices[tmpVertex] = counter
			counter++
		}
	}

	return vertices
}

func generateReachableMap(inputSize, outputSize, n int, transitionMatrix, outputMatrix [][]int, alphabetOutput []string) map[string]bool {
	isReachable := make(map[string]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < outputSize; j++ {
			for k := 0; k < inputSize; k++ {
				edgeVertex := fmt.Sprintf("(%d,%s)", transitionMatrix[i][k], alphabetOutput[outputMatrix[i][k]])
				isReachable[edgeVertex] = true
			}
		}
	}

	return isReachable
}

func printGraph(outputSize int, vertices map[string]int, isReachable map[string]bool, alphabetOutput, alphabetInput []string, transitionMatrix, outputMatrix [][]int) {
	fmt.Println("digraph {")
	fmt.Println("rankdir = LR")

	for i := range vertices {
		if isReachable[i] {
			fmt.Printf("%d [label = \"%s\"]\n", vertices[i], i)
		}
	}

	for i := range vertices {
		if isReachable[i] {
			for k := 0; k < len(alphabetInput); k++ {
				edgeVertex := fmt.Sprintf("(%d,%s)", transitionMatrix[vertices[i]/outputSize][k], alphabetOutput[outputMatrix[vertices[i]/outputSize][k]])
				fmt.Printf("%d -> %d [label = \"%s\"]\n", vertices[i], vertices[edgeVertex], alphabetInput[k])
			}
		}
	}

	fmt.Println("}")
}

func main() {
	var inputSize, outputSize, n int
	fmt.Scan(&inputSize)

	alphabetInput := readStrings(inputSize)

	fmt.Scan(&outputSize)

	alphabetOutput := readStrings(outputSize)

	fmt.Scan(&n)

	transitionMatrix := readIntMatrix(n, inputSize)

	outputMatrix := readIntMatrix(n, inputSize)

	vertices := generateVertices(outputSize, n, alphabetOutput)

	isReachable := generateReachableMap(inputSize, outputSize, n, transitionMatrix, outputMatrix, alphabetOutput)

	printGraph(outputSize, vertices, isReachable, alphabetOutput, alphabetInput, transitionMatrix, outputMatrix)
}
