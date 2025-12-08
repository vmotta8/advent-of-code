package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Distances struct {
	A        string
	B        string
	Distance int64
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))
	fmt.Println("input:", input[:1])

	sample := "162,817,812\n57,618,57\n906,360,560\n592,479,940\n352,342,300\n466,668,158\n542,29,236\n431,825,988\n739,650,466\n52,470,668\n216,146,977\n819,987,18\n117,168,530\n805,96,715\n346,949,466\n970,615,88\n941,993,340\n862,61,35\n984,92,344\n425,690,689"
	fmt.Println("sample:", sample[:1])

	distances, points := parseInput(input)

	result1 := threeLargestCircuits(distances, points, 1000)
	fmt.Println("Part 1:", result1)

	result2 := lastTwoJunctionBoxesX(distances, points)
	fmt.Println("Part 2:", result2)
}

func threeLargestCircuits(distances []Distances, points []string, limit int) int {
	return threeLargestCircuitsBFS(distances[:limit], points)
}

func lastTwoJunctionBoxesX(distances []Distances, points []string) int {
	for limit := 1; limit <= len(distances); limit++ {
		result := threeLargestCircuitsBFS(distances[:limit], points)
		if result == len(points) {
			lastConn := distances[limit-1]
			ax, _, _ := parsePoint(lastConn.A)
			bx, _, _ := parsePoint(lastConn.B)
			return int(ax * bx)
		}
	}

	return 0
}

func parseInput(input string) ([]Distances, []string) {
	lines := strings.Split(input, "\n")
	dists := make([]Distances, 0)
	rows := len(lines)

	for i := range lines {
		for j := i + 1; j < rows; j++ {
			A, B := lines[i], lines[j]
			ax, ay, az := parsePoint(A)
			bx, by, bz := parsePoint(B)

			dx := bx - ax
			dy := by - ay
			dz := bz - az

			dists = append(dists, Distances{
				A:        A,
				B:        B,
				Distance: dx*dx + dy*dy + dz*dz,
			})
		}
	}

	slices.SortFunc(dists, func(a, b Distances) int {
		return int(a.Distance) - int(b.Distance)
	})

	return dists, lines
}

func parsePoint(s string) (int64, int64, int64) {
	parts := strings.Split(s, ",")
	x, _ := strconv.ParseInt(parts[0], 10, 64)
	y, _ := strconv.ParseInt(parts[1], 10, 64)
	z, _ := strconv.ParseInt(parts[2], 10, 64)

	return x, y, z
}

func threeLargestCircuitsBFS(distances []Distances, points []string) int {
	graph := make(map[string][]string)
	for _, dist := range distances {
		graph[dist.A] = append(graph[dist.A], dist.B)
		graph[dist.B] = append(graph[dist.B], dist.A)
	}

	result := make([]int, 0)
	visited := make(map[string]bool)
	queue := make([]string, 0)

	for _, p := range points {
		if visited[p] {
			continue
		}
		size := 0
		queue = append(queue, p)

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			if visited[curr] {
				continue
			}
			visited[curr] = true

			size++
			for _, neighbor := range graph[curr] {
				queue = append(queue, neighbor)
			}
		}

		result = append(result, size)
	}

	slices.Sort(result)
	slices.Reverse(result)

	for len(result) < 3 {
		result = append(result, 1)
	}

	return result[0] * result[1] * result[2]
}

func threeLargestCircuitsDSU(distances []Distances, points []string) int {
	parent := make(map[string]string)
	for _, p := range points {
		parent[p] = p
	}

	for _, dist := range distances {
		union(parent, dist.A, dist.B)
	}

	sizes := make(map[string]int)
	for _, p := range points {
		root := find(parent, p)
		sizes[root]++
	}

	result := make([]int, 0)
	for _, size := range sizes {
		result = append(result, size)
	}

	slices.Sort(result)
	slices.Reverse(result)

	for len(result) < 3 {
		result = append(result, 1)
	}

	return result[0] * result[1] * result[2]
}

func find(parent map[string]string, x string) string {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent map[string]string, a, b string) {
	rootA := find(parent, a)
	rootB := find(parent, b)
	if rootA != rootB {
		parent[rootA] = rootB
	}
}
