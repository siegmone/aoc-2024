package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Robot struct {
	pos Vector2
	vel Vector2
}

func Day14() {
	const input_file = "inputs/day14.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 14 Solutions:\n")
	var start = time.Now()
	sol_1, err := d14_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day14 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d (%s)\n", sol_1, time.Since(start))

	start = time.Now()
	sol_2, err := d14_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day14 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d (%s)\n", sol_2, time.Since(start))
}

func parse_robot(line string) Robot {
	split := strings.Split(strings.TrimSpace(line), " ")
	pos_str := split[0]
	vel_str := split[1]
	split = strings.Split(pos_str, "=")
	split = strings.Split(split[1], ",")
	p_x, _ := strconv.Atoi(split[0])
	p_y, _ := strconv.Atoi(split[1])
	split = strings.Split(vel_str, "=")
	split = strings.Split(split[1], ",")
	v_x, _ := strconv.Atoi(split[0])
	v_y, _ := strconv.Atoi(split[1])
	return Robot{
		pos: Vector2{p_x, p_y},
		vel: Vector2{v_x, v_y},
	}
}

func update_robot(robot *Robot, x_lim, y_lim int) {
	robot.pos = robot.pos.add_vector(&robot.vel)
	if robot.pos.Y >= y_lim {
		robot.pos.Y = robot.pos.Y - y_lim
	} else if robot.pos.Y < 0 {
		robot.pos.Y += y_lim
	}
	if robot.pos.X >= x_lim {
		robot.pos.X = robot.pos.X - x_lim
	} else if robot.pos.X < 0 {
		robot.pos.X += x_lim
	}
}

func grid_check() {
	grid_str := "......2..1. ........... 1.......... .11........ .....1..... ...12...... .1....1...."
	grid := mapfunc(
		strings.Split(strings.TrimSpace(grid_str), " "),
		func(s string) []string {
			return strings.Split(s, "")
		})
	var robots []Vector2
	for y, row := range grid {
		for x, cell := range row {
			if cell == "." {
				continue
			}
			val, _ := strconv.Atoi(cell)
			for range val {
				robots = append(robots, Vector2{x, y})
			}
		}
	}
	fmt.Println(robots)
}

func d14_part_1(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var robots []Robot
	for _, line := range lines {
		robot := parse_robot(line)
		robots = append(robots, robot)
	}

	width := 101
	height := 103
	steps := 100
	for range steps {
		for i := range robots {
			update_robot(&robots[i], width, height)
			if robots[i].pos.X < 0 ||
				robots[i].pos.Y < 0 ||
				robots[i].pos.X >= width ||
				robots[i].pos.Y >= height {
				panic("Something strange happened")
			}
		}
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		x := robot.pos.X
		y := robot.pos.Y
		left := x < (width-1)/2
		right := x > (width-1)/2
		up := y < (height-1)/2
		down := y > (height-1)/2
		if right && up {
			q1 += 1
		}
		if left && up {
			q2 += 1
		}
		if left && down {
			q3 += 1
		}
		if right && down {
			q4 += 1
		}
	}

	ans := q1 * q2 * q3 * q4

	return ans, nil
}

func d14_part_2(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var robots []Robot
	for _, line := range lines {
		robot := parse_robot(line)
		robots = append(robots, robot)
	}

	input := bufio.NewScanner(os.Stdin)

	width := 101
	height := 103

	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
	}

	var x_mean, y_mean, x_var, y_var, variance float64
	print_hide_cursor()
	for frame := 0; ; frame++ {
		for y, row := range grid {
			for x := range row {
				grid[y][x] = " "
			}
		}

		x_mean = 0
		y_mean = 0
		x_var = 0
		y_var = 0
		for i := range robots {
			update_robot(&robots[i], width, height)
			if robots[i].pos.X < 0 ||
				robots[i].pos.Y < 0 ||
				robots[i].pos.X >= width ||
				robots[i].pos.Y >= height {
				panic("Something strange happened")
			}
			grid[robots[i].pos.Y][robots[i].pos.X] = "X"
			x_mean += float64(robots[i].pos.X)
			y_mean += float64(robots[i].pos.Y)
		}
		x_mean = x_mean / float64(len(robots))
		y_mean = x_mean / float64(len(robots))

		for _, robot := range robots {
			x_var += math.Pow(float64(robot.pos.X)-x_mean, 2)
			y_var += math.Pow(float64(robot.pos.Y)-y_mean, 2)
		}
		x_var = x_var / float64(len(robots))
		y_var = x_var / float64(len(robots))

		variance = (x_var + y_var) / 2

		if variance < 200 {
			print_clear()
			fmt.Println("Frame:", frame)
			fmt.Println("Mean X:", x_mean)
			fmt.Println("Mean Y:", y_mean)
			fmt.Println("Variance:", variance)
			fmt.Println("Variance X:", x_var)
			fmt.Println("Variance Y:", y_var)
			print_grid(grid)
			input.Scan()
		}

	}

	ans := 0
	return ans, nil
}
