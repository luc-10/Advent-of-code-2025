package days

import (
	"fmt"

	"github.com/luc-10/Advent-of-code-2025/io"
)

func Day4() {
	Day4Part1()
	Day4Part2()
}

func Day4Part1() {
	data := io.ReadByteLines("inputFiles/day04.txt")

	paperCount := 0

	for i := range data {
		for j := range data[i] {
			if data[i][j] == '@' && countAdjacent(data, i, j, '@') < 4 {
				paperCount++
			}
		}
	}
	fmt.Println(paperCount)
}

func Day4Part2() {
	data := io.ReadByteLines("inputFiles/day04.txt")

	paperCount, removedCount := 1, 0
	for paperCount > 0 {
		paperCount = 0
		for i := range data {
			for j := range data[i] {
				if data[i][j] == '@' && countAdjacent(data, i, j, '@') < 4 {
					paperCount++
					data[i][j] = 'x'
				}
			}
		}
		removedCount += paperCount
	}
	fmt.Println(removedCount)

}

func countAdjacent(mat [][]byte, x, y int, char byte) int {
	count := 0
	for i := max(0, x-1); i <= min(len(mat)-1, x+1); i++ {
		for j := max(0, y-1); j <= min(len(mat[i])-1, y+1); j++ {
			if i == x && j == y {
				continue
			}
			if mat[i][j] == char {
				count++
			}
		}
	}
	return count
}
