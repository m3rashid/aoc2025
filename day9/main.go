package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int64
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

// rectangleArea: inclusive grid-rectangle area between two opposite corners.
func rectangleArea(a, b Point) int64 {
	w := abs64(a.X-b.X) + 1
	h := abs64(a.Y-b.Y) + 1
	return w * h
}

// orientation / cross product: sign tells CW/CCW/collinear.
// >0: left turn, <0: right turn, 0: collinear
func orient(a, b, c Point) int64 {
	return (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
}

// check if point p lies on segment ab (inclusive)
func onSegment(p, a, b Point) bool {
	if orient(a, b, p) != 0 {
		return false
	}
	if p.X < min64(a.X, b.X) || p.X > max64(a.X, b.X) {
		return false
	}
	if p.Y < min64(a.Y, b.Y) || p.Y > max64(a.Y, b.Y) {
		return false
	}
	return true
}

// PointInPolygon: even-odd rule, returns true if inside or on boundary.
func pointInPolygon(p Point, poly []Point) bool {
	n := len(poly)
	if n < 3 {
		return false
	}

	// check boundary first
	for i := range n {
		j := (i + 1) % n
		if onSegment(p, poly[i], poly[j]) {
			return true
		}
	}

	inside := false
	px, py := p.X, p.Y
	for i := range n {
		j := (i + 1) % n
		xi, yi := poly[i].X, poly[i].Y
		xj, yj := poly[j].X, poly[j].Y

		// does horizontal ray at y=py intersect edge (yi,yj)?
		intersects := (yi > py) != (yj > py)
		if intersects {
			// compute x coordinate of intersection as float
			// cast to float64 only here; coordinates are small enough.
			xIntersect := float64(xi) + float64(py-yi)*float64(xj-xi)/float64(yj-yi)
			if float64(px) <= xIntersect {
				inside = !inside
			}
		}
	}

	return inside
}

// properSegmentIntersect returns true if segments ab and cd intersect
// in a "proper" crossing (interior point), not just touching or overlapping.
func properSegmentIntersect(a, b, c, d Point) bool {
	o1 := orient(a, b, c)
	o2 := orient(a, b, d)
	o3 := orient(c, d, a)
	o4 := orient(c, d, b)

	// if any collinear, we treat as non-proper (touching/overlap allowed)
	if o1 == 0 || o2 == 0 || o3 == 0 || o4 == 0 {
		return false
	}

	return (o1 > 0) != (o2 > 0) && (o3 > 0) != (o4 > 0)
}

// rectangleInsidePolygon checks if the entire inclusive rectangle
// [x1..x2] x [y1..y2] is contained in the polygon (boundary allowed).
// poly is the loop formed by the red tiles in order.
func rectangleInsidePolygon(a, b Point, poly []Point) bool {
	x1 := min64(a.X, b.X)
	x2 := max64(a.X, b.X)
	y1 := min64(a.Y, b.Y)
	y2 := max64(a.Y, b.Y)

	// corners
	corners := []Point{
		{x1, y1},
		{x2, y1},
		{x2, y2},
		{x1, y2},
	}

	// 1) all corners must be inside or on boundary
	for _, c := range corners {
		if !pointInPolygon(c, poly) {
			return false
		}
	}

	// 2) no rectangle edge may properly intersect any polygon edge
	rectEdges := [][2]Point{
		{{x1, y1}, {x2, y1}},
		{{x2, y1}, {x2, y2}},
		{{x2, y2}, {x1, y2}},
		{{x1, y2}, {x1, y1}},
	}

	n := len(poly)
	for _, e := range rectEdges {
		ra, rb := e[0], e[1]
		for i := range n {
			j := (i + 1) % n
			pa, pb := poly[i], poly[j]
			if properSegmentIntersect(ra, rb, pa, pb) {
				return false
			}
		}
	}

	return true
}

func parsePositions(_data string) []Point {
	lines := strings.Split(_data, "\n")
	positions := []Point{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.ParseInt(parts[1], 10, 64)
		y, _ := strconv.ParseInt(parts[0], 10, 64)
		positions = append(positions, Point{X: x, Y: y})
	}
	return positions
}

func main() {
	points := parsePositions(data)
	n := len(points)

	// Part 1: largest rectangle from any two red tiles (no restriction)
	var maxArea1 int64 = 0
	for i := range n {
		for j := i + 1; j < n; j++ {
			a := points[i]
			b := points[j]
			area := rectangleArea(a, b)
			if area > maxArea1 {
				maxArea1 = area
			}
		}
	}

	// Part 2: same, but rectangle must lie entirely in polygon whose
	// vertices are the red tiles in input order.
	// points already represent the loop (wrap-around).
	var maxArea2 int64 = 0
	for i := range n {
		for j := i + 1; j < n; j++ {
			a := points[i]
			b := points[j]

			area := rectangleArea(a, b)
			// quick pruning: cannot beat current best
			if area <= maxArea2 {
				continue
			}

			if rectangleInsidePolygon(a, b, points) {
				maxArea2 = area
			}
		}
	}

	fmt.Printf("Part 1: %d\n", maxArea1)
	fmt.Printf("Part 2: %d\n", maxArea2)
}
