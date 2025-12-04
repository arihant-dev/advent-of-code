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

	// create a 2-d array of characters
	grid := [][]rune{}
	rawLines := strings.SplitSeq(content, "\n")
	for line := range rawLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		row := []rune(line)
		grid = append(grid, row)
	}
	// part_1(grid, out)
	part_2(grid, out)

}
func part_1(grid [][]rune, out *os.File) {
	count := 0
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			// check if eight adjacent positions are '@', and if the count is less than 4, add it to ans
			if grid[i][j] == '@' {
				count = 0
				// check all 8 directions
				directions := [][2]int{
					{-1, -1}, {-1, 0}, {-1, 1},
					{0, -1}, {0, 1},
					{1, -1}, {1, 0}, {1, 1},
				}
				for _, dir := range directions {
					ni := i + dir[0]
					nj := j + dir[1]
					if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[ni]) {
						if grid[ni][nj] == '@' {
							count++
						}
					}
				}
				if count < 4 {
					ans++
				}
			}
		}
	}
	fmt.Fprintf(out, "%d", ans)
}
func part_2(grid [][]rune, out *os.File) {
	count := 0
	ans := 0
	removeIndexes := [][]int{}
	loop := true
	x := 0
	for loop {
		x++
		fmt.Fprintf(out, "Starting iteration %d\n", x)
		for i := range grid {
			for j := range grid[i] {
				// check if eight adjacent positions are '@', and if the count is less than 4, add it to ans
				if grid[i][j] == '@' {
					count = 0
					// check all 8 directions
					directions := [][2]int{
						{-1, -1}, {-1, 0}, {-1, 1},
						{0, -1}, {0, 1},
						{1, -1}, {1, 0}, {1, 1},
					}
					for _, dir := range directions {
						ni := i + dir[0]
						nj := j + dir[1]
						if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[ni]) {
							if grid[ni][nj] == '@' {
								count++
							}
						}
					}
					if count < 4 {
						removeIndexes = append(removeIndexes, []int{i, j})
						ans++
					}
				}
			}
		}
		// remove all '@' at removeIndexes
		for _, idx := range removeIndexes {
			i := idx[0]
			j := idx[1]
			grid[i][j] = '.'
		}
		fmt.Fprintf(out, "grid after iteration %d\n", x)
		for _, row := range grid {
			fmt.Fprintf(out, "%s\n", string(row))
		}
		if len(removeIndexes) == 0 {
			loop = false
			break
		}
		removeIndexes = [][]int{}
	}
	fmt.Fprintf(out, "The answer for part 2 is %d", ans)
}
