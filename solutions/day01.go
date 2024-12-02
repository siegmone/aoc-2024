package solutions

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day01() {
	const input_file = "inputs/day01.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 01 Solutions:\n")
	sol_1, err := d1_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day01 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d\n", sol_1)
	sol_2, err := d1_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day01 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d\n", sol_2)
}

func d1_part_1(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	l1 := []int{}
	l2 := []int{}
	for _, line := range lines {
		str_values := strings.Split(line, "   ")
		v1, err := strconv.Atoi(str_values[0])
		if err != nil {
			fmt.Println("Malformed input")
			return 0, err
		}
		l1 = append(l1, v1)

		v2, err := strconv.Atoi(str_values[1])
		if err != nil {
			fmt.Println("Malformed input")
			return 0, err
		}
		l2 = append(l2, v2)
	}

	sort.Slice(l1, func(i, j int) bool { return l1[i] < l1[j] })
	sort.Slice(l2, func(i, j int) bool { return l2[i] < l2[j] })

	sum := 0
	if len(l1) == len(l2) {
		for i := range l1 {
			diff := int(math.Abs(float64(l1[i] - l2[i])))
			sum += diff
		}
	}

	return sum, nil
}

func d1_part_2(data string) (int, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	l1 := []int{}
	var m map[int]int
	m = make(map[int]int)
	for _, line := range lines {
		str_values := strings.Split(line, "   ")

		v1, err := strconv.Atoi(str_values[0])
		if err != nil {
			fmt.Println("Malformed input 1")
			return 0, err
		}
		l1 = append(l1, v1)

		v2, err := strconv.Atoi(str_values[1])
		if err != nil {
			fmt.Println("Malformed input 2")
			return 0, err
		}

		if _, ok := m[v2]; ok {
			m[v2]++
		} else {
			m[v2] = 1
		}
	}

    sum := 0
    for _, v := range l1 {
        sum += v * m[v]
    }

	return sum, nil
}
