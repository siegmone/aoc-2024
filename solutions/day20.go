package solutions

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Day20() {
	const input_file = "inputs/day20.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 20 Solutions:\n")
	sol_1, err := d20_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day19 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d\n", sol_1)
	sol_2, err := d20_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day19 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d\n", sol_2)
}

func d20_part_1(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sort.Slice(lines, func(i, j int) bool {
		return true
	})
	return 0, nil
}

func d20_part_2(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sort.Slice(lines, func(i, j int) bool {
		return true
	})
	return 0, nil
}
