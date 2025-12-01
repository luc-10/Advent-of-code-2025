package io

import (
	"os"
	"strings"
)

func ReadLines(path string)[]string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	text := strings.TrimSpace(string(input))
	return strings.Split(text, "\n")
}