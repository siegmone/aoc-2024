package solutions

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day02() {
	const input_file = "inputs/day02.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 02 Solutions:\n")
	sol_1, err := d02_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day02 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d\n", sol_1)
	sol_2, err := d02_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day02 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d\n", sol_2)
}

func d02_part_1(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	safe_count := 0
	for _, line := range lines {
		levels_str := strings.Split(line, " ")
		var levels []int
		for _, s := range levels_str {
			if n, err := strconv.Atoi(s); err == nil {
				levels = append(levels, n)
			}
		}
		safe := check_safety(levels)
        if safe {
            safe_count++
        }
	}
	return safe_count, nil
}

func d02_part_2(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	safe_count := 0
	for _, line := range lines {
		levels_str := strings.Split(line, " ")
		var levels []int
		for _, s := range levels_str {
			if n, err := strconv.Atoi(s); err == nil {
				levels = append(levels, n)
			}
		}
		safe := check_safety(levels)
        if safe {
            safe_count++
        }
		if !safe {
			for i := 0; i < len(levels); i++ {
				var new_levels []int
				new_levels = append(new_levels, levels[:i]...)
				new_levels = append(new_levels, levels[i+1:]...)
                if check_safety(new_levels) {
                    safe_count++
                    break
                }
			}
		}
	}
	return safe_count, nil
}

func check_safety(levels []int) bool {
	increasing := false
	decreasing := false
	safe := true
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]

		if math.Abs(float64(diff)) <= 0 || math.Abs(float64(diff)) >= 4 {
			safe = false
			break
		}

		if diff < 0 {
			decreasing = true
		} else if diff > 0 {
			increasing = true
		}

		if increasing && decreasing {
			safe = false
			break
		}

	}
	return safe
}
