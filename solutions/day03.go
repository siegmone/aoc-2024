package solutions

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Day03() {
	const input_file = "inputs/day03.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 03 Solutions:\n")
	var start = time.Now()
	sol_1, err := d03_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day03 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d (%s)\n", sol_1, time.Since(start))

	start = time.Now()
	sol_2, err := d03_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day03 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d (%s)\n", sol_2, time.Since(start))
}

func eval_instruction(instruction string) (int, error) {
	start := 0
	current := 0
	var args []int
	for current < len(instruction) {
		start = current
		c := instruction[current]
		current++
		if is_digit(c) {
			peek := instruction[current]
			for is_digit(peek) {
				current++
				peek = instruction[current]
			}
			num, err := strconv.Atoi(instruction[start:current])
			if err != nil {
				fmt.Println("Error parsing the number")
				return -1, err
			}
			args = append(args, num)
		}
	}

	if len(args) != 2 {
		return -1, errors.New("Something went wrong, mul(*,*) doesn't have 2 args")
	}

	result := args[0] * args[1]

	return result, nil
}

func d03_part_1(data string) (int, error) {
	ans := 0
	rgx := regexp.MustCompile(`mul\((\d+,\d+)\)`)
	instructions := rgx.FindAllString(data, -1)
	for _, instr := range instructions {
		mul, err := eval_instruction(instr)
		if err != nil {
			fmt.Println("Wrong evaluation")
			return 0, err
		}
		ans += mul
	}
	return ans, nil
}

func d03_part_2(data string) (int, error) {
	ans := 0

	var sb strings.Builder

	right := strings.Clone(data)
	for {
		split := strings.SplitN(right, "don't()", 2)
		sb.WriteString(split[0])
		if len(split) == 1 {
			break
		}
		right = split[1]
		split = strings.SplitN(right, "do()", 2)
		if len(split) == 1 {
			break
		}
		right = split[1]
	}

	rgx := regexp.MustCompile(`mul\((\d+,\d+)\)`)
	instructions := rgx.FindAllString(sb.String(), -1)

	for _, instr := range instructions {
		mul, err := eval_instruction(instr)
		if err != nil {
			fmt.Println("Wrong evaluation")
			return 0, err
		}
		ans += mul
	}

	return ans, nil
}
