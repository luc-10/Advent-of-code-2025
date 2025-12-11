package main

import (
	"fmt"
	"os"
	"time"

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

		"day5":      days.Day5,
		"day5Part1": days.Day5Part1,
		"day5Part2": days.Day5Part2,

		"day6":      days.Day6,
		"day6Part1": days.Day6Part1,
		"day6Part2": days.Day6Part2,

		"day7":      days.Day7,
		"day7Part1": days.Day7Part1,
		"day7Part2": days.Day7Part2,

		"day8":      days.Day8,
		"day8Part1": days.Day8Part1,
		"day8Part2": days.Day8Part2,

		"day9":      days.Day9,
		"day9Part1": days.Day9Part1,
		"day9Part2": days.Day9Part2,

		"day10":      days.Day10,
		"day10Part1": days.Day10Part1,
		"day10Part2": days.Day10Part2,

		"day11":      days.Day11,
		"day11Part1": days.Day11Part1,
		"day11Part2": days.Day11Part2,
	}

	if f, ok := stringToFuncMap[day]; ok {
		start := time.Now()
		f()
		elapsed := time.Since(start)
		fmt.Println("Time:", elapsed)
	} else {
		fmt.Println("Day not valid:", day)
	}
}
