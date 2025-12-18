package main

import (
	"fmt"
	"strings"
)

func parseData(_data string) map[string][]string {
	lines := strings.Split(_data, "\n")

	data := map[string][]string{}
	for _, line := range lines {
		parts := strings.Split(line, ":")
		source := strings.TrimSpace(parts[0])
		others := strings.Split(strings.TrimSpace(parts[1]), " ")
		data[source] = others
	}

	return data
}

var sampleInput = parseData(sample2)

func countPart1(current string, val int, visited map[string]bool) int {
	fmt.Println("Current: ", current, "Value: ", val)

	visited[current] = true
	for _, other := range sampleInput[current] {
		if other == "out" {
			val += 1
		} else {
			val = countPart1(other, val, visited)
		}
	}
	return val
}

func part1() {
	visited := map[string]bool{}
	res := countPart1("you", 0, visited)
	fmt.Println("Part 1: ", res)
}

func part2() {
	stack := Stack[string]{}
	stack.Push("svr")

	paths := map[string][]string{}

	for !stack.IsEmpty() {
		current, _ := stack.Pop()
		if _, ok := paths[current]; !ok {
			paths[current] = []string{}
		}

	}
}

func main() {
	part1()
	part2()
}
