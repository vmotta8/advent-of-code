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

	result1 := solveByWords(input)
	fmt.Println("Part 1:", result1)

	result2 := solveByChars(input)
	fmt.Println("Part 2:", result2)
}

func solveByWords(input string) int64 {
	grid, operators := parseByWords(input)
	return calculateByWords(grid, operators)
}

func parseByWords(input string) ([][]string, []string) {
	lines := strings.Split(input, "\n")
	numsLines := lines[:len(lines)-1]
	operatorsLine := lines[len(lines)-1]

	nums := make([][]string, len(numsLines))
	for i, numsLine := range numsLines {
		for _, n := range strings.Split(numsLine, " ") {
			if n != "" {
				nums[i] = append(nums[i], n)
			}
		}
	}

	operators := make([]string, 0)
	for _, o := range strings.Split(operatorsLine, " ") {
		if o != "" {
			operators = append(operators, o)
		}
	}

	return nums, operators
}

func calculateByWords(grid [][]string, operators []string) int64 {
	result := int64(0)
	cols := len(grid[0])

	for c := range cols {
		partial := int64(0)
		for r := range len(grid) {
			n, _ := strconv.ParseInt(grid[r][c], 10, 64)
			if operators[c] == "+" {
				partial += n
			}
			if operators[c] == "*" {
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

func solveByChars(input string) int64 {
	grid, operators := parseByChars(input)
	return calculateByChars(grid, operators)
}

func parseByChars(input string) ([][]string, []string) {
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

func calculateByChars(grid [][]string, operators []string) int64 {
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
