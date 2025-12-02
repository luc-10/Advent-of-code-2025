package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/luc-10/Advent-of-code-2025/io"
)

func Day2() {
	Day2Part1()
	Day2Part2()
}

func Day2Part1() {
	data := strings.Split(io.ReadLine("inputFiles/day02.txt"), ",")

	invalidSum := 0

	for _, s := range data {

		bounds := strings.Split(s, "-")
		bottom, err := strconv.Atoi(bounds[0])

		if err != nil {
			panic(err)
		}

		top, err := strconv.Atoi(bounds[1])

		if err != nil {
			panic(err)
		}

		for i := bottom; i <= top; i++ {
			left, right := splitInTwo(i)
			if left == right {
				invalidSum += i
			}
		}

	}

	fmt.Println(invalidSum)
}

func Day2Part2() {
	data := strings.Split(io.ReadLine("inputFiles/day02.txt"), ",")

	invalidSum := 0

	for _, s := range data {

		bounds := strings.Split(s, "-")
		bottom, err := strconv.Atoi(bounds[0])

		if err != nil {
			panic(err)
		}

		top, err := strconv.Atoi(bounds[1])

		if err != nil {
			panic(err)
		}

		for i := bottom; i <= top; i++ {
			len := length(i)
			for partLength := 1; partLength <= len/2; partLength++ {
				if len%partLength != 0 {
					continue
				}
				parts := splitInParts(i, partLength)
				if allEqual(parts) {
					invalidSum += i
					break
				}
			}
		}

	}

	fmt.Println(invalidSum)
}

func length(n int) int {
	len := 0
	for n > 0 {
		len++
		n /= 10
	}
	return len
}

func splitInTwo(num int) (int, int) {
	len := float64(length(num))

	right := num % int(math.Pow(10, math.Floor(len/2)))
	left := num / int(math.Pow(10, math.Floor(len/2)))

	return left, right
}

func splitInParts(num, partLength int) []int {
	len := length(num)
	ret := make([]int, 0)
	for i := partLength; i < len; i += partLength {
		ret = append(ret, num%int(math.Pow(10, float64(partLength))))
		num /= int(math.Pow(10, float64(partLength)))
	}
	ret = append(ret, num%int(math.Pow(10, float64(partLength))))
	return ret
}

func allEqual(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
}
