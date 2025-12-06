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
	freshBatches := [][]int{}
	countOfFreshBatchesPresentInStock := 0
	for line := range rawLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Fprintf(out, "Processing line: %s\n", line)
		if strings.Contains(line, "-") {
			// create a hashmap inclusive of the range
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				panic("invalid range: " + line)
			}
			startInt := 0
			endInt := 0
			_, err := fmt.Sscanf(strings.TrimSpace(parts[0]), "%d", &startInt)
			if err != nil {
				panic(err)
			}
			_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%d", &endInt)
			if err != nil {
				panic(err)
			}
			// store these ranges in the array freshBatches
			freshBatches = append(freshBatches, []int{startInt, endInt})
		} else {
			// single integer
			num := 0
			_, err := fmt.Sscanf(line, "%d", &num)
			if err != nil {
				panic(err)
			}
			// loop through freshBatches and check if num is in any range
			found := false
			for _, batch := range freshBatches {
				if num >= batch[0] && num <= batch[1] {
					found = true
					break
				}
			}
			if found {
				countOfFreshBatchesPresentInStock++
				fmt.Fprintf(out, "Number %d is present in fresh batches\n", num)
			} else {
				fmt.Fprintf(out, "Number %d is NOT present in fresh batches\n", num)
			}
		}
	}
	part_2(freshBatches, out)
	fmt.Println("Count of fresh batches present in stock:", countOfFreshBatchesPresentInStock)
}
func part_2(freshBatches [][]int, out *os.File) {
	// find total number of unique integers covered by freshBatches
	// merge intervals
	if len(freshBatches) == 0 {
		fmt.Fprintf(out, "Total unique integers covered by fresh batches: 0\n")
		return
	}
	// sort freshBatches by start
	sort.Slice(freshBatches, func(i, j int) bool {
		return freshBatches[i][0] < freshBatches[j][0]
	})
	merged := [][]int{}
	current := freshBatches[0]
	for i := 1; i < len(freshBatches); i++ {
		if freshBatches[i][0] <= current[1]+1 {
			// overlap
			if freshBatches[i][1] > current[1] {
				current[1] = freshBatches[i][1]
			}
		} else {
			merged = append(merged, current)
			current = freshBatches[i]
		}
	}
	merged = append(merged, current)

	// calculate total unique integers
	totalUnique := 0
	for _, interval := range merged {
		totalUnique += (interval[1] - interval[0] + 1)
	}
	fmt.Fprintf(out, "Total unique integers covered by fresh batches: %d\n", totalUnique)
}
