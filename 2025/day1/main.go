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

	input := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println("input len:", len(input))

	sample := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
	fmt.Println("sample len:", len(sample))

	nums := inputToInts(input)
	curr := 50
	limit := 99

	result1 := part1(nums, curr, limit)
	fmt.Println("Part 1:", result1)

	result2 := part2(nums, curr, limit)
	fmt.Println("Part 2:", result2)
}

func inputToInts(input []string) []int {
	result := make([]int, len(input))

	for i, val := range input {
		rval := []rune(val)
		rnums := rval[1:]
		num, _ := strconv.Atoi(string(rnums))

		rside := rval[0]
		switch rside {
		case 'L':
			result[i] = -num
		case 'R':
			result[i] = num
		}
	}

	return result
}

func part1(nums []int, curr, limit int) int {
	result := 0
	cycle := limit + 1
	for _, num := range nums {
		moves := num % cycle
		curr += moves
		if curr < 0 {
			curr += cycle
		}
		if curr > limit {
			curr -= cycle
		}
		if curr == 0 {
			result++
		}
	}

	return result
}

func part2(nums []int, curr, limit int) int {
	result := 0
	cycle := limit + 1
	for _, num := range nums {
		cycles := abs(num / cycle)
		result += cycles

		moves := num % cycle
		if moves < 0 && curr != 0 {
			if curr+moves < 0 {
				result++
			}
		}
		if moves > 0 && curr != 0 {
			if curr+moves > cycle {
				result++
			}
		}

		curr += moves
		if curr < 0 {
			curr += limit + 1
		}
		if curr > limit {
			curr -= limit + 1
		}
		if curr == 0 {
			result++
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
