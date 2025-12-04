package main

import (
	"fmt"
	"strings"
)

func getCount(_matrix [][]rune, toReplace bool) ([][]rune, int) {
	count := 0
	_data := _matrix

	numberOfLines := len(_data)
	lineLength := len(_data[0])

	for i, line := range _data {
		for j, ch := range line {
			if ch != '@' {
				continue
			}

			// check all the 8 adjacent cells of the '@'
			countOfAt := 0
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if k == 0 && l == 0 || (i+k < 0 || i+k >= numberOfLines || j+l < 0 || j+l >= lineLength) || (_data[i+k][j+l] != '@') {
						continue
					}

					countOfAt++
				}
			}

			if countOfAt < 4 {
				if toReplace {
					_data[i][j] = '.'
				}
				count++
			}
		}
	}

	return _data, count
}

func prepareData(_data string) [][]rune {
	lines := strings.Split(_data, "\n")

	numberOfLines := len(lines)
	lineLength := len(lines[0])
	currentData := make([][]rune, numberOfLines)

	for i, line := range lines {
		currentData[i] = make([]rune, lineLength)
		for j, ch := range line {
			currentData[i][j] = ch
		}
	}

	return currentData
}

func part1() {
	currentData := prepareData(data)
	_, count := getCount(currentData, false)

	fmt.Printf("Part1 Count: %d\n", count)
}

func part2() {
	count := 0
	prevCount := -1

	currentData := prepareData(data)
	for prevCount != count {
		prevCount = count
		_currentData, _count := getCount(currentData, true)

		currentData = _currentData
		count += _count
	}

	fmt.Printf("Part2 Count: %d\n", count)
}

func main() {
	part1()
	part2()
}
