package solutions

import (
	"fmt"
)

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

func print_reset()       { fmt.Print("\033[0m") }
func print_clear()       { fmt.Print("\033[H\033[2J") }
func print_hide_cursor() { fmt.Print("\033[?25l") }

func print_bold() { fmt.Print("\033[1m") }

func print_red()   { fmt.Print("\033[31m") }
func print_green() { fmt.Print("\033[32m") }
func print_blue()  { fmt.Print("\033[34m") }
