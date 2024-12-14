package solutions

import (
	"fmt"
	"os"
	"strconv"
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

func parse_button(button string) (int64, int64) {
	split := strings.SplitN(strings.TrimSpace(button), ": ", 2)
	split = strings.SplitN(strings.TrimSpace(split[1]), ", ", 2)
	dx_split := strings.SplitN(strings.TrimSpace(split[0]), "+", 2)
	dy_split := strings.SplitN(strings.TrimSpace(split[1]), "+", 2)
	if len(dx_split) != 2 || len(dy_split) != 2 {
		panic("Malformed button")
	}
	dx, err := strconv.ParseInt(dx_split[1], 10, 0)
	if err != nil {
		panic("Couldn't convert dx to integer")
	}
	dy, err := strconv.ParseInt(dy_split[1], 10, 0)
	if err != nil {
		panic("Couldn't convert dy to integer")
	}
	return dx, dy
}

func parse_prize(prize string) (int64, int64) {
	split := strings.SplitN(strings.TrimSpace(prize), ": ", 2)
	split = strings.SplitN(strings.TrimSpace(split[1]), ", ", 2)
	x_split := strings.SplitN(strings.TrimSpace(split[0]), "=", 2)
	y_split := strings.SplitN(strings.TrimSpace(split[1]), "=", 2)
	if len(x_split) != 2 || len(y_split) != 2 {
		panic("Malformed prize")
	}
	x, err := strconv.ParseInt(x_split[1], 10, 0)
	if err != nil {
		panic("Couldn't convert dx to integer")
	}
	y, err := strconv.ParseInt(y_split[1], 10, 0)
	if err != nil {
		panic("Couldn't convert dy to integer")
	}
	return x, y
}

func parse_machine(machine string, part_2 bool) int64 {
	var a, b, a_cost, b_cost int64

	lines := strings.Split(strings.TrimSpace(machine), "\n")

	a_cost = 3
	b_cost = 1

	a_x, a_y := parse_button(lines[0])
	b_x, b_y := parse_button(lines[1])
	p_x, p_y := parse_prize(lines[2])

	if part_2 {
		var offset int64
		offset = 10000000000000
		p_x += offset
		p_y += offset
	}

	den := a_x*b_y - a_y*b_x

	a = (p_x*b_y - p_y*b_x) / den
	b = (p_y*a_x - p_x*a_y) / den
	p_x_expected := a*a_x + b*b_x
	p_y_expected := a*a_y + b*b_y

	if (p_x_expected != p_x) || (p_y_expected != p_y) {
		return 0
	}

	return a*a_cost + b*b_cost
}

func d13_part_1(data string) (int64, error) {
	machines := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	var ans int64
	ans = 0
	for _, machine := range machines {
		ans += parse_machine(machine, false)
	}

	return ans, nil
}

func d13_part_2(data string) (int64, error) {
	machines := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	var ans int64
	ans = 0
	for _, machine := range machines {
		ans += parse_machine(machine, true)
	}

	return ans, nil
}
