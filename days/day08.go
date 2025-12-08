package days

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/luc-10/Advent-of-code-2025/io"
)

type DistBoxes struct {
	dist float64
	box1 [3]int
	box2 [3]int
}

func Day8() {
	Day8Part1()
	Day8Part2()
}

func Day8Part1() {
	data := io.ReadLines("inputFiles/day08.txt")
	coords := getBoxesCoords(data)
	distBoxes := make([]DistBoxes, 0)

	coordIndex := make(map[[3]int]int)

	for i := range coords {
		coordIndex[coords[i]] = i
		for j := i + 1; j < len(coords); j++ {
			distBoxes = append(distBoxes, DistBoxes{getLineDistance(coords[i], coords[j]), coords[i], coords[j]})
		}
	}

	sort.Slice(distBoxes, func(i, j int) bool {
		return distBoxes[i].dist < distBoxes[j].dist
	})

	distBoxes = distBoxes[:1000]
	circuits := make([]int, len(coords))
	circuitIndex := 1

	for _, distBox := range distBoxes {
		circuitBox1, circuitBox2 := circuits[coordIndex[distBox.box1]], circuits[coordIndex[distBox.box2]]
		if circuitBox1 == 0 && circuitBox2 == 0 {
			circuits[coordIndex[distBox.box1]] = circuitIndex
			circuits[coordIndex[distBox.box2]] = circuitIndex
			circuitIndex++
		} else if circuitBox1 == 0 || circuitBox2 == 0 {
			circuits[coordIndex[distBox.box1]] = max(circuitBox1, circuitBox2)
			circuits[coordIndex[distBox.box2]] = max(circuitBox1, circuitBox2)
		} else if circuitBox1 == circuitBox2 {
			continue
		} else {
			replaceInArray(circuits, max(circuitBox1, circuitBox2), min(circuitBox1, circuitBox2))
		}
	}

	circuitSize := make([]int, circuitIndex)

	for _, circuit := range circuits {
		if circuit == 0 {
			continue
		}
		circuitSize[circuit]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(circuitSize)))
	fmt.Println(circuitSize[0] * circuitSize[1] * circuitSize[2])

}

func Day8Part2() {
	data := io.ReadLines("inputFiles/day08.txt")
	coords := getBoxesCoords(data)
	distBoxes := make([]DistBoxes, 0)

	coordIndex := make(map[[3]int]int)

	for i := range coords {
		coordIndex[coords[i]] = i
		for j := i + 1; j < len(coords); j++ {
			distBoxes = append(distBoxes, DistBoxes{getLineDistance(coords[i], coords[j]), coords[i], coords[j]})
		}
	}

	sort.Slice(distBoxes, func(i, j int) bool {
		return distBoxes[i].dist < distBoxes[j].dist
	})

	circuits := make([]int, len(coords))
	circuitIndex := 1

	for _, distBox := range distBoxes {
		circuitBox1, circuitBox2 := circuits[coordIndex[distBox.box1]], circuits[coordIndex[distBox.box2]]
		if circuitBox1 == 0 && circuitBox2 == 0 {
			circuits[coordIndex[distBox.box1]] = circuitIndex
			circuits[coordIndex[distBox.box2]] = circuitIndex
			circuitIndex++
		} else if circuitBox1 == 0 || circuitBox2 == 0 {
			circuits[coordIndex[distBox.box1]] = max(circuitBox1, circuitBox2)
			circuits[coordIndex[distBox.box2]] = max(circuitBox1, circuitBox2)
		} else if circuitBox1 == circuitBox2 {
			continue
		} else {
			replaceInArray(circuits, max(circuitBox1, circuitBox2), min(circuitBox1, circuitBox2))
		}
		if arraySame(circuits) {
			fmt.Println(distBox.box1[0] * distBox.box2[0])
			break
		}
	}

}

func getBoxesCoords(data []string) [][3]int {
	coords := make([][3]int, len(data))

	for i, line := range data {
		stringCoords := strings.Split(line, ",")
		for j := range 3 {
			coords[i][j], _ = strconv.Atoi(stringCoords[j])
		}
	}

	return coords
}

func getLineDistance(box1 [3]int, box2 [3]int) float64 {
	squareDist := 0.0
	for idx := range 3 {
		squareDist += math.Pow(float64(box2[idx]-box1[idx]), 2)
	}
	return math.Sqrt(squareDist)
}

func replaceInArray(arr []int, val1, val2 int) {
	for i := range arr {
		if arr[i] == val1 {
			arr[i] = val2
		}
	}
}

func arraySame(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			return false
		}
	}
	return true
}
