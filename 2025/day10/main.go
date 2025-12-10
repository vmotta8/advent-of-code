package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	Goal     int
	Lights   int
	Buttons  []int
	Joltages []int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))
	fmt.Println("input:", input[:1])

	sample := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}\n[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}\n[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}"
	fmt.Println("sample:", sample[:1])

	machines := parseInput(input)

	result1 := 0
	for _, m := range machines {
		result1 += minButtonPresses(m)
	}
	fmt.Println("Part 1:", result1)

	result2 := 0
	fmt.Println("Part 2:", result2)
}

func parseInput(input string) []Machine {
	lines := strings.Split(input, "\n")

	result := make([]Machine, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")

		goal := 0
		lights := 0
		buttons := make([]int, 0)
		joltages := make([]int, 0)

		for _, part := range parts {
			inner := part[1 : len(part)-1]

			switch part[0] {
			case '[':
				lights = len(inner)
				for j, c := range inner {
					if c == '#' {
						goal |= (1 << j)
					}
				}
			case '(':
				button := 0
				for _, ns := range strings.Split(inner, ",") {
					num, _ := strconv.Atoi(ns)
					button |= (1 << num)
				}
				buttons = append(buttons, button)
			case '{':
				for _, ns := range strings.Split(inner, ",") {
					num, _ := strconv.Atoi(ns)
					joltages = append(joltages, num)
				}
			}
		}

		result[i] = Machine{
			Goal:     goal,
			Lights:   lights,
			Buttons:  buttons,
			Joltages: joltages,
		}
	}

	return result
}

func minButtonPresses(m Machine) int {
	numButtons := len(m.Buttons)
	minPresses := math.MaxInt

	for mask := 0; mask < (1 << numButtons); mask++ {
		lights := 0
		presses := 0

		for i := range numButtons {
			if mask&(1<<i) != 0 {
				lights ^= m.Buttons[i]
				presses++
			}
		}

		if lights == m.Goal && presses < minPresses {
			minPresses = presses
		}
	}

	return minPresses
}
