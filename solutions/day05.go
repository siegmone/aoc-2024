package solutions

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Day05() {
	const input_file = "inputs/day05.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 05 Solutions:\n")
	var start = time.Now()
	sol_1, err := d05_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day05 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d (%s)\n", sol_1, time.Since(start))

	start = time.Now()
	sol_2, err := d05_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day05 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d (%s)\n", sol_2, time.Since(start))
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

		if !rule_check {
			sort.Slice(pages, func(i, j int) bool {
				if check_rules(pages[i], pages[j], rules) < 0 {
					return true
				}
				return false
			})
			middle, err := strconv.Atoi(pages[len(pages)/2])
			if err == nil {
				ans += middle
			}
		}
	}

	return ans, nil
}
