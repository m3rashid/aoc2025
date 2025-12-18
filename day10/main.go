package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Data struct {
	bulbPattern      map[int]bool // position -> bulb is on
	wiringSchemantic [][]int
	joltages         []int
}

func parseData(_str string) []Data {
	lines := strings.Split(_str, "\n")
	data := []Data{}

	for _, line := range lines {
		pos := 1 // 0 is [
		for line[pos] != ']' {
			pos++
		}

		bulbPatternStr := strings.TrimSpace(line[1:pos])
		bulbPattern := map[int]bool{}
		for i, ch := range bulbPatternStr {
			switch ch {
			case '#':
				bulbPattern[i] = true
			case '.':
				bulbPattern[i] = false
			}
		}

		pos += 2 // skip ] ans space
		initialPos := pos
		for line[pos] != '{' {
			pos++
		}
		buttonWiringSchemanticStr := strings.TrimSpace(line[initialPos:pos])
		schemas := strings.Split(buttonWiringSchemanticStr, " ")

		wiringSchemantic := [][]int{}
		for _, sc := range schemas {
			schema := sc[1 : len(sc)-1] // remove the ( and )

			positions := strings.Split(schema, ",")
			positionsInt := []int{}
			for _, pos := range positions {
				posInt, _ := strconv.Atoi(pos)
				positionsInt = append(positionsInt, posInt)
			}
			wiringSchemantic = append(wiringSchemantic, positionsInt)
		}

		pos += 1 // skip } and space
		initialPos = pos
		for line[pos] != '}' {
			pos++
		}
		joltageRequirementsStr := strings.TrimSpace(line[initialPos:pos])
		requirements := strings.Split(joltageRequirementsStr, ",")
		joltageRequirements := []int{}
		for _, req := range requirements {
			reqInt, _ := strconv.Atoi(req)
			joltageRequirements = append(joltageRequirements, reqInt)
		}

		data = append(data, Data{bulbPattern: bulbPattern, wiringSchemantic: wiringSchemantic, joltages: joltageRequirements})
	}

	return data
}

func analyzePart1(data Data) int {
	return 0
}

func part1() {
	inputData := parseData(sample)
	totalClicks := 0

	for _, data := range inputData {
		totalClicks += analyzePart1(data)
	}

	fmt.Println("Part 1 total: ", totalClicks)
}

func part2() {

}

func main() {
	part1()
	part2()
}
