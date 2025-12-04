package main

import (
	"fmt"
	"os"

	"github.com/luc-10/Advent-of-code-2025/days"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
	}

	day := os.Args[1]

	stringToFuncMap := map[string]func(){

		"day1":      days.Day1,
		"day1Part1": days.Day1Part1,
		"day1Part2": days.Day1Part2,

		"day2":      days.Day2,
		"day2Part1": days.Day2Part1,
		"day2Part2": days.Day2Part2,

		"day3":      days.Day3,
		"day3Part1": days.Day3Part1,
		"day3Part2": days.Day3Part2,

		"day4":      days.Day4,
		"day4Part1": days.Day4Part1,
		"day4Part2": days.Day4Part2,
	}

	if f, ok := stringToFuncMap[day]; ok {
		f()
	} else {
		fmt.Println("Day not valid:", day)
	}
}
