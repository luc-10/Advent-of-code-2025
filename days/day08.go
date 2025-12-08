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

func Day8() {
	Day8Part1()
	Day8Part2()
}

func Day8Part1() {
	data := io.ReadLines("inputFiles/day08.txt")
	coords := getBoxesCoords(data)
	junctions := 1000
	pq := dataStructures.NewPriorityQueue(func(a, b *dataStructures.PriorityItem[[2][3]int]) bool {
		return a.Priority > b.Priority
	})

	coordIndex := make(map[[3]int]int)
	for i := range coords {
		coordIndex[coords[i]] = i
		for j := i + 1; j < len(coords); j++ {
			distSquared := int(getLineDistanceSquared(coords[i], coords[j]))
			if pq.Len() < junctions {
				pq.Push([2][3]int{coords[i], coords[j]}, distSquared)
			} else {
				top, _ := pq.Top()
				if distSquared < int(getLineDistanceSquared(top[0], top[1])) {
					pq.Pop()
					pq.Push([2][3]int{coords[i], coords[j]}, distSquared)
				}
			}
		}
	}

	mfs := dataStructures.NewMfset(len(coords))
	for i := 0; i < junctions; i++ {
		boxes, _ := pq.Pop()
		mfs.Merge(coordIndex[boxes[0]], coordIndex[boxes[1]])
	}

	sizes := make([]int, len(coords))
	for i := range coords {
		root := mfs.Find(i)
		sizes[root]++
	}
	sort.Ints(sizes)

	val := 1
	for i := 1; i <= 3; i++ {
		val *= sizes[len(sizes)-i]
	}
	fmt.Println(val)
}

func Day8Part2() {
	data := io.ReadLines("inputFiles/day08.txt")
	coords := getBoxesCoords(data)
	pq := dataStructures.NewPriorityQueue(func(a, b *dataStructures.PriorityItem[[2][3]int]) bool {
		return a.Priority < b.Priority
	})

	coordIndex := make(map[[3]int]int)
	for i := range coords {
		coordIndex[coords[i]] = i
		for j := i + 1; j < len(coords); j++ {
			pq.Push([2][3]int{coords[i], coords[j]}, int(getLineDistanceSquared(coords[i], coords[j])))
		}
	}

	val := 0
	mfs := dataStructures.NewMfset(len(coords))
	for !pq.Empty() {
		boxes, _ := pq.Pop()
		mfs.Merge(coordIndex[boxes[0]], coordIndex[boxes[1]])
		if mfs.CountSets() == 1 {
			val = boxes[0][0] * boxes[1][0]
			break
		}
	}
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

func getLineDistanceSquared(box1 [3]int, box2 [3]int) float64 {
	squareDist := 0.0
	for idx := range 3 {
		squareDist += math.Pow(float64(box2[idx]-box1[idx]), 2)
	}
	return squareDist
}
