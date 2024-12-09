package solutions

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X int
	Y int
}

func (p *Vector2) add_vector(other *Vector2) Vector2 {
	return Vector2{p.X + other.X, p.Y + other.Y}
}

func (p *Vector2) sub_vector(other *Vector2) Vector2 {
	return Vector2{p.X - other.X, p.Y - other.Y}
}

func (p *Vector2) distance(other *Vector2) float64 {
	return math.Sqrt(
		math.Pow(float64(p.X)-float64(other.X), 2) +
			math.Pow(float64(p.Y)-float64(other.Y), 2))
}

func mapfunc[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))
	for _, e := range data {
		res = append(res, f(e))
	}
	return res
}

func is_digit(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	}
	return false
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func print_grid(grid [][]string) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Print(char, " ")
		}
		fmt.Println()
	}
}

type addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | float32 | float64
}

func map_entry_or_default_add[K comparable, V addable](hmap map[K]V, key K, value V) {
	if _, ok := hmap[key]; ok {
		hmap[key] += value
	} else {
		hmap[key] = value
	}
}

func print_reset()       { fmt.Print("\033[0m") }
func print_clear()       { fmt.Print("\033[H\033[2J") }
func print_hide_cursor() { fmt.Print("\033[?25l") }
func print_show_cursor() { fmt.Print("\033[?25h") }
func print_bold()        { fmt.Print("\033[1m") }
func print_red()         { fmt.Print("\033[31m") }
func print_green()       { fmt.Print("\033[32m") }
func print_blue()        { fmt.Print("\033[34m") }
