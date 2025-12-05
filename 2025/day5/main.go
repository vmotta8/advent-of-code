package main

import (
	"fmt"
	"os"
	"slices"
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
	sample := "3-5\n10-14\n16-20\n12-18\n30-40\n32-35\n50-55\n55-60\n\n1\n5\n8\n11\n17\n32\n33\n56"
	fmt.Println("sample:", sample[:1])

	ranges, nums := parseInput(input)
	fmt.Println("ranges:", len(ranges), "nums:", len(nums))

	result1 := part1(ranges, nums)
	result2 := part2(ranges)

	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}

func parseInput(input string) ([][]int64, []int64) {
	parts := strings.Split(input, "\n\n")
	rangeLines := strings.Split(parts[0], "\n")
	numLines := strings.Split(parts[1], "\n")

	rangesNums := make([][]int64, len(rangeLines))
	for i, r := range rangeLines {
		values := strings.Split(r, "-")
		first, _ := strconv.ParseInt(values[0], 10, 64)
		last, _ := strconv.ParseInt(values[1], 10, 64)
		rangesNums[i] = []int64{first, last}
	}

	slices.SortFunc(rangesNums, func(a, b []int64) int {
		return int(a[0] - b[0])
	})

	ranges := make([][]int64, 0)
	prev := rangesNums[0]
	for i, r := range rangesNums {
		if i == 0 {
			continue
		}
		if r[0] <= prev[1] {
			prev[1] = max(prev[1], r[1])
		} else {
			ranges = append(ranges, prev)
			prev = r
		}
	}
	ranges = append(ranges, prev)

	nums := make([]int64, len(numLines))
	for i, n := range numLines {
		nums[i], _ = strconv.ParseInt(n, 10, 64)
	}

	return ranges, nums
}

func part1(ranges [][]int64, nums []int64) int {
	result := 0
	for _, n := range nums {
		for _, r := range ranges {
			if n >= r[0] && n <= r[1] {
				result++
				break
			}
		}
	}

	return result
}

func part2(ranges [][]int64) int64 {
	result := int64(0)
	for _, r := range ranges {
		result += r[1] - r[0] + 1
	}

	return result
}
