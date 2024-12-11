package solutions

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func Day13() {
	const input_file = "inputs/day13.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 13 Solutions:\n")
	var start = time.Now()
	sol_1, err := d13_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day13 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d (%s)\n", sol_1, time.Since(start))

	start = time.Now()
	sol_2, err := d13_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day13 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d (%s)\n", sol_2, time.Since(start))
}

func d13_part_1(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sort.Slice(lines, func(i, j int) bool {
		return true
	})
	return 0, nil
}

func d13_part_2(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sort.Slice(lines, func(i, j int) bool {
		return true
	})
	return 0, nil
}
