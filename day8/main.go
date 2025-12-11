package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
	z int
}

func prepareData(_data string) []Position {
	lines := strings.Split(_data, "\n")
	positions := []Position{}

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		positions = append(positions, Position{x: x, y: y, z: z})
	}

	return positions
}

func getDistance(p1 Position, p2 Position) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2) + math.Pow(float64(p1.z-p2.z), 2))
}

func isConnected(connections [][]int, el1Idx int, el2Idx int) bool {
	for _, connection := range connections {
		if slices.Contains(connection, el1Idx) && slices.Contains(connection, el2Idx) {
			return true
		}
	}
	return false
}

func prepareDistanceMatrix(positions []Position) [][]float64 {
	matrix := make([][]float64, len(positions))
	for i := range positions {
		matrix[i] = make([]float64, len(positions))
	}

	for i := range positions {
		for j := range positions {
			distance := getDistance(positions[i], positions[j])
			matrix[i][j] = distance
			matrix[j][i] = distance
		}
	}

	return matrix
}

func part1() {
	maxConnections := 1000
	positions := prepareData(data)
	sampleSize := len(positions)

	distanceMatrix := prepareDistanceMatrix(positions)

	connections := [][]int{}

	for range maxConnections {
		minDist := math.MaxFloat64
		el1Idx := 0
		el2Idx := 0

		for i := range positions {
			for j := i + 1; j < sampleSize; j++ {
				if distanceMatrix[i][j] < minDist && !isConnected(connections, i, j) {
					minDist = distanceMatrix[i][j]
					el1Idx = i
					el2Idx = j
				}
			}
		}

		added := false
		for i := range connections {
			containsEl1 := slices.Contains(connections[i], el1Idx)
			containsEl2 := slices.Contains(connections[i], el2Idx)

			if containsEl1 && containsEl2 {
				continue
			}

			if containsEl1 {
				connections[i] = append(connections[i], el2Idx)
				added = true
				break
			} else if containsEl2 {
				connections[i] = append(connections[i], el1Idx)
				added = true
				break
			}
		}

		if !added {
			connections = append(connections, []int{el1Idx, el2Idx})
		}
	}

	circuitSizes := make([]int, len(connections))
	for i := range connections {
		circuitSizes[i] = len(connections[i])
	}

	slices.Sort(circuitSizes)
	length := len(circuitSizes)

	sum := 1
	for i := length - 1; i >= length-3; i-- {
		sum *= circuitSizes[i]
	}

	fmt.Printf("Part1 Sum: %d\n", sum)
}

func part2() {}

func main() {
	part1()
	part2()
}
