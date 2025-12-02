package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isInValidPart1(str string) bool {
	length := len(str)
	mid := length / 2
	return str[:mid] == str[mid:]
}

func part1() {
	var total int64 = 0
	ranges := strings.SplitSeq(data, ",")

	for numRange := range ranges {
		parts := strings.Split(numRange, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			if isInValidPart1(str) {
				total += int64(i)
			}
		}
	}

	fmt.Printf("Part1 Total: %d\n", total)
}

func isInValidPart2(str string) bool {
	length := len(str)
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for _, num := range nums {
		if length%num == 0 { // string might have a repeated pattern (pattern length = num, pattern = str[:length/num])
			pattern := str[:length/num]
			repeated := strings.Repeat(pattern, num)
			if repeated == str {
				return true
			}
		}
	}
	return false
}

func part2() {
	var total int64 = 0
	ranges := strings.SplitSeq(data, ",")

	for numRange := range ranges {
		parts := strings.Split(numRange, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			if isInValidPart2(str) {
				total += int64(i)
			}
		}
	}

	fmt.Printf("Part2 Total: %d\n", total)
}

func main() {
	part1()
	part2()
}
