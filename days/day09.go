package days

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/luc-10/Advent-of-code-2025/dataStructures"
	"github.com/luc-10/Advent-of-code-2025/io"
)

func Day9() {
	Day9Part1()
	Day9Part2()
}

func Day9Part1() {
	data := io.ReadLines("inputFiles/day09.txt")
	redTilesCoords := getRedTilesCoords(data)
	rectangleSize := 0

	for i := range redTilesCoords {
		for j := i + 1; j < len(redTilesCoords); j++ {
			rectangleSize = max(rectangleSize, getRectangleSize(redTilesCoords[i], redTilesCoords[j]))

		}
	}
	fmt.Println(rectangleSize)
}

func Day9Part2() {
	data := io.ReadLines("inputFiles/day09.txt")
	redTilesCoords := getRedTilesCoords(data)
	initialOrder := make(map[[2]int]int)
	for i := range redTilesCoords {
		initialOrder[redTilesCoords[i]] = i
	}

	sort.Slice(redTilesCoords, func(i, j int) bool {
		return redTilesCoords[i][0] < redTilesCoords[j][0]
	})
	bigifyX := make(map[int]int)
	smallifyX := make(map[int]int)
	val := 1
	for _, point := range redTilesCoords {
		if smallifyX[point[0]] == 0 {
			smallifyX[point[0]] = val
			bigifyX[val] = point[0]
			val++
		}
	}

	maxX := val
	sort.Slice(redTilesCoords, func(i, j int) bool {
		return redTilesCoords[i][1] < redTilesCoords[j][1]
	})

	bigifyY := make(map[int]int)
	smallifyY := make(map[int]int)
	val = 1

	for _, point := range redTilesCoords {
		if smallifyY[point[1]] == 0 {
			smallifyY[point[1]] = val
			bigifyY[val] = point[1]
			val++
		}
	}

	maxY := val
	sort.Slice(redTilesCoords, func(i, j int) bool {
		return initialOrder[redTilesCoords[i]] < initialOrder[redTilesCoords[j]]
	})
	smallRedTilesCoords := make([][2]int, len(redTilesCoords))
	for i := range redTilesCoords {
		smallRedTilesCoords[i] = smallify(redTilesCoords[i], smallifyX, smallifyY)
	}

	grid := createGrid([2]int{maxX, maxY})
	createBorder(grid, smallRedTilesCoords)
	fillOutside(grid)
	rectangleSize := 0
	for i := range smallRedTilesCoords {
		for j := i + 2; j < len(smallRedTilesCoords); j++ {

			if isInside(smallRedTilesCoords[i], smallRedTilesCoords[j], grid) {
				rectangleSize = max(rectangleSize, getBigRectangleSize(smallRedTilesCoords[i], smallRedTilesCoords[j], bigifyX, bigifyY))
			}

		}
	}
	fmt.Println(rectangleSize)
}

func getRedTilesCoords(lines []string) [][2]int {
	redTilesList := make([][2]int, len(lines))
	for i, line := range lines {
		redTilesCoords := strings.Split(line, ",")
		redTilesList[i][0], _ = strconv.Atoi(redTilesCoords[0])
		redTilesList[i][1], _ = strconv.Atoi(redTilesCoords[1])
	}
	return redTilesList
}

func getRectangleSize(rect1, rect2 [2]int) int {
	return int((math.Abs(float64(rect2[0]-rect1[0])) + 1) * (math.Abs(float64(rect2[1]-rect1[1])) + 1))
}

func smallify(point [2]int, smallifyX, smallifyY map[int]int) [2]int {
	return [2]int{smallifyX[point[0]], smallifyY[point[1]]}
}

func createGrid(corner [2]int) [][]byte {
	grid := make([][]byte, corner[0]+2)
	for i := range len(grid) {
		grid[i] = make([]byte, corner[1]+2)
	}
	return grid
}

func createBorder(grid [][]byte, redTilesList [][2]int) {
	for i := range redTilesList {
		p1, p2 := redTilesList[i], redTilesList[(i+1)%len(redTilesList)]
		dx, dy := sign(p2[0]-p1[0]), sign(p2[1]-p1[1])
		for p1 != p2 {
			grid[p1[0]][p1[1]] = 'B'
			p1[0] += dx
			p1[1] += dy
		}
	}
}

func sign(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}

func fillOutside(grid [][]byte) {
	var q dataStructures.Queue[[2]int]
	q.Push([2]int{0, 0})
	grid[0][0] = 'O'
	for !q.Empty() {
		top, _ := q.Pop()
		if top[0]-1 >= 0 && grid[top[0]-1][top[1]] == byte(0) {
			grid[top[0]-1][top[1]] = 'O'
			q.Push([2]int{top[0] - 1, top[1]})
		}
		if top[1]-1 >= 0 && grid[top[0]][top[1]-1] == byte(0) {
			grid[top[0]][top[1]-1] = 'O'
			q.Push([2]int{top[0], top[1] - 1})
		}
		if top[0]+1 < len(grid) && grid[top[0]+1][top[1]] == byte(0) {
			grid[top[0]+1][top[1]] = 'O'
			q.Push([2]int{top[0] + 1, top[1]})
		}
		if top[1]+1 < len(grid[top[0]]) && grid[top[0]][top[1]+1] == byte(0) {
			grid[top[0]][top[1]+1] = 'O'
			q.Push([2]int{top[0], top[1] + 1})
		}
	}
}

func getBigRectangleSize(rect1, rect2 [2]int, bigifyX, bigifyY map[int]int) int {
	return getRectangleSize([2]int{bigifyX[rect1[0]], bigifyY[rect1[1]]}, [2]int{bigifyX[rect2[0]], bigifyY[rect2[1]]})
}

func isInside(point1, point2 [2]int, grid [][]byte) bool {
	for i := min(point1[0], point2[0]); i <= max(point1[0], point2[0]); i++ {
		for j := min(point1[1], point2[1]); j <= max(point1[1], point2[1]); j++ {
			if grid[i][j] == 'O' {
				return false
			}
		}
	}
	return true
}
