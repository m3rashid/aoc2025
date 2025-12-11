package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1() {
	lines := strings.Split(data, "\n")
	numberOfLines := len(lines)

	operators := []string{}
	operatorLine := strings.SplitSeq(lines[numberOfLines-1], " ")
	for op := range operatorLine {
		operator := strings.TrimSpace(op)
		if operator != "" {
			operators = append(operators, operator)
		}
	}

	numbers := [][]int64{}
	for i := range numberOfLines - 1 {
		line := strings.SplitSeq(lines[i], " ")
		localNos := []int64{}
		for nm := range line {
			number := strings.TrimSpace(nm)
			if number != "" {
				num, _ := strconv.ParseInt(number, 10, 64)
				localNos = append(localNos, num)
			}
		}
		numbers = append(numbers, localNos)
	}

	var res int64 = 0
	n := len(operators)
	for i := range n {
		op := operators[i]
		var localRes int64
		if op == "*" {
			localRes = 1
		} else {
			localRes = 0
		}

		for j := range numberOfLines - 1 {
			num := numbers[j][i]
			switch op {
			case "*":
				localRes *= num
			case "+":
				localRes += num
			}
		}

		res += localRes
	}

	fmt.Println("Part 1: ", res)
}

type operator struct {
	val      string
	digitLen int
	numbers  []string
}

func part2() {
	lines := strings.Split(sample, "\n")
	numberOfDataLines := len(lines) - 1

	operators := []operator{}
	operatorLine := []rune(lines[numberOfDataLines])

	digitLen := 1
	for i := len(operatorLine) - 1; i >= 0; i-- {
		if operatorLine[i] == ' ' {
			digitLen++
		} else {
			operators = append(operators, operator{val: string(operatorLine[i]), digitLen: digitLen})
			digitLen = 0
		}
	}

	// reverse the operators
	for i, j := 0, len(operators)-1; i < j; i, j = i+1, j-1 {
		operators[i], operators[j] = operators[j], operators[i]
	}

	allNumbers := [][]string{}

	for i := range numberOfDataLines {
		line := lines[i]
		number := []string{}

		currentNo := ""

		for j, ch := range line {
			if ch != ' ' {
				currentNo += string(ch)

				continue
			}
		}
	}
}

func main() {
	// part1()
	part2()
}
