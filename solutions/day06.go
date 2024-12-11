package solutions

/* With animation!!! */

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Guard struct {
	X         int
	Y         int
	Direction int /* 0 ^ | 1 > | 2 v | 3 < */
}

type End int

const (
	WALL End = iota
	FINISH
	CLEAR
)

var directions [4]string = [4]string{"^", ">", "v", "<"}
var animate = false

func Day06() {
	input_file := "inputs/day06.txt"
	if animate {
		input_file = "inputs/day06_ex.txt"
	}
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	ani_txt := ""
	if animate {
		ani_txt = " (with animations)"
	}

	fmt.Printf("Day 06 Solutions%s:\n", ani_txt)
	var start = time.Now()
	sol_1, err := d06_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day06 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d (%s)\n", sol_1, time.Since(start))

	start = time.Now()
	sol_2, err := d06_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day06 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d (%s)\n", sol_2, time.Since(start))
}

func turn_right(g *Guard) {
	g.Direction = (g.Direction + 1) % 4
}

func guard_move(g *Guard, grid [][]string) End {
	height := len(grid)
	width := len(grid[0])
	switch g.Direction {
	case 0: // ^
		{
			if g.Y == 0 {
				return FINISH
			} else if grid[g.Y-1][g.X] == "#" || grid[g.Y-1][g.X] == "O" {
				turn_right(g)
				return WALL
			} else {
				g.Y--
			}
		}
	case 1: // >
		{
			if g.X == width-1 {
				return FINISH
			} else if grid[g.Y][g.X+1] == "#" || grid[g.Y][g.X+1] == "O" {
				turn_right(g)
				return WALL
			} else {
				g.X++
			}
		}
	case 2: // v
		{
			if g.Y == height-1 {
				return FINISH
			} else if grid[g.Y+1][g.X] == "#" || grid[g.Y+1][g.X] == "O" {
				turn_right(g)
				return WALL
			} else {
				g.Y++
			}
		}
	case 3: // <
		{
			if g.X == 0 {
				return FINISH
			}
			if grid[g.Y][g.X-1] == "#" || grid[g.Y][g.X-1] == "O" {
				turn_right(g)
				return WALL
			} else {
				g.X--
			}
		}
	default:
		panic("WTF")
	}
	return CLEAR
}

func print_guard_grid(grid [][]string) {
	width := len(grid[0])
	for range 2*width + 2 {
		fmt.Print("*")
	}
	fmt.Println()

	for _, row := range grid {
		fmt.Print("*")
		for _, char := range row {
			if char == "^" || char == ">" || char == "v" || char == "<" {
				print_bold()
				print_green()
				fmt.Print(char, " ")
				print_reset()
				continue
			}
			if char == "#" {
				print_bold()
				print_red()
				fmt.Print(char, " ")
				print_reset()
				continue
			}
			if char == "O" {
				print_bold()
				print_blue()
				fmt.Print(char, " ")
				print_reset()
				continue
			}
			fmt.Print("  ")
		}
		fmt.Print("*")
		fmt.Println()
	}
	for range 2*width + 2 {
		fmt.Print("*")
	}
	fmt.Println()
}

func animate_guard(g *Guard, grid [][]string, part int) {
	grid[g.Y][g.X] = directions[g.Direction]
	print_hide_cursor()
	print_bold()
	fmt.Printf("Part %d\n", part)
	print_reset()
	print_guard_grid(grid)
	grid[g.Y][g.X] = "."
	height := len(grid)
	for range height + 3 {
		fmt.Print("\033[A")
	}
	time.Sleep(10 * time.Millisecond)
	print_show_cursor()
}

func d06_part_1(data string) (int, error) {
	grid := mapfunc(strings.Split(strings.TrimSpace(string(data)), "\n"), func(s string) []string {
		return strings.Split(s, "")
	})

	guard := Guard{X: 0, Y: 0, Direction: 0}

	for y, row := range grid {
		for x, pos := range row {
			if pos == "^" {
				guard.X = x
				guard.Y = y
				guard.Direction = 0
			}
		}
	}

	height := len(grid)
	width := len(grid[0])

	visited := make([][]int, height)
	for i := range visited {
		visited[i] = make([]int, width)
	}

	ans := 1
	visited[guard.Y][guard.X] = 1

	for {
		what := guard_move(&guard, grid)
		if animate {
			animate_guard(&guard, grid, 1)
		}
		if visited[guard.Y][guard.X] != 1 {
			ans++
			visited[guard.Y][guard.X] = 1
		}
		if what == FINISH {
			break
		}
	}

	for range height + 3 {
		fmt.Print("\033[B")
	}

	return ans, nil
}

func d06_part_2(data string) (int, error) {
	grid := mapfunc(strings.Split(strings.TrimSpace(string(data)), "\n"), func(s string) []string {
		return strings.Split(s, "")
	})

	height := len(grid)

	guard := Guard{X: 0, Y: 0, Direction: 0}
	var start_pos Vector2

	for y, row := range grid {
		for x, pos := range row {
			if pos == "^" {
				start_pos.X = x
				start_pos.Y = y
			}
		}
	}

	var visited []Vector2
	guard.X = start_pos.X
	guard.Y = start_pos.Y

	for {
		what := guard_move(&guard, grid)
		if animate {
			animate_guard(&guard, grid, 2)
		}
		pos := Vector2{guard.X, guard.Y}
		if !slices.Contains(visited, pos) && pos != start_pos {
			visited = append(visited, pos)
		}
		if what == FINISH {
			break
		}
	}

	ans := 0
	var wall_map map[Vector2]int
	wall_map = make(map[Vector2]int)

	for _, v := range visited {
		// reset the guard
		guard.X = start_pos.X
		guard.Y = start_pos.Y
		guard.Direction = 0
		grid[v.Y][v.X] = "O"
		for {
			what := guard_move(&guard, grid)
			if animate {
				animate_guard(&guard, grid, 2)
			}
			if what == WALL {
				pos := Vector2{X: guard.X, Y: guard.Y}
				if _, ok := wall_map[pos]; ok {
					wall_map[pos]++
					if wall_map[pos] > 2 {
						ans++
						break
					}
				} else {
					wall_map[pos] = 1
				}
			}
			if what == FINISH {
				break
			}
		}
		grid[v.Y][v.X] = "."
		for p := range wall_map {
			delete(wall_map, p)
		}
	}

	for range height + 3 {
		fmt.Print("\033[B")
	}

	return ans, nil
}
