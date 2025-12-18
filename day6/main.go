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

type Operator struct {
	op       string
	maxDigit int
}

func getOperators(line string) []Operator {
	operators := []Operator{}

	lastOp := ""
	currentSize := 0

	for _, ch := range line {
		if ch == ' ' {
			currentSize++
			continue
		}

		if lastOp != "" {
			operators = append(operators, Operator{op: lastOp, maxDigit: currentSize})
		}

		lastOp = string(ch)
		currentSize = 0
	}

	if lastOp != "" {
		operators = append(operators, Operator{op: lastOp, maxDigit: currentSize + 1})
	}

	return operators
}

func part2() {
	lines := strings.Split(data, "\n")
	numberOfDataLines := len(lines) - 1
	operators := getOperators(lines[numberOfDataLines])

	currentXOffset := 0
	totalSum := 0

	for _, op := range operators {
		digits := op.maxDigit
		operator := op.op

		globalRes := 0
		if operator == "*" {
			globalRes = 1
		}

		for i := range digits {
			localRes := 0
			for j := range numberOfDataLines {
				digit := lines[j][currentXOffset+i]
				if digit != ' ' {
					localRes = localRes*10 + int(digit-'0')
				}
			}

			switch operator {
			case "*":
				globalRes *= localRes
			case "+":
				globalRes += localRes
			}
		}

		totalSum += globalRes

		currentXOffset += op.maxDigit + 1 // +1 for the space
	}

	fmt.Println("Part 2: ", totalSum)
}

func main() {
	part1()
	part2()
}
