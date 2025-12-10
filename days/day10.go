package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luc-10/Advent-of-code-2025/dataStructures"
	"github.com/luc-10/Advent-of-code-2025/io"
	"github.com/lukpank/go-glpk/glpk"
)

func Day10() {
	Day10Part1()
	Day10Part2()
}

func Day10Part1() {
	data := io.ReadLines("inputFiles/day10.txt")
	sum := 0
	for i := range data {
		configuration := strings.Split(data[i], " ")
		lights := make([]bool, len(configuration[0])-2)
		for i := range lights {
			if configuration[0][i+1] == '#' {
				lights[i] = true
			} else {
				lights[i] = false
			}
		}
		buttons := make([][]int, len(configuration)-2)
		for i := 1; i < len(configuration)-1; i++ {

			nums := strings.Split(strings.Trim(configuration[i], "()"), ",")
			for _, num := range nums {
				n, _ := strconv.Atoi(num)
				buttons[i-1] = append(buttons[i-1], n)
			}
		}

		sum += exploreButtons(lights, buttons)

	}
	fmt.Println(sum)
}

func Day10Part2() {

	data := io.ReadLines("inputFiles/day10.txt")
	sum := 0
	for i := range data {
		configuration := strings.Split(data[i], " ")
		lights := make([]bool, len(configuration[0])-2)
		for i := range lights {
			if configuration[0][i+1] == '#' {
				lights[i] = true
			} else {
				lights[i] = false
			}
		}
		buttons := make([][]int, len(configuration)-2)
		for i := 1; i < len(configuration)-1; i++ {

			nums := strings.Split(strings.Trim(configuration[i], "()"), ",")
			for _, num := range nums {
				n, _ := strconv.Atoi(num)
				buttons[i-1] = append(buttons[i-1], n)
			}
		}

		nums := strings.Split(strings.Trim(configuration[len(configuration)-1], "{}"), ",")
		joltage := make([]int, len(nums))
		for i := range nums {
			joltage[i], _ = strconv.Atoi(nums[i])
		}
		sum += solveGlpk(createMatrix(buttons, joltage), joltage)

	}
	fmt.Println(sum)
}

func exploreButtons(lights []bool, buttons [][]int) int {
	visited := make(map[string]bool)
	pos := make([]bool, len(lights))
	visited[turnIntoString(pos)] = true
	var q dataStructures.Queue[[]bool]
	q.Push(pos)
	dist, toPop := 0, 1
	for !q.Empty() {
		pos, _ := q.Pop()
		toPop--
		for _, button := range buttons {
			next := pressButton(pos, button)
			if !visited[turnIntoString(next)] {
				if equalLights(lights, next) {
					return dist + 1
				}
				visited[turnIntoString(next)] = true
				q.Push(next)
			}
		}

		if toPop == 0 {
			toPop = q.Length()
			dist++
		}
	}
	return -1
}

func turnIntoString(lights []bool) string {
	ret := make([]byte, len(lights))
	for _, b := range lights {
		if b {
			ret = append(ret, '1')
		} else {
			ret = append(ret, '0')
		}
	}
	return string(ret)
}

func pressButton(lights []bool, button []int) []bool {
	newLights := make([]bool, len(lights))
	copy(newLights, lights)
	for i := range button {
		newLights[button[i]] = !newLights[button[i]]
	}
	return newLights
}

func equalLights(lights1, lights2 []bool) bool {
	for i := range lights1 {
		if lights1[i] != lights2[i] {
			return false
		}
	}
	return true
}

func solveGlpk(A [][]int, b []int) int {
	rows := len(A)
	cols := len(A[0])

	lp := glpk.New()
	lp.SetProbName("day10Part2")
	lp.SetObjName("day10b")
	lp.SetObjDir(glpk.MIN)
	lp.AddCols(cols)
	for j := 0; j < cols; j++ {
		lp.SetColBnds(j+1, glpk.LO, 0, 0)
		lp.SetColKind(j+1, glpk.IV)
		lp.SetObjCoef(j+1, 1)
	}
	floatA := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		floatA[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			floatA[i][j] = float64(A[i][j])
		}
	}
	lp.AddRows(rows)
	for i := 1; i <= rows; i++ {
		lp.SetRowBnds(i, glpk.FX, float64(b[i-1]), float64(b[i-1]))

		colIdx := []int32{0}
		colVal := []float64{0}

		for j := 1; j <= cols; j++ {
			colIdx = append(colIdx, int32(j))
			colVal = append(colVal, float64(A[i-1][j-1]))

		}

		lp.SetMatRow(i, colIdx, colVal)
	}

	parm := glpk.NewSmcp()
	parm.SetMsgLev(glpk.MSG_OFF)
	lp.Simplex(parm)

	mipParm := glpk.NewIocp()
	mipParm.SetMsgLev(glpk.MSG_OFF)
	lp.Intopt(mipParm)
	xSol := make([]int, cols)
	for j := 1; j <= cols; j++ {
		xSol[j-1] = int(lp.MipColVal(j))
	}

	sum := 0
	for i := range xSol {
		sum += xSol[i]
	}
	lp.Delete()
	return sum

}

func createMatrix(buttons [][]int, joltage []int) [][]int {
	rows := len(joltage)
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, len(buttons))
	}
	for i := range buttons {
		for btn := range buttons[i] {
			matrix[buttons[i][btn]][i] = 1
		}
	}
	return matrix
}
