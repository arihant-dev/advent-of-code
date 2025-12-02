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
	rawLines := strings.Split(content, "\n")

	// starting point is 50
	start := 50
	endpointCount := 0 // number of times position equals 0 after a rotation
	passCount := 0     // number of times 0 was reached during rotations (including passing through and landing)

	for _, line := range rawLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		dir := line[0]
		numStr := strings.TrimSpace(line[1:])
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		prev := start
		occurrences := 0

		switch dir {
		case 'R':
			// steps until next 0 when moving right
			occurrences = (prev + num) / 100
			start = (prev + num) % 100
			fmt.Fprintf(out, "R%d: %d -> %d (passed 0 %d times)\n", num, prev, start, occurrences)
		case 'L':
			// steps until next 0 when moving left
			// smallest positive step to hit 0 when moving left is t0 = prev (if prev>0) else 100
			t0 := prev
			if t0 == 0 {
				t0 = 100
			}
			if num >= t0 {
				occurrences = 1 + (num-t0)/100
			} else {
				occurrences = 0
			}
			start = (prev - num) % 100
			if start < 0 {
				start += 100
			}
			fmt.Fprintf(out, "L%d: %d -> %d (passed 0 %d times)\n", num, prev, start, occurrences)
		default:
			fmt.Fprintf(out, "skipping: %q\n", line)
			continue
		}

		passCount += occurrences
		if start == 0 {
			endpointCount++
		}
	}

	fmt.Fprintf(out, "\nTotal times passed through 0 during moves: %d\n", passCount)
	fmt.Fprintf(out, "Total times ended at 0 after a move: %d\n", endpointCount)
}
