package days

import (
	"fmt"
	"strconv"

	"github.com/luc-10/Advent-of-code-2025/io"
)

func Day1() {
	Day1Part1()
	Day1Part2()
}

func Day1Part1() {
	data := io.ReadLines("inputFiles/day01.txt")

	zeroCount, position := 0, 50

	for _, s := range data {
		movement, err := strconv.Atoi(s[1:])

		if err != nil {
			panic(err)
		}

		if s[0] == 'L' {
			movement *= -1
		}
		position += movement
		position = ((position % 100) + 100) % 100
		if position == 0 {
			zeroCount++
		}

	}

	fmt.Println(zeroCount)
}

func Day1Part2() {
	data := io.ReadLines("inputFiles/day01.txt")

	zeroCount, position := 0, 50

	for _, s := range data {
		movement, err := strconv.Atoi(s[1:])

		if err != nil {
			panic(err)
		}

		direction := 1

		if s[0] == 'L' {
			direction = -1
		}

		for movement > 0 {
			position += direction
			movement--
			if position == -1 {
				position += 100
			}
			if position == 100 {
				position -= 100
			}
			if position == 0 {
				zeroCount++
			}
		}

	}

	fmt.Println(zeroCount)
}
