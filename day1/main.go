package main

import (
	"fmt"
	"strconv"
	"strings"
)

var temp = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func Part1() {
	lines := strings.Split(sample, "\n")
	point := 50
	zeroes := 0

	for _, line := range lines {
		dir := 1
		if line[0] == 'L' {
			dir = -1
		}

		num, _ := strconv.Atoi(line[1:])
		point = (point + (dir * num)) % 100
		if point == 0 {
			zeroes++
		}
	}

	fmt.Printf("\n\n\nPassword: %d\n", zeroes)
}

func Part2() {
	lines := strings.Split(sample, "\n")
	point := 50
	zeroes := 0

	for _, line := range lines {
		dir := 1
		if line[0] == 'L' {
			dir = -1
		}

		num, _ := strconv.Atoi(line[1:])
		zeroes += (num / 100)
		remainder := num % 100

		if remainder > 0 {
			new_pt := point + (dir * remainder)
			if (dir > 0 && new_pt >= 100) || (dir < 0 && new_pt <= 0 && point > 0) {
				zeroes++
			}
			point = (new_pt + 100) % 100
		}

		fmt.Printf("%s: %d %d\n", line, point, zeroes)
	}

	fmt.Printf("\n\n\nPassword: %d\n", zeroes)
}

func main() {
	Part1()
}
