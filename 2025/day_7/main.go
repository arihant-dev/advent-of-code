package main

import (
	"fmt"
	"os"
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
	tachyonManifold := [][]rune{}
	for line := range rawLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		tachyonManifold = append(tachyonManifold, []rune(line))
	}
	part_1(tachyonManifold, out)
	part_2(tachyonManifold, out)

}
func part_1(tachyonManifold [][]rune, out *os.File) {
	rows := len(tachyonManifold)
	cols := len(tachyonManifold[0])
	splits := 0
	startPoint := 0
	for c := range cols {
		if tachyonManifold[0][c] == 'S' {
			startPoint = c
			break
		}
	}
	beams := make(map[int]bool)
	beams[startPoint] = true
	for r := 1; r < rows; r++ {
		newBeams := make(map[int]bool)
		for c := range beams {
			if c < 0 || c >= cols {
				continue
			}
			cell := tachyonManifold[r][c]
			switch cell {
			case '.':
				newBeams[c] = true
			case '^':
				splits++
				newBeams[c-1] = true
				newBeams[c+1] = true
			}
		}
		beams = newBeams
	}
	fmt.Fprintf(out, "Part 1: Total splits = %d\n", splits)
}
func part_2(tachyonManifold [][]rune, out *os.File) {
	rows := len(tachyonManifold)
	if rows == 0 {
		return
	}
	cols := len(tachyonManifold[0])

	startPoint := -1
	for c := 0; c < cols; c++ {
		if tachyonManifold[0][c] == 'S' {
			startPoint = c
			break
		}
	}

	if startPoint == -1 {
		fmt.Fprintf(out, "Start point 'S' not found")
		return
	}

	// Use a map to track the number of beams (timelines) at each column index
	beams := make(map[int]int)
	beams[startPoint] = 1

	// Iterate through each row starting from 1
	for r := 1; r < rows; r++ {
		nextBeams := make(map[int]int)

		for c, count := range beams {
			// Boundary check
			if c < 0 || c >= cols {
				continue
			}

			cell := tachyonManifold[r][c]

			switch cell {
			case '.':
				nextBeams[c] += count
			case '^':
				nextBeams[c-1] += count
				nextBeams[c+1] += count
			}
		}
		beams = nextBeams
	}

	totalTimelines := 0
	for _, count := range beams {
		totalTimelines += count
	}

	fmt.Fprintf(out, "Part 2: Total active timelines = %d\n", totalTimelines)
}
