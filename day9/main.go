package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func parsePositions(_data string) []Position {
	lines := strings.Split(_data, "\n")
	positions := []Position{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[1])
		y, _ := strconv.Atoi(parts[0])
		positions = append(positions, Position{x: x, y: y})
	}
	return positions
}

func getArea(p1 Position, p2 Position) int {
	return int(math.Abs(float64(p1.x-p2.x+1) * float64(p1.y-p2.y+1)))
}

func part1() {
	positions := parsePositions(data)

	size := len(positions)
	areaMatrix := make([][]int, size)
	for i := range size {
		areaMatrix[i] = make([]int, size)
	}

	for i := range size {
		currentRow := make([]int, size)

		for j := i + 1; j < size; j++ {
			area := getArea(positions[i], positions[j])
			currentRow[j] = area
		}

		areaMatrix[i] = currentRow
	}

	for i := range size {
		for j := range i {
			areaMatrix[i][j] = areaMatrix[j][i]
		}
	}

	maxArea := 0
	for i := range size {
		for j := range i {
			maxArea = max(maxArea, areaMatrix[i][j])
		}
	}
	fmt.Println("Part 1 Max Area:", maxArea)
}

func part2() {
	positions := parsePositions(data)
	_ = positions
}

func main() {
	part1()
	part2()
}
