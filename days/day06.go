package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luc-10/Advent-of-code-2025/io"
)

func Day6() {
	Day6Part1()
	Day6Part2()
}

func Day6Part1() {
	data := io.ReadLines("inputFiles/day06.txt")
	worksheet, ops := getNumbersAndOps(data)
	table := transpose(worksheet)
	grandTotal := 0
	for i, op := range ops {
		grandTotal += applyOp(table[i], op)
	}
	fmt.Println(grandTotal)
}

func Day6Part2() {
	data := io.ReadLines("inputFiles/day06.txt")

	_, ops := getNumbersAndOps(data)
	grandTotal, opIndex := 0, 0
	nums := make([]int, 0)
	for j := range len(data[0]) {
		val := 0
		for i := range len(data) - 1 {
			if data[i][j] == ' ' {
				continue
			}
			val *= 10
			val += int(data[i][j] - '0')
		}
		if val == 0 {

			grandTotal += applyOp(nums, ops[opIndex])
			opIndex++
			nums = nums[:0]
		} else {
			nums = append(nums, val)
		}
	}
	grandTotal += applyOp(nums, ops[opIndex])
	fmt.Println(grandTotal)
}

func getNumbersAndOps(data []string) ([][]int, []string) {
	worksheet := make([][]int, len(data)-1)
	for i, s := range data[:len(data)-1] {
		line := strings.Fields(s)
		for _, num := range line {
			n, _ := strconv.Atoi(num)
			worksheet[i] = append(worksheet[i], n)
		}
	}

	return worksheet, strings.Fields(data[len(data)-1])
}

func transpose(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])
	tMat := make([][]int, cols)

	for i := range tMat {
		tMat[i] = make([]int, rows)
	}
	for i := range rows {
		for j := range cols {
			tMat[j][i] = mat[i][j]
		}
	}
	return tMat
}

func calc(table [][]int, index int, f func(a, b int) int) int {
	val := 1 - f(1, 0)
	for i := range table {
		val = f(val, table[i][index])
	}
	return val
}

func fold(arr []int, f func(a, b int) int) int {
	val := arr[0]
	for i := 1; i < len(arr); i++ {
		val = f(val, arr[i])
	}
	return val
}

func applyOp(arr []int, op string) int {
	switch op {
	case "+":
		return fold(arr, func(a, b int) int { return a + b })
	case "*":
		return fold(arr, func(a, b int) int { return a * b })
	}

	return 0
}
