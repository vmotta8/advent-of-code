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

	sdata := strings.TrimSpace(string(data))
	input := strings.Split(sdata, "\n")
	fmt.Println("input len:", len(input))

	sinput := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
	fmt.Println("sinput len:", len(sinput))

	result1 := int64(0)
	result2 := int64(0)
	for _, val := range input {
		result1 += part1(val)
		result2 += part2(val)
	}

	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}

func part1(input string) int64 {
	fidx := 0
	fval, _ := strconv.ParseInt(string(input[fidx]), 10, 64)

	for i := 0; i < len(input)-1; i++ {
		val := input[i]
		nval, _ := strconv.ParseInt(string(val), 10, 64)
		if nval > fval {
			fidx = i
			fval = nval
		}
	}

	sidx := fidx + 1
	sval, _ := strconv.ParseInt(string(input[sidx]), 10, 64)
	for i := sidx + 1; i < len(input); i++ {
		val := input[i]
		nval, _ := strconv.ParseInt(string(val), 10, 64)
		if nval > sval {
			sidx = i
			sval = nval
		}
	}

	num := string(input[fidx]) + string(input[sidx])
	result, _ := strconv.ParseInt(num, 10, 64)
	return result
}

func part2(input string) int64 {
	nremove := len(input) - 12
	fidx := 0
	fval, _ := strconv.ParseInt(string(input[fidx]), 10, 64)

	for i := range nremove + 1 {
		val := input[i]
		nval, _ := strconv.ParseInt(string(val), 10, 64)
		if nval > fval {
			fidx = i
			fval = nval
		}
	}

	input = input[fidx:]

	nremove -= fidx
	for nremove > 0 {
		finished := true
		lidx := 0
		lowest, _ := strconv.ParseInt(string(input[0]), 10, 64)
		curr, _ := strconv.ParseInt(string(input[0]), 10, 64)

		for i, val := range input {
			nval, _ := strconv.ParseInt(string(val), 10, 64)
			if nval > curr {
				input = input[:lidx] + input[lidx+1:]

				nremove--
				finished = false
				break
			}

			if nval < lowest {
				lidx = i
				lowest = nval
			}

			curr = nval
		}

		if finished {
			input = input[:len(input)-nremove]
			break
		}
	}

	result, _ := strconv.ParseInt(string(input), 10, 64)

	return result
}
