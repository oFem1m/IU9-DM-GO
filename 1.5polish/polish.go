package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expr := scanner.Text()
	result := evalPrefixExpression(expr)
	fmt.Println(result)
}

func evalPrefixExpression(expr string) int {
	var stack []int
	operators := "+-*/"
	expr = removeBracketsAndSpaces(expr)
	for i := len(expr) - 1; i >= 0; i-- {
		char := string(expr[i])
		if strings.Contains(operators, char) {
			if len(stack) < 2 {
				panic("Not enough numbers in stack")
			}
			var op1, op2 int
			op1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			op2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			switch char {
			case "+":
				stack = append(stack, op1+op2)
			case "-":
				stack = append(stack, op1-op2)
			case "*":
				stack = append(stack, op1*op2)
			case "/":
				stack = append(stack, op1/op2)
			}
		} else {
			digit, _ := strconv.Atoi(char)
			stack = append(stack, digit)
		}
	}
	if len(stack) != 1 {
		panic("Invalid expression format")
	}
	return stack[0]
}

func removeBracketsAndSpaces(str string) string {
	var newStr []rune // создаем пустой слайс рун
	for _, char := range str {
		if char != '(' && char != ')' && char != ' ' { // если символ не скобка, то добавляем его в новую строку
			newStr = append(newStr, char)
		}
	}
	return string(newStr) // конвертируем слайс рун в строку и возвращаем ее
}
