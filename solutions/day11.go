package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day11() {
	const input_file = "inputs/day11.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 11 Solutions:\n")
	var start = time.Now()
	sol_1, err := d11_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day11 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d (%s)\n", sol_1, time.Since(start))

	start = time.Now()
	sol_2, err := d11_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day11 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d (%s)\n", sol_2, time.Since(start))
}

func blink(stone_map map[uint64]uint64) map[uint64]uint64 {
	new_stone_map := make(map[uint64]uint64)

	for stone, count := range stone_map {
		stone_str := strconv.FormatUint(stone, 10)
		if stone == 0 {
			map_entry_or_default_add(new_stone_map, 1, count)
		} else if len(stone_str)%2 == 0 {
			left := stone_str[:len(stone_str)/2]
			right := stone_str[len(stone_str)/2:]
			ls, _ := strconv.ParseUint(left, 10, 0)
			rs, _ := strconv.ParseUint(right, 10, 0)
			map_entry_or_default_add(new_stone_map, ls, count)
			map_entry_or_default_add(new_stone_map, rs, count)
		} else {
			map_entry_or_default_add(new_stone_map, stone*2024, count)
		}
	}

	return new_stone_map
}

func run_blinks(stones []uint64, n int) uint64 {
	stone_map := make(map[uint64]uint64)
	for _, stone := range stones {
		stone_map[stone] = 1
	}

	for range n {
		stone_map = blink(stone_map)
	}

	var result uint64
	result = 0
	for _, v := range stone_map {
		result += v
	}
	return result
}

func d11_part_1(data string) (uint64, error) {
	stones := mapfunc(
		strings.Split(
			strings.TrimSpace(string(data)),
			" "),
		func(s string) uint64 {
			v, _ := strconv.Atoi(s)
			return uint64(v)
		})

	ans := run_blinks(stones, 25)

	return ans, nil
}

func d11_part_2(data string) (uint64, error) {
	stones := mapfunc(
		strings.Split(
			strings.TrimSpace(string(data)),
			" "),
		func(s string) uint64 {
			v, _ := strconv.Atoi(s)
			return uint64(v)
		})

	ans := run_blinks(stones, 75)

	return ans, nil
}
