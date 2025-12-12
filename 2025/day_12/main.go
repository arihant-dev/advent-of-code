package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Shape struct {
	ID     int
	Pixels [][]bool
	Size   int
}

type Region struct {
	Width, Height int
	Reqs          []int
}

func main() {
	content, err := os.ReadFile("input_final.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	var shapes []Shape
	var regions []Region

	var currentShapePixels [][]bool

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasSuffix(line, ":") && !strings.Contains(line, "x") {
			if len(currentShapePixels) > 0 {
				shapes = append(shapes, createShape(len(shapes), currentShapePixels))
				currentShapePixels = nil
			}
			continue
		}

		if strings.Contains(line, "x") {
			if len(currentShapePixels) > 0 {
				shapes = append(shapes, createShape(len(shapes), currentShapePixels))
				currentShapePixels = nil
			}

			parts := strings.Split(line, ":")
			dims := strings.Split(strings.TrimSpace(parts[0]), "x")
			w, _ := strconv.Atoi(dims[0])
			h, _ := strconv.Atoi(dims[1])

			reqsStr := strings.Fields(strings.TrimSpace(parts[1]))
			reqs := make([]int, len(reqsStr))
			for i, s := range reqsStr {
				reqs[i], _ = strconv.Atoi(s)
			}

			regions = append(regions, Region{Width: w, Height: h, Reqs: reqs})
			continue
		}

		row := make([]bool, len(line))
		for i, ch := range line {
			if ch == '#' {
				row[i] = true
			}
		}
		currentShapePixels = append(currentShapePixels, row)
	}
	if len(currentShapePixels) > 0 {
		shapes = append(shapes, createShape(len(shapes), currentShapePixels))
	}

	variants := make([][]Shape, len(shapes))
	for i, s := range shapes {
		variants[i] = generateVariants(s)
	}

	count := 0
	for _, r := range regions {
		if solve(r, shapes, variants) {
			count++
		}
	}
	fmt.Printf("Part 1: %d\n", count)
}

func createShape(id int, pixels [][]bool) Shape {
	size := 0
	for _, row := range pixels {
		for _, val := range row {
			if val {
				size++
			}
		}
	}
	return Shape{ID: id, Pixels: pixels, Size: size}
}

func generateVariants(s Shape) []Shape {
	unique := make(map[string]Shape)
	current := s.Pixels
	for i := 0; i < 4; i++ {
		addVariant(unique, s.ID, current)
		flipped := flip(current)
		addVariant(unique, s.ID, flipped)
		current = rotate(current)
	}
	res := make([]Shape, 0, len(unique))
	for _, v := range unique {
		res = append(res, v)
	}
	return res
}

func addVariant(unique map[string]Shape, id int, pixels [][]bool) {
	trimmed := trim(pixels)
	key := shapeToString(trimmed)
	unique[key] = createShape(id, trimmed)
}

func trim(pixels [][]bool) [][]bool {
	if len(pixels) == 0 {
		return pixels
	}
	minR, maxR := len(pixels), -1
	minC, maxC := len(pixels[0]), -1

	for r, row := range pixels {
		for c, val := range row {
			if val {
				if r < minR {
					minR = r
				}
				if r > maxR {
					maxR = r
				}
				if c < minC {
					minC = c
				}
				if c > maxC {
					maxC = c
				}
			}
		}
	}

	if maxR == -1 {
		return [][]bool{}
	}

	h := maxR - minR + 1
	w := maxC - minC + 1
	res := make([][]bool, h)
	for r := 0; r < h; r++ {
		res[r] = make([]bool, w)
		for c := 0; c < w; c++ {
			res[r][c] = pixels[r+minR][c+minC]
		}
	}
	return res
}

func shapeToString(pixels [][]bool) string {
	var sb strings.Builder
	for _, row := range pixels {
		for _, val := range row {
			if val {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func rotate(pixels [][]bool) [][]bool {
	h := len(pixels)
	w := len(pixels[0])
	res := make([][]bool, w)
	for r := 0; r < w; r++ {
		res[r] = make([]bool, h)
		for c := 0; c < h; c++ {
			res[r][c] = pixels[h-1-c][r]
		}
	}
	return res
}

func flip(pixels [][]bool) [][]bool {
	h := len(pixels)
	w := len(pixels[0])
	res := make([][]bool, h)
	for r := 0; r < h; r++ {
		res[r] = make([]bool, w)
		copy(res[r], pixels[r])
		for c := 0; c < w/2; c++ {
			res[r][c], res[r][w-1-c] = res[r][w-1-c], res[r][c]
		}
	}
	return res
}

func solve(region Region, shapes []Shape, variants [][]Shape) bool {
	totalArea := 0
	toPlace := []int{}
	for id, count := range region.Reqs {
		for range count {
			toPlace = append(toPlace, id)
			totalArea += shapes[id].Size
		}
	}

	if totalArea > region.Width*region.Height {
		return false
	}

	sort.Ints(toPlace)

	grid := make([][]bool, region.Height)
	for i := range grid {
		grid[i] = make([]bool, region.Width)
	}

	return backtrack(grid, toPlace, variants, -1, 0)
}

func backtrack(grid [][]bool, toPlace []int, variants [][]Shape, prevPieceID int, prevFlatIdx int) bool {
	if len(toPlace) == 0 {
		return true
	}

	pieceID := toPlace[0]
	remaining := toPlace[1:]

	startIdx := 0
	if pieceID == prevPieceID {
		startIdx = prevFlatIdx
	}

	H := len(grid)
	W := len(grid[0])

	pieceVariants := variants[pieceID]

	for _, v := range pieceVariants {
		vH := len(v.Pixels)
		vW := len(v.Pixels[0])

		for i := startIdx; i < H*W; i++ {
			r := i / W
			c := i % W

			if r+vH > H || c+vW > W {
				continue
			}

			if canPlace(grid, v.Pixels, r, c) {
				place(grid, v.Pixels, r, c, true)
				if backtrack(grid, remaining, variants, pieceID, i) {
					return true
				}
				place(grid, v.Pixels, r, c, false)
			}
		}
	}
	return false
}

func canPlace(grid [][]bool, pixels [][]bool, r, c int) bool {
	for pr, row := range pixels {
		for pc, val := range row {
			if val {
				if grid[r+pr][c+pc] {
					return false
				}
			}
		}
	}
	return true
}

func place(grid [][]bool, pixels [][]bool, r, c int, val bool) {
	for pr, row := range pixels {
		for pc, v := range row {
			if v {
				grid[r+pr][c+pc] = val
			}
		}
	}
}
