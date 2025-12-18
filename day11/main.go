package main

import (
	"fmt"
	"strings"
)

type Graph map[string][]string

func parseData(_data string) Graph {
	lines := strings.Split(_data, "\n")

	data := Graph{}
	for _, line := range lines {
		parts := strings.Split(line, ":")
		source := strings.TrimSpace(parts[0])
		others := strings.Split(strings.TrimSpace(parts[1]), " ")
		data[source] = others
	}

	return data
}

func countPathsDFS(graph Graph, src, dst string, memo map[string]int) int {
	if src == dst {
		return 1
	}

	if v, ok := memo[src]; ok {
		return v
	}

	total := 0
	for _, next := range graph[src] {
		total += countPathsDFS(graph, next, dst, memo)
	}

	memo[src] = total
	return total
}

func hasPath(graph Graph, src, dst string, visited map[string]bool) bool {
	if src == dst {
		return true
	}

	if visited[src] {
		return false
	}
	visited[src] = true

	for _, next := range graph[src] {
		if hasPath(graph, next, dst, visited) {
			return true
		}
	}

	return false
}

func countPaths(graph Graph, src, dst string) int {
	return countPathsDFS(graph, src, dst, make(map[string]int))
}

func countPathsVia(graph Graph, src, dst string, viaA, viaB string) int {
	// Determine order
	if hasPath(graph, viaA, viaB, make(map[string]bool)) {
		return pathsThrough(graph, src, viaA, viaB, dst)
	}
	return pathsThrough(graph, src, viaB, viaA, dst)
}

func pathsThrough(graph Graph, src, mid1, mid2, dst string) int {
	p1 := countPathsDFS(graph, src, mid1, make(map[string]int))
	p2 := countPathsDFS(graph, mid1, mid2, make(map[string]int))
	p3 := countPathsDFS(graph, mid2, dst, make(map[string]int))
	return p1 * p2 * p3
}

func main() {
	graph := parseData(data)

	part1 := countPaths(graph, "you", "out")
	fmt.Printf("Part 1: %d\n", part1)

	part2 := countPathsVia(graph, "svr", "out", "dac", "fft")
	fmt.Printf("Part 2: %d\n", part2)
}
