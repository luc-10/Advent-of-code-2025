package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/luc-10/Advent-of-code-2025/io"
)

func Day5() {
	Day5Part1()
	Day5Part2()
}

func Day5Part1() {
	data := io.ReadLines("inputFiles/day05.txt")
	ranges, ids := getRangesAndIDs(data)

	rangeIndex, idIndex, freshIngredients := 0, 0, 0
	for idIndex < len(ids) && rangeIndex < len(ranges) {
		if ids[idIndex] >= ranges[rangeIndex][0] && ids[idIndex] <= ranges[rangeIndex][1] {
			freshIngredients++
			idIndex++
		} else if ids[idIndex] < ranges[rangeIndex][0] {
			idIndex++
		} else if ids[idIndex] > ranges[rangeIndex][1] {
			rangeIndex++
		}
	}
	fmt.Println(freshIngredients)
}

func Day5Part2() {
	data := io.ReadLines("inputFiles/day05.txt")
	ranges, _ := getRangesAndIDs(data)

	freshIngredients, prevTop := 0, 0
	for _, pair := range ranges {
		if prevTop >= pair[1] {
			continue
		} else if prevTop == pair[0] || (prevTop > pair[0] && prevTop < pair[1]) {
			pair[0] = prevTop + 1
		}

		freshIngredients += (pair[1] - pair[0] + 1)
		prevTop = pair[1]
	}

	fmt.Println(freshIngredients)
}

func getRangesAndIDs(data []string) ([][2]int, []int) {
	i := 0
	for ; i < len(data); i++ {
		if data[i] == "" {
			break
		}
	}
	ranges := make([][2]int, 0)
	for _, s := range data[:i] {
		rangesString := strings.Split(s, "-")
		bottom, _ := strconv.Atoi(rangesString[0])
		top, _ := strconv.Atoi(rangesString[1])
		ranges = append(ranges, [2]int{bottom, top})
	}

	ids := make([]int, 0)
	for _, s := range data[i+1:] {
		id, _ := strconv.Atoi(s)
		ids = append(ids, id)
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}
		return ranges[i][0] < ranges[j][0]
	})
	sort.Ints(ids[:])

	return ranges, ids
}
