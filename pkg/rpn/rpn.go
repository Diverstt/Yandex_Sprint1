package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Stack представляет стек чисел
type Stack []float64

// Add добавляет элемент на вершину стека
func (s *Stack) Add(value float64) {
	*s = append(*s, value)
}

// Del удаляет элемент с вершины стека
func (s *Stack) Del() (float64, error) {
	if len(*s) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = append((*s)[:index])

	return value, nil
}

// OperatorsStack представляет стек операторов
type OperatorsStack []string

// Add добавляет оператор на вершину стека
func (s *OperatorsStack) Add(value string) {
	*s = append(*s, value)
}

// Del удаляет оператор с вершины стека
func (s *OperatorsStack) Del() (string, error) {
	if len(*s) == 0 {
		return "", fmt.Errorf("stack is empty")
	}

	index := len(*s) - 1
	value := (*s)[index]
	*s = append((*s)[:index])

	return value, nil
}

// DoOperation выполняет арифметическую операцию над двумя числами
func doOperation(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", operator)
	}
}

// FindMistake проверяет корректность ввода
func findMistake(str string) error {
	re := regexp.MustCompile(`^[*/+\-]|[*/+\-]$`)
	matches := re.FindAllString(str, -1)
	if len(matches) != 0 {
		return fmt.Errorf("false input format")
	}

	re = regexp.MustCompile(`[*/+\-]{2,}`)
	matches = re.FindAllString(str, -1)
	if len(matches) != 0 {
		return fmt.Errorf("false input format")
	}

	re = regexp.MustCompile(`[()]`)
	matches = re.FindAllString(str, -1)
	if len(matches)%2 != 0 {
		return fmt.Errorf("false input format")
	}

	return nil
}

// IsDigit проверяет, является ли строка числом
func isDigit(item string) (float64, error) {
	num, err := strconv.ParseFloat(item, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", item)
	}

	return num, nil
}

// GetTokenAndOperand возвращает список токенов и операндов
func getTokenAndOperand(str string) ([]string, error) {
	if err := findMistake(str); err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`[0-9]+|[*/+\-()]`)

	return re.FindAllString(str, -1), nil
}

// CheckPriority возвращает приоритет оператора
func checkPriority(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

// processNumber обрабатывает число и добавляет его в стек
func processNumber(token string, stack *Stack) error {
	num, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return fmt.Errorf("invalid number: %s", token)
	}
	stack.Add(num)

	return nil
}

// processOperator обрабатывает оператор
func processOperator(token string, stack *Stack, operators *OperatorsStack) error {
	for len(*operators) > 0 {
		op, _ := operators.Del()
		if op == "(" || checkPriority(token) > checkPriority(op) {
			operators.Add(op)
			break
		}
		b, _ := stack.Del()
		a, _ := stack.Del()
		result, err := doOperation(a, b, op)
		if err != nil {
			return err
		}
		stack.Add(result)
	}
	operators.Add(token)

	return nil
}

// processClosingBracket обрабатывает закрывающую скобку
func processClosingBracket(stack *Stack, operators *OperatorsStack) error {
	for len(*operators) > 0 {
		op, _ := operators.Del()
		if op == "(" {
			break
		}
		b, _ := stack.Del()
		a, _ := stack.Del()
		result, err := doOperation(a, b, op)
		if err != nil {
			return err
		}
		stack.Add(result)
	}

	return nil
}

// processToken обрабатывает токен
func processToken(token string, stack *Stack, operators *OperatorsStack) error {
	if num, err := strconv.ParseFloat(token, 64); err == nil {
		stack.Add(num)
	} else {
		switch token {
		case "(":
			operators.Add(token)
		case ")":
			return processClosingBracket(stack, operators)
		default:
			return processOperator(token, stack, operators)
		}
	}

	return nil
}

// Calc вычисляет результат выражения
func Calc(expression string) (float64, error) {
	tokens, err := getTokenAndOperand(expression)
	if err != nil {
		return 0, err
	}

	stack := Stack{}
	operators := OperatorsStack{}

	for _, token := range tokens {
		if err := processToken(token, &stack, &operators); err != nil {
			return 0, err
		}
	}

	for len(operators) > 0 {
		op, _ := operators.Del()
		b, _ := stack.Del()
		a, _ := stack.Del()
		result, err := doOperation(a, b, op)
		if err != nil {
			return 0, err
		}
		stack.Add(result)
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	result, _ := stack.Del()

	return result, nil
}
