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

	input := strings.TrimSpace(string(data))
	fmt.Println("input:", input[:1])
	sample := ".......S.......\n...............\n.......^.......\n...............\n......^.^......\n...............\n.....^.^.^.....\n...............\n....^.^...^....\n...............\n...^.^...^.^...\n...............\n..^...^.....^..\n...............\n.^.^.^.^.^...^.\n..............."
	fmt.Println("sample:", sample[:1])

	grid := parseInput(sample)

	result1 := countHits(grid)
	fmt.Println("Part 1:", result1)

	result2 := countWays(grid)
	fmt.Println("Part 2:", result2)
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")

	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	return grid
}

func countHits(grid [][]string) int {
	start := 0
	for col, cell := range grid[0] {
		if cell == "S" {
			start = col
			break
		}
	}

	beams := map[int]bool{start: true}
	splits := 0

	for row := 1; row < len(grid); row++ {
		next := map[int]bool{}
		for col := range beams {
			if grid[row][col] == "^" {
				splits++
				next[col-1] = true
				next[col+1] = true
			} else {
				next[col] = true
			}
		}
		beams = next
	}

	return splits
}

func countWays(grid [][]string) int64 {
	start := 0
	for col, cell := range grid[0] {
		if cell == "S" {
			start = col
			break
		}
	}

	counts := make([]int64, len(grid[0]))
	counts[start] = 1

	for row := 1; row < len(grid); row++ {
		next := make([]int64, len(grid[0]))
		for col, count := range counts {
			if count == 0 {
				continue
			}
			if grid[row][col] == "^" {
				next[col-1] += count
				next[col+1] += count
			} else {
				next[col] += count
			}
		}
		counts = next
	}

	var total int64
	for _, c := range counts {
		total += c
	}
	return total
}
