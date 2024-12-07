package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operator int

const (
	PLUS Operator = iota
	MULT
	CONCAT
)

var operators_p1 []Operator = []Operator{PLUS, MULT}
var operators_p2 []Operator = []Operator{PLUS, MULT, CONCAT}

func Day07() {
	const input_file = "inputs/day07.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 07 Solutions:\n")
	sol_1, err := d07_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day07 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d\n", sol_1)
	sol_2, err := d07_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day07 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d\n", sol_2)
}

func parse_result_values(l string) (int, []int) {
	split := strings.SplitN(l, ": ", 2)
	result, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Println("Error converting result")
		os.Exit(1)
	}
	values := mapfunc(strings.Split(split[1], " "), func(s string) int {
		v, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error converting values")
			return 0
		}
		return v
	})
	return result, values
}

func apply_operator(a int, b int, op Operator) int {
	var result int
	switch op {
	case PLUS:
		{
			result = a + b
		}
	case MULT:
		{
			result = a * b
		}
	case CONCAT:
		{
			a_str := strconv.Itoa(a)
			b_str := strconv.Itoa(b)
			c, err := strconv.Atoi(a_str + b_str)
			if err != nil {
				return 0
			}
			result = c
		}
	default:
		panic("WTF")
	}
	return result
}

func eval_expression(values []int, operators []Operator) (int, error) {
	if len(values) != len(operators)+1 {
		return 0, fmt.Errorf("Numbers of values and operators are not compatible")
	}
	if len(values) == 2 {
		return apply_operator(values[0], values[1], operators[0]), nil
	}
	result, err := eval_expression(values[:len(values)-1], operators[:len(operators)-1])
	return apply_operator(result, values[len(values)-1], operators[len(operators)-1]), err
}

func generate_permutations[T any](elements []T, length int) [][]T {
	var result [][]T
	var permutation []T

	var helper func()
	helper = func() {
		if len(permutation) == length {
			temp := make([]T, length)
			copy(temp, permutation)
			result = append(result, temp)
			return
		}
		for _, element := range elements {
			permutation = append(permutation, element)
			helper()
			permutation = permutation[:len(permutation)-1]
		}
	}

	helper()
	return result
}

func d07_part_1(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	ans := 0

	for _, line := range lines {
		ok := false
		test_value, values := parse_result_values(line)
		operator_permutations := generate_permutations(operators_p1, len(values)-1)
		for _, op_perm := range operator_permutations {
			result, err := eval_expression(values, op_perm)
			if err != nil {
				return 0, err
			}
			ok = (result == test_value)
			if ok {
				break
			}
		}
		if ok {
			ans += test_value
		}
	}

	return ans, nil
}

func d07_part_2(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	ans := 0

	for _, line := range lines {
		ok := false
		test_value, values := parse_result_values(line)
		operator_permutations := generate_permutations(operators_p2, len(values)-1)
		for _, op_perm := range operator_permutations {
			result, err := eval_expression(values, op_perm)
			if err != nil {
				return 0, err
			}
			ok = (result == test_value)
			if ok {
				break
			}
		}
		if ok {
			ans += test_value
		}
	}

	return ans, nil
}
