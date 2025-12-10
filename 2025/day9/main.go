package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))
	fmt.Println("input:", input[:1])

	sample := "7,1\n11,1\n11,7\n9,7\n9,5\n2,5\n2,3\n7,3"
	fmt.Println("sample:", sample[:1])

	points := parseInput(input)

	result1 := largestRectangleArea(points)
	fmt.Println("Part 1:", result1)

	result2 := 0
	fmt.Println("Part 2:", result2)
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")

	result := make([][]int, len(lines))
	for i, line := range lines {
		spoint := strings.Split(line, ",")
		point := make([]int, 2)
		for j, p := range spoint {
			n, _ := strconv.Atoi(p)
			point[j] = n
		}
		result[i] = point
	}

	return result
}

func largestRectangleArea(points [][]int) int {
	result := 0
	n := len(points)
	for i := 0; i < n; i++ {
		p1 := points[i]
		for j := i + 1; j < n; j++ {
			p2 := points[j]

			dx := abs(p1[0]-p2[0]) + 1
			dy := abs(p1[1]-p2[1]) + 1
			result = max(result, dx*dy)
		}
	}

	return result
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
