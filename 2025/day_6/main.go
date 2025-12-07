package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	out, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	part2(lines, out)
}

func part2(lines []string, out *os.File) {
	n := len(lines)
	if n == 0 {
		return
	}

	// Determine max line length (m)
	m := 0
	for _, line := range lines {
		if len(line) > m {
			m = len(line)
		}
	}

	totalSum := int64(0)
	col := 0

	for col < m {
		// Check if the current column is empty (separator)
		isEmpty := true
		for row := 0; row < n; row++ {
			if col < len(lines[row]) && lines[row][col] != ' ' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			col++
			continue
		}

		// Identify the block of the problem (contiguous non-empty columns)
		problemStart := col
		problemEnd := col
		for problemEnd < m {
			colEmpty := true
			for row := 0; row < n; row++ {
				if problemEnd < len(lines[row]) && lines[row][problemEnd] != ' ' {
					colEmpty = false
					break
				}
			}
			if colEmpty {
				break
			}
			problemEnd++
		}

		// Find the operator in the last row of this block
		var op byte
		for c := problemStart; c < problemEnd; c++ {
			if c < len(lines[n-1]) {
				ch := lines[n-1][c]
				if ch == '+' || ch == '*' {
					op = ch
					break
				}
			}
		}

		// Calculate result for this problem block
		var problemResult int64
		if op == '+' {
			problemResult = 0
		} else {
			problemResult = 1
		}

		// Iterate columns right-to-left (as per C code / problem description)
		for c := problemEnd - 1; c >= problemStart; c-- {
			var numStrBuilder strings.Builder

			// Extract digits from top to bottom (excluding the last operator row)
			for row := 0; row < n-1; row++ {
				if c < len(lines[row]) {
					ch := lines[row][c]
					if ch >= '0' && ch <= '9' {
						numStrBuilder.WriteByte(ch)
					}
				}
			}

			numStr := numStrBuilder.String()
			if len(numStr) > 0 {
				num, err := strconv.ParseInt(numStr, 10, 64)
				if err != nil {
					panic(err)
				}

				if op == '+' {
					problemResult += num
				} else if op == '*' {
					problemResult *= num
				}
			}
		}

		fmt.Fprintf(out, "Problem at cols [%d, %d): op='%c', result=%d\n", problemStart, problemEnd, op, problemResult)
		totalSum += problemResult

		// Move to the next block
		col = problemEnd
	}

	fmt.Fprintf(out, "\nTotal sum of all columns in part 2: %d\n", totalSum)
	fmt.Printf("Total sum: %d\n", totalSum)
}
