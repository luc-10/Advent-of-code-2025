package io

import (
	"os"
	"strings"
)

func ReadLines(path string) []string {
	return strings.Split(ReadLine(path), "\n")
}

func ReadLine(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(input))
}

func ReadDigitLines(path string) [][]int {
	lines := ReadLines(path)
	mat := make([][]int, len(lines))

	for i, line := range lines {
		row := make([]int, len(line))
		for j, char := range line {
			row[j] = int(char - '0')
		}
		mat[i] = row
	}
	return mat
}
