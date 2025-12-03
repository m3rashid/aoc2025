package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getPart1Battery(line string) int {
	length := len(line)

	largestIdx := 0
	for i := 1; i < length; i++ {
		if line[i] > line[largestIdx] && i != length-1 {
			largestIdx = i
		}
	}

	secondLargestIdx := largestIdx + 1
	for i := largestIdx + 2; i < length; i++ {
		if line[i] > line[secondLargestIdx] {
			secondLargestIdx = i
		}
	}

	number, _ := strconv.Atoi(string(line[largestIdx]) + string(line[secondLargestIdx]))
	return number
}

func part1() {
	battery := 0
	lines := strings.SplitSeq(data, "\n")

	for line := range lines {
		battery += getPart1Battery(line)
	}

	fmt.Printf("Part1 Battery: %d\n", battery)
}

func getPart2Battery(line string) int64 {
	length := len(line)
	requiredLength := 12

	dropsLeft := length - requiredLength
	stack := []rune{}

	for _, digit := range line {
		for dropsLeft > 0 && len(stack) > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			dropsLeft--
		}
		stack = append(stack, digit)
	}

	number := string(stack[:requiredLength])
	num, _ := strconv.ParseInt(number, 10, 64)
	return num
}

func part2() {
	var battery int64 = 0
	lines := strings.SplitSeq(data, "\n")

	for line := range lines {
		battery += getPart2Battery(line)
	}

	fmt.Printf("Part2 Battery: %d\n", battery)
}

func main() {
	part1()
	part2()
}
