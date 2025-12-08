package main

import (
	"fmt"
	"os"
	"sort"
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
	rawLines := strings.SplitSeq(content, "\n")

	// Process each line
	arrayOfPoints := make([]Point, 0)
	for line := range rawLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			panic("invalid point: " + line)
		}
		x := 0
		y := 0
		z := 0
		fmt.Sscanf(strings.TrimSpace(parts[0]), "%d", &x)
		fmt.Sscanf(strings.TrimSpace(parts[1]), "%d", &y)
		fmt.Sscanf(strings.TrimSpace(parts[2]), "%d", &z)
		arrayOfPoints = append(arrayOfPoints, Point{x: x, y: y, z: z})
	}
	part_1(arrayOfPoints, out)
	part_2(arrayOfPoints, out)

}

type Point struct {
	x int
	y int
	z int
}

type Edge struct {
	u, v   int
	distSq int
}

func part_1(points []Point, out *os.File) {
	n := len(points)
	var edges []Edge
	for i := range n {
		for j := i + 1; j < n; j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z
			distSq := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{u: i, v: j, distSq: distSq})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distSq < edges[j].distSq
	})

	limit := min(len(edges), 1000)

	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}

	union := func(i, j int) {
		rootI := find(i)
		rootJ := find(j)
		if rootI != rootJ {
			parent[rootI] = rootJ
		}
	}

	for i := range limit {
		union(edges[i].u, edges[i].v)
	}

	// Calculate component sizes
	sizes := make(map[int]int)
	for i := range n {
		root := find(i)
		sizes[root]++
	}

	var sizeList []int
	for _, s := range sizes {
		sizeList = append(sizeList, s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizeList)))

	ans := 1
	for i := 0; i < 3 && i < len(sizeList); i++ {
		ans *= sizeList[i]
	}

	fmt.Fprintf(out, "Part 1: Product of top 3 circuit sizes = %d\n", ans)
}
func part_2(points []Point, out *os.File) {
	n := len(points)
	if n < 2 {
		return
	}
	var edges []Edge
	for i := range n {
		for j := i + 1; j < n; j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z
			distSq := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{u: i, v: j, distSq: distSq})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distSq < edges[j].distSq
	})

	parent := make([]int, n)
	for i := range n {
		parent[i] = i
	}

	var find func(int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}

	union := func(i, j int) bool {
		rootI := find(i)
		rootJ := find(j)
		if rootI != rootJ {
			parent[rootI] = rootJ
			return true
		}
		return false
	}

	edgesCount := 0
	var lastU, lastV int

	for _, edge := range edges {
		if union(edge.u, edge.v) {
			edgesCount++
			lastU = edge.u
			lastV = edge.v
			if edgesCount == n-1 {
				break
			}
		}
	}

	ans := points[lastU].x * points[lastV].x
	fmt.Fprintf(out, "Part 2: Product of X coordinates of last connection = %d\n", ans)
}
