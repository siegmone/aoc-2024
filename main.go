package main

import (
	"aoc/solutions"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	argv := os.Args
	var day int
	if argc := len(argv); argc == 1 {
		day = 0
	} else {
		arg1, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error: '%s' not a valid argument.\n", os.Args[1])
		}
		day = int(arg1)
	}

	if day == 0 {
		runAll()
	} else {
		runDay(day)
	}
}

func downloadInput(filepath string, url string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	sessionCookie, exists := os.LookupEnv("AOC_SESSION_COOKIE")
	if !exists {
		return fmt.Errorf("No session cookie provided. Could't download the input")
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("writing to file: %w", err)
	}
	return nil
}

func runDay(day int) {
	input_file := fmt.Sprintf("inputs/day%02d.txt", day)
	if _, err := os.Stat(input_file); errors.Is(err, os.ErrNotExist) {
		url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
		if err := downloadInput(input_file, url); err != nil {
			fmt.Println(err)
		}
	}

	switch day {
	case 1:
		solutions.Day01()
	case 2:
		solutions.Day02()
	case 3:
		solutions.Day03()
	case 4:
		solutions.Day04()
	case 5:
		solutions.Day05()
	case 6:
		solutions.Day06()
	case 7:
		solutions.Day07()
	case 8:
		solutions.Day08()
	case 9:
		solutions.Day09()
	case 10:
		solutions.Day10()
	case 11:
		solutions.Day11()
	case 12:
		solutions.Day12()
	case 13:
		solutions.Day13()
	case 14:
		solutions.Day14()
	case 15:
		solutions.Day15()
	case 16:
		solutions.Day16()
	case 17:
		solutions.Day17()
	case 18:
		solutions.Day18()
	case 19:
		solutions.Day19()
	case 20:
		solutions.Day20()
	case 21:
		solutions.Day21()
	case 22:
		solutions.Day22()
	case 23:
		solutions.Day23()
	case 24:
		solutions.Day24()
	case 25:
		solutions.Day25()
	default:
		fmt.Println("Invalid day, please specify between 1 and 25.")
	}
}

func runAll() {
	for i := 1; i <= 25; i++ {
		runDay(i)
	}
}
