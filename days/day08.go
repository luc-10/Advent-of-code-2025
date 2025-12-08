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

	mfset := dataStructures.NewMfset(len(coords))
	for _, distBox := range distBoxes {
		mfset.Merge(coordIndex[distBox.box1], coordIndex[distBox.box2])
	}

	circuits := mfset.GetSetSizes()
	sort.Ints(circuits)
	val := 1
	for i := 1; i <= 3; i++ {
		val *= circuits[len(circuits)-i]
	}
	fmt.Println(val)
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
	mfset := dataStructures.NewMfset(len(coords))
	var lastBoxes DistBoxes
	for _, distBox := range distBoxes {
		lastBoxes = distBox
		mfset.Merge(coordIndex[distBox.box1], coordIndex[distBox.box2])
		if mfset.CountSets() == 1 {
			break
		}
	}
	val := lastBoxes.box1[0] * lastBoxes.box2[0]
	fmt.Println(val)

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
