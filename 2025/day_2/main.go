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

	// Get ranges (comma-separated). Allow whitespace/newlines around entries.
	content := strings.ReplaceAll(string(f), "\r\n", "\n")
	parts := strings.Split(content, ",")

	sum := 0

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		bounds := strings.Split(part, "-")
		if len(bounds) != 2 {
			panic("invalid range: " + part)
		}
		startInt, err := strconv.Atoi(strings.TrimSpace(bounds[0]))
		if err != nil {
			panic(err)
		}
		endInt, err := strconv.Atoi(strings.TrimSpace(bounds[1]))
		if err != nil {
			panic(err)
		}

		for i := startInt; i <= endInt; i++ {
			s := strconv.Itoa(i)
			n := len(s)
			// only consider even-length numbers where first half equals second half
			if n%2 != 0 {
				continue
			}
			half := n / 2
			if s[:half] == s[half:] {
				sum += i
				fmt.Fprintln(out, s)
			}
		}
	}

	fmt.Println("Sum:", sum)
}
