package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// read filename from the args
	filename := "input_final.txt"
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

	rawLines := strings.SplitSeq(content, "\n")

	// create the adjacency list
	adjacencyList := make(map[string][]string)
	for line := range rawLines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ":")
		node := strings.TrimSpace(parts[0])
		neighbors := strings.Split(strings.TrimSpace(parts[1]), " ")
		adjacencyList[node] = neighbors
	}

	// call part 1 and part 2 functions

	part_1(out, adjacencyList)
	part_2(out, adjacencyList)
}
func part_1(out *os.File, adjacencyList map[string][]string) {
	start := "you"
	end := "out"
	//find number of paths from start to end using DP (Memoization)
	memo := make(map[string]int)

	var dp func(node string) int
	dp = func(node string) int {
		if node == end {
			return 1
		}
		if val, ok := memo[node]; ok {
			return val
		}

		count := 0
		for _, neighbor := range adjacencyList[node] {
			count += dp(neighbor)
		}
		memo[node] = count
		return count
	}

	pathCount := dp(start)
	fmt.Fprintf(out, "Number of paths from %s to %s: %d\n", start, end, pathCount)
}
func part_2(out *os.File, adjacencyList map[string][]string) {
	start := "svr"
	end := "out"
	specialNodes := []string{"dac", "fft"}
	specialMap := make(map[string]int)
	for i, node := range specialNodes {
		specialMap[node] = i
	}
	targetMask := (1 << len(specialNodes)) - 1

	// Memoization cache: "node|mask", value is count
	memo := make(map[string]int)

	var dp func(node string, mask int) int
	dp = func(node string, mask int) int {
		// Update mask if current node is special
		if idx, ok := specialMap[node]; ok {
			mask |= (1 << idx)
		}

		if node == end {
			if mask == targetMask {
				return 1
			}
			return 0
		}

		key := fmt.Sprintf("%s|%d", node, mask)
		if val, ok := memo[key]; ok {
			return val
		}

		count := 0
		for _, neighbor := range adjacencyList[node] {
			count += dp(neighbor, mask)
		}
		memo[key] = count
		return count
	}

	pathCount := dp(start, 0)
	fmt.Fprintf(out, "Number of paths from %s to %s visiting all special nodes: %d\n", start, end, pathCount)
}
