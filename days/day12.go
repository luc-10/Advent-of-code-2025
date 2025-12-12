package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luc-10/Advent-of-code-2025/io"
)

func Day12() {
	Day12Part1()
}

func Day12Part1() {
	data := io.ReadLine("inputFiles/day12.txt")
	shapeAreas, gridInfos := parseInput(data)
	count := 0
	for i := range gridInfos {
		gridArea := gridInfos[i][0] * gridInfos[i][1]
		for j := 2; j < len(gridInfos[i]); j++ {
			gridArea -= shapeAreas[j-2] * gridInfos[i][j]
		}
		if gridArea >= 0 {
			count++
		}
	}
	fmt.Println(count)

}

func parseInput(data string) ([]int, [][]int) {
	splitData := strings.Split(data, "\n\n")
	shapes := splitData[:len(splitData)-1]
	grids := strings.Split(splitData[len(splitData)-1], "\n")

	shapeAreas := make([]int, len(shapes))
	for i, sh := range shapes {
		area := 0
		for _, char := range sh {
			if char == '#' {
				area++
			}
		}
		shapeAreas[i] = area
	}
	gridInfos := make([][]int, len(grids))
	for i := range grids {
		grid := strings.Split(grids[i], " ")
		gridSizes := strings.Split(strings.Trim(grid[0], ":"), "x")
		gridX, _ := strconv.Atoi(gridSizes[0])
		gridY, _ := strconv.Atoi(gridSizes[1])
		gridInfos[i] = append(gridInfos[i], gridX, gridY)
		for j := 1; j < len(grid); j++ {
			occurrence, _ := strconv.Atoi(grid[j])
			gridInfos[i] = append(gridInfos[i], occurrence)
		}

	}

	return shapeAreas, gridInfos
}
