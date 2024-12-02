package solutions

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Day15() {
	const input_file = "inputs/day15.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 15 Solutions:\n")
	sol_1, err := d15_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day15 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d\n", sol_1)
	sol_2, err := d15_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day15 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d\n", sol_2)
}

func d15_part_1(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sort.Slice(lines, func(i, j int) bool {
		return true
	})
	return 0, nil
}

func d15_part_2(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sort.Slice(lines, func(i, j int) bool {
		return true
	})
	return 0, nil
}
