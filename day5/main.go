package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Database struct {
	idRanges [][2]int64
	ids      []int64
}

func prepareData(_data string) Database {
	separatedData := strings.Split(_data, "\n\n")

	idRangesStr := strings.Split(separatedData[0], "\n")
	idRanges := make([][2]int64, len(idRangesStr))

	for i, idRange := range idRangesStr {
		parts := strings.Split(idRange, "-")
		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end, _ := strconv.ParseInt(parts[1], 10, 64)
		idRanges[i] = [2]int64{start, end}
	}

	sort.Slice(idRanges, func(i, j int) bool {
		if idRanges[i][0] == idRanges[j][0] { // if the start is the same, sort by end
			return idRanges[i][1] < idRanges[j][1]
		}
		// otherwise, sort by start
		return idRanges[i][0] < idRanges[j][0]

	})

	idsStr := strings.Split(separatedData[1], "\n")
	ids := make([]int64, len(idsStr))
	for i, id := range idsStr {
		id, _ := strconv.ParseInt(id, 10, 64)
		ids[i] = id
	}
	slices.Sort(ids)

	return Database{idRanges: idRanges, ids: ids}
}

func part1() {
	database := prepareData(data)
	count := 0

	for _, id := range database.ids {
		for _, rng := range database.idRanges {
			if id >= rng[0] && id <= rng[1] {
				count++
				break
			}
		}
	}

	fmt.Printf("Part1 Count: %d\n", count)
}

// 334291616815636 is too low
func part2() {
	database := prepareData(data)
	var count int64 = 0

	var prevEnd int64 = 0
	for _, rng := range database.idRanges {
		start := max(prevEnd, rng[0])
		end := rng[1]

		if end >= start {
			count += end - start + 1
			prevEnd = end + 1
		}
	}

	fmt.Printf("Part2 Count: %d\n", count)
}

func main() {
	part1()
	part2()
}
