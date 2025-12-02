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
