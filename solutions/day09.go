package solutions

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Block struct {
	pos  int
	size int
}

func print_disk(disk []int) {
	fmt.Println(strings.Join(mapfunc(disk, func(i int) string {
		if i == -1 {
			return "."
		}
		return string(i + '0')
	}), ""))
}

func Day09() {
	const input_file = "inputs/day09.txt"
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Day 09 Solutions:\n")
	sol_1, err := d09_part_1(string(data))
	if err != nil {
		fmt.Println("Error during Day9 part 1")
		return
	}
	fmt.Printf("\tPart 1: %d\n", sol_1)
	sol_2, err := d09_part_2(string(data))
	if err != nil {
		fmt.Println("Error during Day9 part 2")
		return
	}
	fmt.Printf("\tPart 2: %d\n", sol_2)
}

func d09_part_1(data string) (int, error) {
	input := strings.TrimSpace(string(data))

	var repr []int

	file_id := 0
	for i, block := range input {
		block_size := int(block - '0')
		if i%2 == 0 {
			for range block_size {
				repr = append(repr, file_id)
			}
			file_id++
		} else {
			for range block_size {
				repr = append(repr, -1)
			}
		}
	}

	start := 0
	end := len(repr) - 1
	for start < end {
		if repr[start] == -1 {
			if repr[end] != -1 {
				repr[start] = repr[end]
				repr[end] = -1
			}
			end--
		} else {
			start++
		}
	}

	ans := 0
	for pos, id := range repr {
		if id == -1 {
			break
		}
		ans += pos * id
	}

	return ans, nil
}

func d09_part_2(data string) (int, error) {
	input := strings.TrimSpace(string(data))

	var file_blocks []Block
	var free_blocks []Block
	var disk []int

	pos := 0
	file_id := 0
	for i, block := range input {
		block_size := int(block - '0')
		if block_size < 1 {
			continue
		}
		if i%2 == 0 {
			file_blocks = append(file_blocks, Block{
				pos:  pos,
				size: block_size,
			})
			for range block_size {
				disk = append(disk, file_id)
			}
			file_id++
		} else {
			free_blocks = append(free_blocks, Block{
				pos:  pos,
				size: block_size,
			})
			for range block_size {
				disk = append(disk, -1)
			}
		}
		pos += block_size
	}

	slices.Reverse(file_blocks)
	for _, file := range file_blocks {
		for free_idx, free := range free_blocks {
			if free.pos >= file.pos {
				break
			}
			if free.size >= file.size {
				for offset := 0; offset < file.size; offset++ {
					disk[free.pos+offset] = disk[file.pos+offset]
					disk[file.pos+offset] = -1
				}
				excess := free.size - file.size
				if excess > 0 {
					free_blocks[free_idx] = Block{
						pos:  free.pos + file.size,
						size: excess,
					}
				} else {
					free_blocks = remove(free_blocks, free_idx)
				}
				break
			}
		}
	}

	ans := 0
	for pos, id := range disk {
		if id == -1 {
			continue
		}
		ans += pos * id
	}

	return ans, nil
}
