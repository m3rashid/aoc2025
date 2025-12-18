package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/google/btree"
)

type Box struct {
	x, y, z int
}

type Connection struct {
	boxA, boxB Box
	distance   float64
}

// ConnectionItem is the type stored in the B-tree, ordered by distance.
type ConnectionItem Connection

// Less implements btree.Item and sorts connections by ascending distance.
func (c ConnectionItem) Less(other btree.Item) bool {
	return c.distance < other.(ConnectionItem).distance
}

// boxDistance returns the 3D Euclidean distance between two boxes.
func boxDistance(a, b Box) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	dz := float64(a.z - b.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func prepareData(_data string) []Box {
	lines := strings.Split(_data, "\n")
	positions := []Box{}

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		positions = append(positions, Box{x: x, y: y, z: z})
	}

	return positions
}

// buildConnectionTree builds a B-tree containing all possible connections
// between pairs of boxes, sorted by straight-line distance.
//
// Only unique pairs (i < j) are considered, so each pair of boxes appears once.
func buildConnectionTree(boxes []Box) *btree.BTree {
	tree := btree.New(32)

	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			conn := ConnectionItem{
				boxA:     boxes[i],
				boxB:     boxes[j],
				distance: boxDistance(boxes[i], boxes[j]),
			}
			tree.ReplaceOrInsert(conn)
		}
	}

	return tree
}

// allBoxesInSingleCircuit reports whether all boxes belong to a single circuit,
// according to the circuit IDs stored in circuitByBox.
func allBoxesInSingleCircuit(circuitByBox map[Box]int, totalBoxes int) bool {
	// If some boxes haven't appeared in the map yet, they are still isolated.
	if len(circuitByBox) != totalBoxes {
		return false
	}

	first := true
	var ref int

	for _, circuitID := range circuitByBox {
		if first {
			ref = circuitID
			first = false
			continue
		}

		if circuitID != ref {
			return false
		}
	}

	return true
}

// simulateConnections walks the sorted list of possible connections
// (shortest first) and simulates plugging in strings of lights.
//
// - totalBoxes: total number of junction boxes in the playground
// - tree:       B-tree of potential connections, sorted by distance
// - maxConnections:
//   - if >= 0: stop after considering this many shortest connections
//   - if   -1: keep connecting until all boxes are in one circuit
//
// It returns:
// - a map assigning each box to a "circuit ID"
// - the last ConnectionItem that was processed
func simulateConnections(totalBoxes int, tree *btree.BTree, maxConnections int) (map[Box]int, ConnectionItem) {
	circuitByBox := make(map[Box]int) // box -> circuit ID
	nextCircuitID := 0
	var lastConnection ConnectionItem

	tree.Ascend(func(item btree.Item) bool {
		// Stop condition for part 2: all boxes are in a single circuit.
		if maxConnections == -1 && allBoxesInSingleCircuit(circuitByBox, totalBoxes) {
			return false
		}

		// Stop condition for part 1: we've considered maxConnections shortest pairs.
		if maxConnections != -1 && nextCircuitID == maxConnections {
			return false
		}

		conn := item.(ConnectionItem)
		lastConnection = conn
		nextCircuitID++

		a, b := conn.boxA, conn.boxB
		circuitA, hasA := circuitByBox[a]
		circuitB, hasB := circuitByBox[b]

		switch {
		case !hasA && !hasB:
			// Both boxes are new: start a brand new circuit.
			circuitByBox[a] = nextCircuitID
			circuitByBox[b] = nextCircuitID

		case hasA && !hasB:
			// Box B joins A's circuit.
			circuitByBox[b] = circuitA

		case !hasA && hasB:
			// Box A joins B's circuit.
			circuitByBox[a] = circuitB

		default:
			// Both boxes already belong to circuits:
			// merge the two circuits into a new circuit ID.
			for box, cid := range circuitByBox {
				if cid == circuitA || cid == circuitB {
					circuitByBox[box] = nextCircuitID
				}
			}
		}

		return true
	})

	return circuitByBox, lastConnection
}

// sortedCircuitSizes takes the circuit assignments and returns a slice of
// circuit sizes, sorted from largest to smallest.
func sortedCircuitSizes(circuitByBox map[Box]int) []int {
	// Count how many boxes belong to each circuit ID.
	boxesPerCircuit := make(map[int]int)
	for _, circuitID := range circuitByBox {
		boxesPerCircuit[circuitID]++
	}

	// Convert to slice of (circuitID, size) pairs for sorting.
	type circuitSize struct {
		id   int
		size int
	}

	pairs := make([]circuitSize, 0, len(boxesPerCircuit))
	for id, size := range boxesPerCircuit {
		pairs = append(pairs, circuitSize{id: id, size: size})
	}

	// Sort by size descending (largest circuits first).
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].size > pairs[j].size })

	// Extract just sizes in sorted order.
	sizes := make([]int, len(pairs))
	for i, p := range pairs {
		sizes[i] = p.size
	}

	return sizes
}

func main() {
	boxes := prepareData(data)
	tree := buildConnectionTree(boxes)

	circuits, _ := simulateConnections(len(boxes), tree, 1000)
	sizes := sortedCircuitSizes(circuits)
	fmt.Println("PART 1:", sizes[0]*sizes[1]*sizes[2])

	_, finalConnection := simulateConnections(len(boxes), tree, -1)
	fmt.Println("PART 2:", finalConnection.boxA.x*finalConnection.boxB.x)
}
