package solutions

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Day08() {
	const input_file = "inputs/day08.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 08 Solutions:\n")
	var start = time.Now()
	sol_1, err := d08_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day08 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d (%s)\n", sol_1, time.Since(start))

	start = time.Now()
	sol_2, err := d08_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day08 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d (%s)\n", sol_2, time.Since(start))
}

func d08_part_1(data string) (int, error) {
	grid := mapfunc(strings.Split(strings.TrimSpace(string(data)), "\n"), func(s string) []string {
		return strings.Split(s, "")
	})
	antenna_positions := make(map[string][]Vector2)
	for y, row := range grid {
		for x, c := range row {
			if c != "." {
				antenna_positions[c] = append(antenna_positions[c], Vector2{X: x, Y: y})
			}
		}
	}

	height := len(grid)
	width := len(grid[0])
	anti_nodes := make([][]int, height)
	for i := range anti_nodes {
		anti_nodes[i] = make([]int, width)
	}

	grid_cpy := make([][]string, height)
	for i := range grid_cpy {
		grid_cpy[i] = make([]string, width)
	}
	copy(grid_cpy, grid)
	fmt.Println()

	ans := 0
	for _, antennas := range antenna_positions {
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a1 == a2 {
					continue
				}
				diff := a1.sub_vector(&a2)
				an1 := a1.add_vector(&diff)
				an2 := a2.sub_vector(&diff)
				if an1.X >= 0 && an1.X < width && an1.Y >= 0 && an1.Y < height {
					if anti_nodes[an1.Y][an1.X] != 1 {
						ans++
						anti_nodes[an1.Y][an1.X] = 1
						grid_cpy[an1.Y][an1.X] = "#"
					}
				}
				if an2.X >= 0 && an2.X < width && an2.Y >= 0 && an2.Y < height {
					if anti_nodes[an2.Y][an2.X] != 1 {
						ans++
						anti_nodes[an2.Y][an2.X] = 1
						grid_cpy[an2.Y][an2.X] = "#"
					}
				}

				print_hide_cursor()
				print_grid(grid_cpy)
				for range height {
					fmt.Print("\033[A")
				}
				time.Sleep(10 * time.Millisecond)
				print_show_cursor()
			}
		}
	}
	for range height {
		fmt.Print("\033[B")
	}
	fmt.Println()

	return ans, nil
}

func d08_part_2(data string) (int, error) {
	grid := mapfunc(strings.Split(strings.TrimSpace(string(data)), "\n"), func(s string) []string {
		return strings.Split(s, "")
	})
	antenna_positions := make(map[string][]Vector2)
	for y, row := range grid {
		for x, c := range row {
			if c != "." {
				antenna_positions[c] = append(antenna_positions[c], Vector2{X: x, Y: y})
			}
		}
	}

	height := len(grid)
	width := len(grid[0])
	anti_nodes := make([][]int, height)
	for i := range anti_nodes {
		anti_nodes[i] = make([]int, width)
	}

	grid_cpy := make([][]string, height)
	for i := range grid_cpy {
		grid_cpy[i] = make([]string, width)
	}
	copy(grid_cpy, grid)
	fmt.Println()

	ans := 0
	for _, antennas := range antenna_positions {
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a1 == a2 {
					continue
				}
				step := a2.sub_vector(&a1)
				current := a1
				for current.X >= 0 && current.X < width && current.Y >= 0 && current.Y < height {
					if anti_nodes[current.Y][current.X] != 1 {
						ans++
						anti_nodes[current.Y][current.X] = 1
						grid_cpy[current.Y][current.X] = "#"
					}
					current = current.add_vector(&step)

					print_hide_cursor()
					print_grid(grid_cpy)
					for range height {
						fmt.Print("\033[A")
					}
					time.Sleep(10 * time.Millisecond)
					print_show_cursor()
				}
			}
		}
	}
	for range height {
		fmt.Print("\033[B")
	}
	fmt.Println()

	return ans, nil
}
