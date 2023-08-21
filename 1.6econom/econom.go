package main

import (
	"bufio"
	"fmt"
	"os"
)

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func eco(expr string) int {
	var ops []string
	var bracesId []int
	for i := len(expr) - 1; i >= 0; i-- {
		if rune(expr[i]) == ')' {
			bracesId = append([]int{i}, bracesId...)
		} else if rune(expr[i]) == '(' {
			op := expr[i+1 : bracesId[0]]
			if !stringInSlice(op, ops) {
				ops = append(ops, op)
			}
			bracesId = bracesId[1:]
		}
	}
	return len(ops)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expr := scanner.Text()
	fmt.Println(eco(expr))
}
