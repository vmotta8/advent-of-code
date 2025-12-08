package main

import (
	"fmt"
	"os"
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
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}
	fmt.Println("sample len:", len(sample))

	limit := 3
	result1 := countIsolated(input, limit)
	fmt.Println("Part 1:", result1)

	result2 := countRemovedIteratively(input, limit)
	fmt.Println("Part 2:", result2)
}

func inputToGrid(input []string) [][]string {
	result := make([][]string, len(input))
	for i, val := range input {
		arr := strings.Split(val, "")
		result[i] = arr
	}

	return result
}

func countIsolated(input []string, limit int) int {
	grid := inputToGrid(input)
	result := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == "@" {
				if hasFewerAdjacents(grid, row, col, limit) {
					result++
				}
			}
		}
	}

	return result
}

func countRemovedIteratively(input []string, limit int) int {
	grid := inputToGrid(input)
	result := 0

	var exec func () bool
	exec = func () bool {
		updated := false
		for row := range grid {
			for col := range grid[row] {
				if grid[row][col] == "@" {
					if hasFewerAdjacents(grid, row, col, limit) {
						grid[row][col] = "."
						updated = true
						result++
					}
				}
			}
		}

		return updated
	}

	updated := exec()
	for updated {
		updated = exec()
	}

	return result
}

func hasFewerAdjacents(grid [][]string, x, y, limit int) bool {
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
	counter := 0
	rows := len(grid)
	cols := len(grid[0])

	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		nx, ny := x+dx, y+dy

		if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
			continue
		}

		if grid[nx][ny] == "@" {
			counter++
		}

		if counter > limit {
			return false
		}
	}

	return true
}
