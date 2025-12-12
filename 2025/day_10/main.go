package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read filename from the args
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	// read the file
	f, err := os.ReadFile(filename)
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
	lights := make([][]int, 0)
	buttonSequences := make([][][]int, 0)
	joltageRatingsList := make([][]int, 0)
	for _, line := range rawLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
		// array of lights is inside [], . means off, # means on
		parts := strings.Split(line, " ")
		lightsPart := parts[0]
		lightsPart = strings.TrimPrefix(lightsPart, "[")
		lightsPart = strings.TrimSuffix(lightsPart, "]")
		light := make([]int, 0)
		for _, ch := range lightsPart {
			if ch == '#' {
				light = append(light, 1)
			} else {
				light = append(light, 0)
			}
		}
		lights = append(lights, light)

		// button sequences are in the remaining parts before the {
		// then after space, we have () with comma separated integers
		buttonSeq := make([][]int, 0)
		for _, p := range parts[1:] {
			p = strings.TrimSpace(p)
			if strings.HasPrefix(p, "{") {
				break
			}
			p = strings.TrimPrefix(p, "(")
			p = strings.TrimSuffix(p, ")")
			numsStr := strings.Split(p, ",")
			nums := make([]int, 0)
			for _, ns := range numsStr {
				ns = strings.TrimSpace(ns)
				if ns == "" {
					continue
				}
				n, _ := strconv.Atoi(ns)
				nums = append(nums, n)
			}
			buttonSeq = append(buttonSeq, nums)
		}
		buttonSequences = append(buttonSequences, buttonSeq)

		// parse the joltage ratings into a slice of integers
		joltageParts := strings.Split(parts[len(parts)-1], ",")
		joltageRatings := make([]int, 0)
		for _, jp := range joltageParts {
			jp = strings.TrimSpace(jp)
			jp = strings.TrimPrefix(jp, "{")
			jp = strings.TrimSuffix(jp, "}")
			if jp == "" {
				continue
			}
			j, _ := strconv.Atoi(jp)
			joltageRatings = append(joltageRatings, j)
		}
		joltageRatingsList = append(joltageRatingsList, joltageRatings)
	}

	part_1(lights, buttonSequences, out)
	part_2(buttonSequences, joltageRatingsList, out)
}
func part_1(lights [][]int, buttonSeqs [][][]int, out *os.File) {
	totalPresses := 0
	for i := range lights {
		target := lights[i]
		buttons := buttonSeqs[i]
		n := len(target)
		numButtons := len(buttons)
		minPresses := -1

		// Brute force all combinations of buttons
		// 1 << numButtons
		for mask := 0; mask < (1 << numButtons); mask++ {
			// Current state starts at all 0s
			current := make([]int, n)
			pressCount := 0

			for b := range numButtons {
				if (mask & (1 << b)) != 0 {
					pressCount++
					// Apply button b
					for _, lightIdx := range buttons[b] {
						if lightIdx >= 0 && lightIdx < n {
							current[lightIdx] ^= 1 // Toggle
						}
					}
				}
			}

			// Check if current matches target
			match := true
			for k := range n {
				if current[k] != target[k] {
					match = false
					break
				}
			}

			if match {
				if minPresses == -1 || pressCount < minPresses {
					minPresses = pressCount
				}
			}
		}

		if minPresses != -1 {
			fmt.Fprintf(out, "Row %d: Minimum button presses = %d\n", i+1, minPresses)
			totalPresses += minPresses
		} else {
			fmt.Fprintf(out, "Row %d: Impossible\n", i+1)
		}
	}
	fmt.Fprintf(out, "Part 1: Total presses = %d\n", totalPresses)
}
func part_2(buttonSeqs [][][]int, joltageRatingsList [][]int, out *os.File) {
	totalPresses := 0
	for i := range joltageRatingsList {
		targets := joltageRatingsList[i]
		buttons := buttonSeqs[i]

		// Build Matrix
		rows := len(targets)
		cols := len(buttons)
		A := make([][]float64, rows)
		b := make([]float64, rows)
		for r := 0; r < rows; r++ {
			A[r] = make([]float64, cols)
			b[r] = float64(targets[r])
		}
		for c, btn := range buttons {
			for _, r := range btn {
				if r >= 0 && r < rows {
					A[r][c] = 1.0
				}
			}
		}

		minPresses := solveGaussianInteger(A, b)

		if minPresses != -1 {
			fmt.Fprintf(out, "Part 2: Machine %d: Minimum button presses = %d\n", i+1, minPresses)
			totalPresses += minPresses
		} else {
			fmt.Fprintf(out, "Part 2: Machine %d: Impossible\n", i+1)
		}
	}
	fmt.Fprintf(out, "Part 2: Total presses = %d\n", totalPresses)
}

func solveGaussianInteger(A [][]float64, b []float64) int {
	rows := len(A)
	cols := len(A[0])

	// Gauss-Jordan
	pivotRow := 0
	colToRow := make(map[int]int)

	for c := 0; c < cols && pivotRow < rows; c++ {
		sel := -1
		for r := pivotRow; r < rows; r++ {
			if math.Abs(A[r][c]) > 1e-5 {
				sel = r
				break
			}
		}
		if sel == -1 {
			continue
		}

		A[pivotRow], A[sel] = A[sel], A[pivotRow]
		b[pivotRow], b[sel] = b[sel], b[pivotRow]

		div := A[pivotRow][c]
		for j := c; j < cols; j++ {
			A[pivotRow][j] /= div
		}
		b[pivotRow] /= div

		for r := 0; r < rows; r++ {
			if r != pivotRow {
				factor := A[r][c]
				if math.Abs(factor) > 1e-5 {
					for j := c; j < cols; j++ {
						A[r][j] -= factor * A[pivotRow][j]
					}
					b[r] -= factor * b[pivotRow]
				}
			}
		}
		colToRow[c] = pivotRow
		pivotRow++
	}

	for r := pivotRow; r < rows; r++ {
		if math.Abs(b[r]) > 1e-5 {
			return -1
		}
	}

	var freeVars []int
	for c := 0; c < cols; c++ {
		if _, ok := colToRow[c]; !ok {
			freeVars = append(freeVars, c)
		}
	}

	// Determine limit for free variables
	limit := 2000

	minTotal := -1
	assignment := make([]int, cols)

	var backtrack func(idx int, currentFreeSum int)
	backtrack = func(idx int, currentFreeSum int) {
		if minTotal != -1 && currentFreeSum >= minTotal {
			return
		}

		if idx == len(freeVars) {
			// Compute basic vars
			currentSum := currentFreeSum
			valid := true

			for c := 0; c < cols; c++ {
				if row, isBasic := colToRow[c]; isBasic {
					val := b[row]
					for _, f := range freeVars {
						val -= A[row][f] * float64(assignment[f])
					}

					if val < -1e-5 {
						valid = false
						break
					}
					intVal := int(math.Round(val))
					if math.Abs(val-float64(intVal)) > 1e-5 {
						valid = false
						break
					}

					assignment[c] = intVal
					currentSum += intVal
					if minTotal != -1 && currentSum >= minTotal {
						valid = false
						break
					}
				}
			}

			if valid {
				if minTotal == -1 || currentSum < minTotal {
					minTotal = currentSum
				}
			}
			return
		}

		fVar := freeVars[idx]
		for val := 0; val <= limit; val++ {
			assignment[fVar] = val
			backtrack(idx+1, currentFreeSum+val)
		}
	}

	backtrack(0, 0)
	return minTotal
}
