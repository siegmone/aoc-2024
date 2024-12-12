package solutions

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func print_direction(d Direction) {
	switch d {
	case RIGHT:
		fmt.Println("RIGHT")
	case LEFT:
		fmt.Println("LEFT")
	case UP:
		fmt.Println("UP")
	case DOWN:
		fmt.Println("DOWN")
	}
}

func turn(current Direction, where Direction) Direction {
	var next Direction
	switch where {
	case RIGHT:
		next = (current + 1) % 4
	case LEFT:
		if current == 0 {
			next = 3
		} else {
			next = (current - 1) % 4
		}
	}
	return next
}

func on_the_right(pos Vector2, dir Direction) Vector2 {
	var result Vector2
	switch dir {
	case RIGHT:
		result = pos.add_vector(&Vector2{0, 1})
	case LEFT:
		result = pos.add_vector(&Vector2{0, -1})
	case UP:
		result = pos.add_vector(&Vector2{1, 0})
	case DOWN:
		result = pos.add_vector(&Vector2{-1, 0})
	}
	return result
}

func move_vec(vec Vector2, dir Direction) Vector2 {
	var next Vector2
	switch dir {
	case RIGHT:
		next = vec.add_vector(&Vector2{1, 0})
	case LEFT:
		next = vec.add_vector(&Vector2{-1, 0})
	case UP:
		next = vec.add_vector(&Vector2{0, -1})
	case DOWN:
		next = vec.add_vector(&Vector2{0, 1})
	}
	return next
}

func Day12() {
	const input_file = "inputs/day12.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 12 Solutions:\n")
	var start = time.Now()
	sol_1, err := d12_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day12 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d (%s)\n", sol_1, time.Since(start))

	start = time.Now()
	sol_2, err := d12_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day12 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d (%s)\n", sol_2, time.Since(start))
}

func fill_plant_region(grid [][]string, pos Vector2, plant string) []Vector2 {
	var region []Vector2
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	q := []Vector2{pos}
	x_dir := [4]int{-1, 0, 0, 1}
	y_dir := [4]int{0, -1, 1, 0}

	for len(q) > 0 {
		var current_pos Vector2
		current_pos, q = q[0], q[1:]
		if visited[current_pos.Y][current_pos.X] {
			continue
		}
		visited[current_pos.Y][current_pos.X] = true
		region = append(region, current_pos)
		for i := range x_dir {
			next := Vector2{
				X: current_pos.X + x_dir[i],
				Y: current_pos.Y + y_dir[i]}
			if position_in_grid(grid, next) &&
				grid[next.Y][next.X] == plant &&
				!visited[next.Y][next.X] {
				q = append(q, next)
			}
		}
	}
	return region
}

func region_perimeter(region []Vector2) int {
	x_dir := [4]int{-1, 0, 0, 1}
	y_dir := [4]int{0, -1, 1, 0}

	perimeter := 0

	for _, cell := range region {
		for i := range x_dir {
			neighbor := Vector2{
				X: cell.X + x_dir[i],
				Y: cell.Y + y_dir[i],
			}
			if !slices.Contains(region, neighbor) {
				perimeter++
			}
		}
	}
	return perimeter
}

func region_area(region []Vector2) int {
	return len(region)
}

func region_sides(region []Vector2) int {
	sides := 0

	region_map := make(map[Vector2]bool)
	for _, cell := range region {
		region_map[cell] = true
	}

	for _, cell := range region {
		// get the directions
		left := cell.add_vector(&Vector2{-1, 0})
		right := cell.add_vector(&Vector2{1, 0})
		top := cell.add_vector(&Vector2{0, -1})
		bottom := cell.add_vector(&Vector2{0, 1})

		top_left := cell.add_vector(&Vector2{-1, -1})
		top_right := cell.add_vector(&Vector2{1, -1})
		bottom_left := cell.add_vector(&Vector2{-1, 1})
		bottom_right := cell.add_vector(&Vector2{1, 1})

		// check if it's part of the region
		_, ok_left := region_map[left]
		_, ok_right := region_map[right]
		_, ok_top := region_map[top]
		_, ok_bottom := region_map[bottom]

		_, ok_tl := region_map[top_left]
		_, ok_tr := region_map[top_right]
		_, ok_bl := region_map[bottom_left]
		_, ok_br := region_map[bottom_right]

		// top-left
		tl_cond := (!ok_left && !ok_top) || ((ok_top && ok_left) && !ok_tl)
		if tl_cond {
			sides++
		}

		// top-right
		tr_cond := (!ok_right && !ok_top) || ((ok_top && ok_right) && !ok_tr)
		if tr_cond {
			sides++
		}

		// bottom-left
		bl_cond := (!ok_left && !ok_bottom) || ((ok_bottom && ok_left) && !ok_bl)
		if bl_cond {
			sides++
		}

		// bottom-right
		br_cond := (!ok_right && !ok_bottom) || ((ok_bottom && ok_right) && !ok_br)
		if br_cond {
			sides++
		}

	}

	return sides
}

func d12_part_1(data string) (int, error) {
	grid := mapfunc(
		strings.Split(strings.TrimSpace(string(data)), "\n"),
		func(s string) []string {
			return strings.Split(s, "")
		})

	plants_seen := make(map[string][]Vector2)
	ans := 0
	for y, row := range grid {
		for x, c := range row {
			pos := Vector2{x, y}
			if !slices.Contains(plants_seen[c], pos) {
				region := fill_plant_region(grid, Vector2{x, y}, c)
				plants_seen[c] = append(plants_seen[c], region...)
				area := region_area(region)
				perimeter := region_perimeter(region)
				ans += perimeter * area
			}
		}
	}

	return ans, nil
}

func d12_part_2(data string) (int, error) {
	grid := mapfunc(
		strings.Split(strings.TrimSpace(string(data)), "\n"),
		func(s string) []string {
			return strings.Split(s, "")
		})

	plants_seen := make(map[string][]Vector2)
	ans := 0
	for y, row := range grid {
		for x, c := range row {
			pos := Vector2{x, y}
			if !slices.Contains(plants_seen[c], pos) {
				region := fill_plant_region(grid, Vector2{x, y}, c)
				plants_seen[c] = append(plants_seen[c], region...)
				area := region_area(region)
				sides := region_sides(region)
				ans += sides * area
			}
		}
	}

	return ans, nil
}
