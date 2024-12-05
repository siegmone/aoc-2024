package solutions

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day05() {
	const input_file = "inputs/day05.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 05 Solutions:\n")
	sol_1, err := d05_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day05 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d\n", sol_1)
	sol_2, err := d05_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day05 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d\n", sol_2)
}

func mapfunc[T, U any](data []T, f func(T) U) []U {

	res := make([]U, 0, len(data))

	for _, e := range data {
		res = append(res, f(e))
	}

	return res
}

func check_rules(a string, b string, rules []string) int {
	rule := fmt.Sprintf("%s|%s", a, b)
	if slices.Contains(rules, rule) {
		return -1
	}
	return 1
}

func d05_part_1(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	rules := strings.Split(lines[0], "\n")
	updates := strings.Split(lines[1], "\n")

	ans := 0
	for _, u := range updates {
		rule_check := true
		pages := strings.Split(u, ",")
		for i := len(pages) - 1; i > 0; i-- {
			if check_rules(pages[i], pages[i-1], rules) < 0 {
				rule_check = false
			}
		}

		if rule_check {
			middle, err := strconv.Atoi(pages[len(pages)/2])
			if err == nil {
				ans += middle
			}
		}
	}

	return ans, nil
}

func d05_part_2(data string) (int, error) {
	// lines := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	ans := 0

	return ans, nil
}
