package days

import (
	"fmt"

	"github.com/luc-10/Advent-of-code-2025/dataStructures"
	"github.com/luc-10/Advent-of-code-2025/io"
)

func Day7() {
	Day7Part1()
	Day7Part2()
}

func Day7Part1() {
	data := io.ReadByteLines("inputFiles/day07.txt")

	var q dataStructures.Queue[[2]int]
	q.Push(locateInMatrix(data, 'S'))

	splits := 0
	for !q.Empty() {
		pos, _ := q.Pop()
		if data[pos[0]][pos[1]] == '|' {
			continue
		} else {
			data[pos[0]][pos[1]] = '|'
		}
		if pos[0]+1 >= len(data) {
			continue
		} else if data[pos[0]+1][pos[1]] == '.' {
			q.Push([2]int{pos[0] + 1, pos[1]})
		} else if data[pos[0]+1][pos[1]] == '^' {
			splits++
			q.Push([2]int{pos[0] + 1, pos[1] - 1})
			q.Push([2]int{pos[0] + 1, pos[1] + 1})
		}
	}
	fmt.Println(splits)
}

func Day7Part2() {
	data := io.ReadByteLines("inputFiles/day07.txt")

	fmt.Println(tachyonMemoization(data, locateInMatrix(data, 'S'), make(map[[2]int]int)))
}

func locateInMatrix(mat [][]byte, char byte) [2]int {
	for i, row := range mat {
		for j, c := range row {
			if c == char {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

func tachyonMemoization(mat [][]byte, pos [2]int, dp map[[2]int]int) int {
	if pos[0] >= len(mat) {
		return 1
	}
	if dp[pos] == 0 {
		if mat[pos[0]][pos[1]] == '^' {
			dp[pos] += tachyonMemoization(mat, [2]int{pos[0], pos[1] - 1}, dp) + tachyonMemoization(mat, [2]int{pos[0], pos[1] + 1}, dp)
		} else {
			dp[pos] += tachyonMemoization(mat, [2]int{pos[0] + 1, pos[1]}, dp)
		}
	}
	return dp[pos]
}
