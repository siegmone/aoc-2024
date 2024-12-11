package solutions

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
    "time"
)

func Day10() {
	const input_file = "inputs/day10.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 10 Solutions:\n")
	var start = time.Now()
	sol_1, err := d10_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day10 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d (%s)\n", sol_1, time.Since(start))

	start = time.Now()
	sol_2, err := d10_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day10 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d (%s)\n", sol_2, time.Since(start))
}

func position_in_grid(grid [][]int, pos Vector2) bool {
	height := len(grid)
	width := len(grid[0])
	if pos.X >= 0 && pos.X < width && pos.Y >= 0 && pos.Y < height {
		return true
	} else {
		return false
	}
}

func traverse_trailhead(grid [][]int, start Vector2) int {
	score := 0
	var visited []Vector2
	x_dir := [4]int{-1, 0, 0, 1}
	y_dir := [4]int{0, -1, 1, 0}

	var helper func(grid [][]int, s Vector2)
	helper = func(grid [][]int, s Vector2) {
		visited = append(visited, s)
		current_height := grid[s.Y][s.X]
		if current_height == 9 {
			score++
			return
		}
		for i := range x_dir {
			next := Vector2{X: s.X + x_dir[i], Y: s.Y + y_dir[i]}
			if !position_in_grid(grid, next) {
				continue
			}
			if slices.Contains(visited, next) {
				continue
			}
			next_height := grid[next.Y][next.X]
			if next_height != current_height+1 {
				continue
			}
			helper(grid, next)
		}
	}

	helper(grid, start)

	return score
}

func get_possible_paths(grid [][]int, start Vector2) int {
	rating := 0
	var q []Vector2

	x_dir := [4]int{-1, 0, 0, 1}
	y_dir := [4]int{0, -1, 1, 0}

	q = append(q, start)

	for len(q) > 0 {
		var current Vector2
		current, q = q[0], q[1:]
		current_height := grid[current.Y][current.X]
		if current_height == 9 {
			rating++
		}
		for i := range x_dir {
			next := Vector2{X: current.X + x_dir[i], Y: current.Y + y_dir[i]}
			if !position_in_grid(grid, next) {
				continue
			}
			next_height := grid[next.Y][next.X]
			if next_height != current_height+1 {
				continue
			}
			q = append(q, next)
		}
	}
	return rating
}

func d10_part_1(data string) (int, error) {
	grid := mapfunc(
		strings.Split(strings.TrimSpace(string(data)), "\n"),
		func(s string) []int {
			return mapfunc(strings.Split(s, ""), func(ss string) int {
				r, _ := strconv.Atoi(ss)
				return r
			})
		})
	ans := 0

	for y, row := range grid {
		for x, el := range row {
			if el == 0 {
				score := traverse_trailhead(grid, Vector2{x, y})
				ans += score
			}
		}
	}
	return ans, nil
}

func d10_part_2(data string) (int, error) {
	grid := mapfunc(
		strings.Split(strings.TrimSpace(string(data)), "\n"),
		func(s string) []int {
			return mapfunc(strings.Split(s, ""), func(ss string) int {
				r, _ := strconv.Atoi(ss)
				return r
			})
		})
	ans := 0

	for y, row := range grid {
		for x, el := range row {
			if el == 0 {
				rating := get_possible_paths(grid, Vector2{x, y})
				ans += rating
			}
		}
	}
	return ans, nil
}
