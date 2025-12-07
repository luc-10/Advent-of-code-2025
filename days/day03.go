package days

import (
	"fmt"

	"github.com/luc-10/Advent-of-code-2025/io"
)

func Day3() {
	Day3Part1()
	Day3Part2()
}

func Day3Part1() {
	data := io.ReadIntLines("inputFiles/day03.txt")

	joltageSum := 0
	for _, row := range data {
		maxNum, maxIndex, joltage := 0, 0, 0
		for j := 0; j < len(row)-1; j++ {

			if row[j] > maxNum {
				maxNum = row[j]
				maxIndex = j
			}
		}

		joltage += maxNum * 10
		maxNum = 0
		for k := maxIndex + 1; k < len(row); k++ {
			if row[k] > maxNum {
				maxNum = row[k]
			}
		}

		joltage += maxNum
		joltageSum += joltage
	}

	fmt.Println(joltageSum)
}

func Day3Part2() {
	data := io.ReadIntLines("inputFiles/day03.txt")
	joltageSum := 0
	for _, row := range data {
		dp := initializeMat(len(row), 12, -1)
		joltageSum += joltageMemoization(dp, len(row)-1, 11, row)
	}
	fmt.Println(joltageSum)

}

func initializeMat(rows, cols, val int) [][]int {
	mat := make([][]int, rows)

	for i := 0; i < rows; i++ {
		mat[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			mat[i][j] = val
		}
	}
	return mat
}

func joltageMemoization(dp [][]int, i, j int, row []int) int {
	if i < 0 || j < 0 {
		return 0
	}
	if dp[i][j] == -1 {
		dp[i][j] = max(row[i]+10*joltageMemoization(dp, i-1, j-1, row), joltageMemoization(dp, i-1, j, row))
	}
	return dp[i][j]
}
