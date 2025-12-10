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
		lights, buttons, _ := getLightsButtonsJoltage(data[i])
		sum += buttonsBFS(lights, buttons)
	}

	fmt.Println(sum)
}

func Day10Part2() {
	data := io.ReadLines("inputFiles/day10.txt")

	sum := 0
	for i := range data {
		_, buttons, joltage := getLightsButtonsJoltage(data[i])
		sum += solveGlpk(createMatrix(buttons, len(joltage)), joltage)

	}

	fmt.Println(sum)
}

func getLightsButtonsJoltage(line string) ([]bool, [][]int, []int) {
	splitLine := strings.Split(line, " ")
	lights := make([]bool, len(splitLine[0])-2)
	for i := range lights {
		if splitLine[0][i+1] == '#' {
			lights[i] = true
		} else {
			lights[i] = false
		}
	}
	buttons := make([][]int, len(splitLine)-2)
	for i := 1; i < len(splitLine)-1; i++ {

		buttonString := strings.Split(strings.Trim(splitLine[i], "()"), ",")
		for _, s := range buttonString {
			num, _ := strconv.Atoi(s)
			buttons[i-1] = append(buttons[i-1], num)
		}
	}

	joltageString := strings.Split(strings.Trim(splitLine[len(splitLine)-1], "{}"), ",")
	joltage := make([]int, len(joltageString))
	for i := range joltageString {
		joltage[i], _ = strconv.Atoi(joltageString[i])
	}
	return lights, buttons, joltage
}

func buttonsBFS(lights []bool, buttons [][]int) int {
	visited := make(map[string]bool)
	pos := make([]bool, len(lights))
	visited[turnLightsToString(pos)] = true
	var q dataStructures.Queue[[]bool]
	q.Push(pos)
	dist, toPop := 0, 1
	for !q.Empty() {
		pos, _ := q.Pop()
		toPop--
		for _, button := range buttons {
			next := pressButton(pos, button)
			if !visited[turnLightsToString(next)] {
				if equalLights(lights, next) {
					return dist + 1
				}
				visited[turnLightsToString(next)] = true
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

func turnLightsToString(lights []bool) string {
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
	lp.SetObjDir(glpk.MIN)

	lp.AddCols(cols)
	for j := 1; j <= cols; j++ {
		lp.SetColBnds(j, glpk.LO, 0, 0)
		lp.SetColKind(j, glpk.IV)
		lp.SetObjCoef(j, 1)
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

	sum := 0
	for j := 1; j <= cols; j++ {
		sum += int(lp.MipColVal(j))
	}
	lp.Delete()

	return sum

}

func createMatrix(buttons [][]int, rows int) [][]int {
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
