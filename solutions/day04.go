package solutions

import (
	"fmt"
	"os"
	"strings"
)

func Day04() {
	const input_file = "inputs/day04.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 04 Solutions:\n")
	sol_1, err := d04_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day04 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d\n", sol_1)
	sol_2, err := d04_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day04 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d\n", sol_2)
}

func matrix_el(m []string, cols int, i int, j int) string {
	return m[(i*cols)+j]
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}

func search(grid []string, rows int, cols int, i int, j int, word string) int {
	result := 0

	if matrix_el(grid, cols, i, j) != string(word[0]) {
		return result
	}

	x := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	y := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}

	for dir := range 8 {
		curr_r := i + x[dir]
		curr_c := j + y[dir]
		k := 1

		for k < len(word) {
			if curr_r >= rows || curr_r < 0 || curr_c >= cols || curr_c < 0 {
				break
			}

			if matrix_el(grid, cols, curr_r, curr_c) != string(word[k]) {
				break
			}

			curr_r += x[dir]
			curr_c += y[dir]
			k++

		}

		if k == len(word) {
			result++
		}
	}

	return result
}

func search_mas(grid []string, rows int, cols int, i int, j int) bool {
	if matrix_el(grid, cols, i, j) != "A" {
		return false
	}

	x := [8]int{-1, -1, 1, 1}
	y := [8]int{-1, 1, -1, 1}

	for dir := range 4 {
		curr_r := i + x[dir]
		curr_c := j + y[dir]

		if curr_r >= rows || curr_r < 0 || curr_c >= cols || curr_c < 0 {
			return false
		}

		el := matrix_el(grid, cols, curr_r, curr_c)

		if el != "M" && el != "S" {
			return false
		}

		curr_r += x[dir]
		curr_c += y[dir]
	}

    nw := matrix_el(grid, cols, i - 1, j - 1)
    ne := matrix_el(grid, cols, i - 1, j + 1)
    sw := matrix_el(grid, cols, i + 1, j - 1)
    se := matrix_el(grid, cols, i + 1, j + 1)

    if nw == se || ne == sw {
        return false
    }

	return true
}

func d04_part_1(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var grid []string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, "")...)
	}

	ans := 0

	rows := len(lines)
	cols := len(lines[0])

	for i := range rows {
		for j := range cols {
			found := search(grid, rows, cols, i, j, "XMAS")
			if found > 0 {
				ans += found
			}
		}
	}

	return ans, nil
}

func d04_part_2(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var grid []string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, "")...)
	}

	ans := 0

	rows := len(lines)
	cols := len(lines[0])

	for i := range rows {
		for j := range cols {
			found := search_mas(grid, rows, cols, i, j)
			if found {
				ans++
			}
		}
	}

	return ans, nil
}
