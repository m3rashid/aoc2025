package main

import (
	"fmt"
	"strings"
)

func part1() {
	lines := strings.Split(data, "\n")
	numberOfLines := len(lines)
	lineWidth := len(lines[0])

	prevRays := make([]int, lineWidth)

	totalSplits := 0
	for i := range lineWidth {
		if lines[0][i] == 'S' {
			prevRays[i] = 1
		}
	}

	for i := 1; i < numberOfLines; i++ { // ignore the first line since we already processed it
		currentRays := make([]int, lineWidth)

		for j := range lineWidth {
			if prevRays[j] == 1 {
				if lines[i][j] == '^' {
					totalSplits++
					if j-1 >= 0 {
						currentRays[j-1] = 1
					}
					if j+1 < lineWidth {
						currentRays[j+1] = 1
					}
				} else {
					currentRays[j] = prevRays[j]
				}
			}
		}

		prevRays = currentRays
	}

	fmt.Printf("Part1 Total Rays: %d\n", totalSplits)
}

/*
1
1 1
1 1
1 2 1
1 2 1
1 3 3 1
1 3 3 1
1 4 3 3 1 1
1 4 3 3 1 1
1 5 4 3 4 2 1
1 5 4 3 4 2 1
1 1 5 4 7 4 2 1 1
1 1 5 4 7 4 2 1 1
1 2 10 11 11 2 1 1 1
*/

func part2() {
	lines := strings.Split(data, "\n")
	numberOfLines := len(lines)
	lineWidth := len(lines[0])

	allRays := map[int]int{} // index -> sum map

	for i := range lineWidth {
		if lines[0][i] == 'S' {
			allRays[i] = 1
		}
	}

	for i := 1; i < numberOfLines; i++ { // ignore the first line since we already processed it
		for j := range lineWidth {
			if val, ok := allRays[j]; ok && val > 0 {
				if lines[i][j] == '^' {
					delete(allRays, j) // remove the key

					allRays[j-1] += val
					allRays[j+1] += val
				}
			}
		}
	}

	total := 0
	for _, ray := range allRays {
		total += ray
	}
	fmt.Println("Part2 Result: ", total)
}

func main() {
	part1()
	part2()
}
