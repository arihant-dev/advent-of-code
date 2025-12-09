package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	out, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Normalize the line endings and split into lines
	content := strings.ReplaceAll(string(f), "\r\n", "\n")

	// create a 2-d array of integers
	rawLines := strings.Split(content, "\n")
	arrayOfPoints := make([][]int, 0)
	for _, line := range rawLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		a0, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		a1, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		arrayOfPoints = append(arrayOfPoints, []int{a0, a1})
	}
	part_1(arrayOfPoints, out)
	part_2(arrayOfPoints, out)
}

func part_1(points [][]int, out *os.File) {
	rectangleMaxArea := 0
	for _, p := range points {
		for _, v := range points {
			// Part 1 logic was just finding max area of any rectangle defined by two points
			// The prompt implies "opposite corners", so we just take any two points.
			area := abs(p[0]-v[0]+1) * abs(p[1]-v[1]+1)
			if area > rectangleMaxArea {
				rectangleMaxArea = area
			}
		}
	}
	fmt.Fprintf(out, "Part 1: Max Rectangle Area: %d\n", rectangleMaxArea)
}

func part_2(points [][]int, out *os.File) {
	n := len(points)
	maxArea := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := points[i]
			p2 := points[j]

			// Calculate potential area
			w := abs(p1[0]-p2[0]) + 1
			h := abs(p1[1]-p2[1]) + 1
			area := w * h

			if area <= maxArea {
				continue
			}

			// Check if valid
			if isValidRectangle(p1, p2, points) {
				maxArea = area
			}
		}
	}
	fmt.Fprintf(out, "Part 2: Max Valid Rectangle Area: %d\n", maxArea)
}

func isValidRectangle(p1, p2 []int, points [][]int) bool {
	xMin, xMax := min(p1[0], p2[0]), max(p1[0], p2[0])
	yMin, yMax := min(p1[1], p2[1]), max(p1[1], p2[1])

	// 1. Check if any polygon vertex is strictly inside the rectangle
	for _, p := range points {
		if p[0] > xMin && p[0] < xMax && p[1] > yMin && p[1] < yMax {
			return false
		}
	}

	// 2. Check if any polygon edge splits the rectangle
	n := len(points)
	for i := range n {
		u, v := points[i], points[(i+1)%n]

		// Vertical edge x=u[0]
		if u[0] == v[0] {
			if u[0] > xMin && u[0] < xMax {
				if min(u[1], v[1]) <= yMin && max(u[1], v[1]) >= yMax {
					return false
				}
			}
		}
		// Horizontal edge y=u[1]
		if u[1] == v[1] {
			if u[1] > yMin && u[1] < yMax {
				if min(u[0], v[0]) <= xMin && max(u[0], v[0]) >= xMax {
					return false
				}
			}
		}
	}

	// 3. Check if the center of the rectangle is inside the polygon
	// We check a point (xMin + 0.5, yMin + 0.5)
	checkX := float64(xMin) + 0.5
	checkY := float64(yMin) + 0.5
	intersections := 0
	for i := range n {
		u, v := points[i], points[(i+1)%n]
		// Check vertical edges for ray casting to the right
		if u[0] == v[0] {
			// Edge must be to the right of checkX
			if float64(u[0]) > checkX {
				y1, y2 := float64(u[1]), float64(v[1])
				if y1 > y2 {
					y1, y2 = y2, y1
				}
				// Check if checkY is within the y-range of the edge
				if checkY > y1 && checkY < y2 {
					intersections++
				}
			}
		}
	}

	return intersections%2 == 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
