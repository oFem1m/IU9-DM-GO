package main

import "fmt"

var array = [5]int{3, 5, 2, 4, 1}

func main() {
	qsort(5, less, swap)
	fmt.Println(array)
}

func less(i, j int) bool {
	return array[i] < array[j]
}

func swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	if n > 1 {
		stack := []int{0, n - 1}
		for len(stack) > 0 {
			h := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			l := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i := l - 1
			for j := l; j < h; j++ {
				if less(j, h) {
					i++
					swap(i, j)
				}
			}
			swap(i+1, h)
			pivot := i + 1

			if pivot-1 > l {
				stack = append(stack, l, pivot-1)
			}

			if pivot+1 < h {
				stack = append(stack, pivot+1, h)
			}
		}
	}
}
