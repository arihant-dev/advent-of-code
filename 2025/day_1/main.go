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

	content := strings.ReplaceAll(string(f), "\r\n", "\n")
	rawLines := strings.Split(content, "\n")

	// starting point is 50
	start := 50
	count := 0
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

		switch dir {
		case 'R':
			prev := start
			start = (start + num) % 100
			fmt.Fprintf(out, "R%d: %d -> %d\n", num, prev, start)
		case 'L':
			prev := start
			start = (start - num) % 100
			if start < 0 {
				start += 100
			}
			fmt.Fprintf(out, "L%d: %d -> %d\n", num, prev, start)
		default:
			fmt.Fprintf(out, "unknown direction")
		}

		if start == 0 {
			count++
		}
	}
	fmt.Fprintln(out, count)
}
