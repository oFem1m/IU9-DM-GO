package main

import (
	"fmt"
	"math"
	"sort"
)

func findDivisors(n int) []int {
	divisors := []int{}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(divisors)))

	return divisors
}
func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func getDivisorDivisors(A []int) [][]int {
	result := make([][]int, len(A)-1)
	for i := 0; i < len(A)-1; i++ {
		result[i] = findDivisors(A[i])[1:]
	}

	return result
}
func main() {
	var n int
	fmt.Scanf("%d", &n)
	divisors := findDivisors(n)
	groupDivisors := getDivisorDivisors(divisors)
	for i := 0; i < len(groupDivisors); i++ {
		for j := 0; j < len(groupDivisors[i]); j++ {
			if i+1 < len(groupDivisors) {
				for k := i + 1; k < len(groupDivisors); k++ {
					if divisors[i]%divisors[k] == 0 {
						for l := 0; l < len(groupDivisors[k]); l++ {
							if groupDivisors[i][j] == groupDivisors[k][l] {
								groupDivisors[i] = removeElement(groupDivisors[i], j)
								j--
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("graph {")
	for i := 0; i < len(divisors); i++ {
		fmt.Println(divisors[i])
	}
	for i := 0; i < len(groupDivisors); i++ {
		for j := 0; j < len(groupDivisors[i]); j++ {
			fmt.Printf("%d -- %d\n", divisors[i], groupDivisors[i][j])
		}
	}
	fmt.Println("}")
}
