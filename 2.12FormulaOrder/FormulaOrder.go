package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Formula struct {
	Variables   []string
	Expressions []string
	Dependents  []int
	Visited     bool
}

var formulas []Formula
var sortedFormulas []int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Чтение формул из стандартного потока ввода
	for scanner.Scan() {
		line := scanner.Text()

		// Разделение формулы на левую и правую части
		parts := strings.Split(line, "=")
		left := strings.TrimSpace(parts[0])
		right := strings.TrimSpace(parts[1])

		// Проверка наличия переменных в левой части
		variables := strings.Split(left, ",")
		for _, variable := range variables {
			if !isValidVariable(variable) {
				fmt.Println("syntax error")
				return
			}
		}

		// Проверка наличия переменных в правой части
		expressions := strings.Split(right, ",")
		for _, expression := range expressions {
			if !isValidExpression(expression, variables) {
				fmt.Println("syntax error")
				return
			}
		}

		formula := Formula{
			Variables:   variables,
			Expressions: expressions,
			Dependents:  make([]int, 0),
			Visited:     false,
		}

		// Поиск зависимостей формул
		for i, f := range formulas {
			for _, variable := range variables {
				if contains(f.Variables, variable) {
					formula.Dependents = append(formula.Dependents, i)
				}
			}
		}

		formulas = append(formulas, formula)
	}

	// Проверка наличия формулы для каждой переменной
	for _, formula := range formulas {
		for _, variable := range formula.Variables {
			if !hasFormulaForVariable(variable) {
				fmt.Println("syntax error")
				return
			}
		}
	}

	// Выполнение топологической сортировки
	for i, formula := range formulas {
		if !formula.Visited {
			err := depthFirstSearch(i)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	// Вывод отсортированных формул
	for _, index := range sortedFormulas {
		formula := formulas[index]
		fmt.Printf("%s = %s\n", strings.Join(formula.Variables, ", "), strings.Join(formula.Expressions, ", "))
	}
}

func depthFirstSearch(index int) error {
	formula := formulas[index]
	formula.Visited = true

	for _, dependent := range formula.Dependents {
		if formulas[dependent].Visited {
			return fmt.Errorf("cycle")
		}

		err := depthFirstSearch(dependent)
		if err != nil {
			return err
		}
	}

	sortedFormulas = append(sortedFormulas, index)
	return nil
}

func isValidVariable(variable string) bool {
	// Проверка имени переменной
	// Ваша логика проверки может отличаться
	// Я просто проверяю, что имя начинается с буквы и состоит из букв и цифр
	if len(variable) < 1 || !isLetter(variable[0]) {
		return false
	}

	for i := 1; i < len(variable); i++ {
		if !isLetter(variable[i]) && !isDigit(variable[i]) {
			return false
		}
	}

	return true
}

func isValidExpression(expression string, variables []string) bool {
	// Проверка выражения на корректность
	// Ваша логика проверки может отличаться
	// Я просто проверяю, что выражение содержит только переменные и арифметические операции

	for _, variable := range variables {
		expression = strings.ReplaceAll(expression, variable, "1")
	}

	_, err := strconv.ParseFloat(expression, 64)
	if err != nil {
		return false
	}

	return true
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func hasFormulaForVariable(variable string) bool {
	for _, formula := range formulas {
		if contains(formula.Variables, variable) {
			return true
		}
	}
	return false
}
