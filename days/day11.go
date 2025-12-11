package days

import (
	"fmt"
	"strings"

	"github.com/luc-10/Advent-of-code-2025/io"
)

func Day11() {
	Day11Part1()
	Day11Part2()
}

func Day11Part1() {
	adj := make(map[string][]string)
	data := io.ReadLines("inputFiles/day11.txt")
	for i := range data {
		line := strings.Split(data[i], ":")
		from := line[0]
		adj[from] = strings.Split(strings.Trim(line[1], " "), " ")
	}
	dp := make(map[string]int)
	for key := range adj {
		dp[key] = -1
	}
	fmt.Println(getPaths("you", "out", adj, dp))
}

func Day11Part2() {
	adj := make(map[string][]string)
	data := io.ReadLines("inputFiles/day11.txt")
	for i := range data {
		line := strings.Split(data[i], ":")
		from := line[0]
		adj[from] = strings.Split(strings.Trim(line[1], " "), " ")
	}
	dp := make(map[string]int)

	fmt.Println(getPathsThroughNodes("svr", "out", "fft", "dac", adj, dp) + getPathsThroughNodes("svr", "out", "dac", "fft", adj, dp))
}

func getPathsThroughNodes(from, to, node1, node2 string, adj map[string][]string, dp map[string]int) int {
	adjWithoutSecondNode := getAdjWithoutNode(adj, node2)
	paths := 1

	resetDp(adj, dp)
	paths *= getPaths(from, node1, adjWithoutSecondNode, dp)

	resetDp(adj, dp)
	paths *= getPaths(node1, node2, adj, dp)

	resetDp(adj, dp)
	paths *= getPaths(node2, to, adj, dp)

	return paths
}

func getPaths(from, to string, adj map[string][]string, dp map[string]int) int {
	if from == to {
		return 1
	}
	if dp[from] == -1 {
		sum := 0
		for _, next := range adj[from] {
			sum += getPaths(next, to, adj, dp)
		}
		dp[from] = sum
	}
	return dp[from]
}

func getAdjWithoutNode(adj map[string][]string, node string) map[string][]string {
	newAdj := make(map[string][]string)
	for key, value := range adj {
		newValue := make([]string, 0)
		for _, s := range value {
			if s != node {
				newValue = append(newValue, s)
			}
		}
		newAdj[key] = newValue
	}
	return newAdj
}

func resetDp(adj map[string][]string, dp map[string]int) {
	for key := range adj {
		dp[key] = -1
	}
}
