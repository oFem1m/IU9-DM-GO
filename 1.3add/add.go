package main

import "fmt"

func main() {
	a := []int32{1, 2, 3}
	b := []int32{2, 3, 4}
	p := 10
	result := add(a, b, p)
	fmt.Println(result)
}

func add(a, b []int32, p int) []int32 {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	result := make([]int32, maxLen+1)
	carry := int32(0)
	for i := 0; i < maxLen || carry > 0; i++ {
		sum := carry
		if i < len(a) {
			sum += a[i]
		}
		if i < len(b) {
			sum += b[i]
		}
		result[i] = sum % int32(p)
		carry = sum / int32(p)
	}
	// удаление нулей в начале
	for len(result) > 1 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}
	return result
}
