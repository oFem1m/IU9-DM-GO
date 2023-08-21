package main

import (
	"fmt"
)

func main() {
	var n, m, q0 int
	fmt.Scan(&n, &m, &q0)

	transition := make([][]int, n)
	output := make([][]string, n)

	for i := 0; i < n; i++ {
		transition[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&transition[i][j])
		}
	}

	for i := 0; i < n; i++ {
		output[i] = make([]string, m)
		for j := 0; j < m; {
			var tmp string
			fmt.Scan(&tmp)
			if len(tmp) != 0 {
				output[i][j] = tmp
				j++
			}
		}
	}

	fmt.Println("digraph {")
	fmt.Println("    rankdir = LR")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("    %d -> %d [label = \"%c(%s)\"]\n", i, transition[i][j], 'a'+j, output[i][j])
		}
	}
	fmt.Println("}")
}
