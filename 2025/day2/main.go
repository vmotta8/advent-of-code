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
	fmt.Println("input:", input)

	sample := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	fmt.Println("sample:", sample)

	nums := inputToInts(input)

	result1 := calculate(nums, isValid1)
	fmt.Println("Part 1:", result1)

	result2 := calculate(nums, isValid2)
	fmt.Println("Part 2:", result2)
}

func inputToInts(input string) [][]int64 {
	lines := strings.Split(input, ",")
	nums := make([][]int64, len(lines))
	for i, item := range lines {
		parts := strings.Split(item, "-")

		idx, _ := strconv.ParseInt(parts[0], 10, 64)
		idy, _ := strconv.ParseInt(parts[1], 10, 64)
		ids := []int64{idx, idy}
		nums[i] = ids
	}
	return nums
}

func calculate(nums [][]int64, isValid func(num int64) bool) int64 {
	sum := make([]int64, 0)
	for _, values := range nums {
		n1, n2 := values[0], values[1]

		for n := n1; n <= n2; n++ {
			valid := isValid(n)
			if !valid {
				sum = append(sum, n)
			}
		}
	}

	result := int64(0)
	for _, s := range sum {
		result += s
	}

	return result
}

func isValid1(num int64) bool {
	snum := strconv.FormatInt(num, 10)
	n := len(snum)
	if n%2 != 0 {
		return true
	}

	p1 := snum[:n/2]
	p2 := snum[n/2:]

	if p1 == p2 {
		return false
	}

	return true
}

func isValid2(num int64) bool {
	snum := strconv.FormatInt(num, 10)
	n := len(snum)

	for step := 1; step <= n/2; step++ {
		if !validate(snum, step) {
			return false
		}
	}

	return true
}

func validate(snum string, step int) bool {
	prev := snum[0:step]
	snum = snum[step:]

	for step <= len(snum) {
		curr := snum[0:step]
		snum = snum[step:]

		if prev != curr {
			return true
		}

		prev = curr
	}

	if len(snum) > 0 {
		return true
	}

	return false
}
