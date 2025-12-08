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
	sample := "123 328  51 64\n 45 64  387 23\n  6 98  215 314\n*   +   *   +"
	fmt.Println("sample:", sample[:1])

	result1 := part1(input)
	result2 := part2(input)

	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}

func part1(input string) int64 {
	matrix, operators := inputToMatrix1(input)
	grid := transposeMatrix1(matrix)
	result := calculateColumns1(grid, operators)
	return result
}

func inputToMatrix1(input string) ([][]string, []string) {
	lines := strings.Split(input, "\n")
	numsLines := lines[:len(lines)-1]
	operatorsLine := lines[len(lines)-1]

	formattedNums := make([][]string, len(numsLines))
	nums := make([][]string, len(numsLines))
	for i, numsLine := range numsLines {
		nums[i] = strings.Split(numsLine, " ")
		for _, n := range nums[i] {
			if n != "" {
				formattedNums[i] = append(formattedNums[i], n)
			}
		}
	}

	formattedOperators := make([]string, 0)
	operators := strings.Split(operatorsLine, " ")
	for _, o := range operators {
		if o != "" {
			formattedOperators = append(formattedOperators, o)
		}
	}

	return formattedNums, formattedOperators
}

func transposeMatrix1(matrix [][]string) [][]int64 {
	grid := make([][]int64, len(matrix[0]))
	for row := range grid {
		grid[row] = make([]int64, len(matrix))
	}

	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			n, _ := strconv.ParseInt(matrix[i][j], 10, 64)
			grid[j][i] = n
		}
	}

	return grid
}

func calculateColumns1(grid [][]int64, operators []string) int64 {
	result := int64(0)
	for i, row := range grid {
		partial := int64(0)
		for _, n := range row {
			if operators[i] == "+" {
				partial += n
			}
			if operators[i] == "*" {
				if partial == 0 {
					partial = 1
				}
				partial *= n
			}
		}

		result += partial
	}

	return result
}

func part2(input string) int64 {
	matrix, operators := inputToMatrix2(input)
	result := calculateColumns2(matrix, operators)
	return result
}

func inputToMatrix2(input string) ([][]string, []string) {
	lines := strings.Split(input, "\n")
	numsLines := lines[:len(lines)-1]
	operatorsLine := lines[len(lines)-1]

	nums := make([][]string, len(numsLines))
	for i, numsLine := range numsLines {
		nums[i] = strings.Split(numsLine, "")
	}

	operators := strings.Split(operatorsLine, "")

	return nums, operators
}

func calculateColumns2(grid [][]string, operators []string) int64 {
	cols := 0
	for _, row := range grid {
		cols = max(len(row), cols)
	}

	operator := ""
	result := int64(0)
	partial := int64(0)
	for c := range cols {
		if len(operators) > c && operators[c] != " " {
			operator = operators[c]
			result += partial
			partial = int64(0)
		}

		sn := ""
		for r := range len(grid) {
			if len(grid[r]) <= c {
				continue
			}
			if grid[r][c] != " " {
				sn += grid[r][c]
			}
		}

		if operator == "+" && sn != "" {
			n, _ := strconv.ParseInt(sn, 10, 64)
			partial += n
		}
		if operator == "*" && sn != "" {
			n, _ := strconv.ParseInt(sn, 10, 64)
			if partial == 0 {
				partial = 1
			}
			partial *= n
		}
	}

	result += partial
	return result
}
