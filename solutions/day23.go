package solutions

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func Day23() {
	const input_file = "inputs/day23.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 23 Solutions:\n")
	var start = time.Now()
	sol_1, err := d23_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day23 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d (%s)\n", sol_1, time.Since(start))

	start = time.Now()
	sol_2, err := d23_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day23 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d (%s)\n", sol_2, time.Since(start))
}

func d23_part_1(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sort.Slice(lines, func(i, j int) bool {
		return true
	})
	return 0, nil
}

func d23_part_2(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sort.Slice(lines, func(i, j int) bool {
		return true
	})
	return 0, nil
}
