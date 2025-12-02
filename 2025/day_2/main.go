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
		// an ID is invalid if it is made only of some sequence of digits repeated at least twice

		for i := startInt; i <= endInt; i++ {
			s := strconv.Itoa(i)
			n := len(s)
			valid := false
			for l := 1; l <= n/2; l++ {
				if n%l != 0 {
					continue
				}
				substr := s[0:l]
				repeated := true
				for j := l; j < n; j += l {
					if s[j:j+l] != substr {
						repeated = false
						break
					}
				}
				if repeated {
					valid = true
					break
				}
			}
			if valid {
				fmt.Fprintf(out, "Valid ID: %d\n", i)
				sum += i
			}
		}
	}

	fmt.Println("Sum:", sum)
}
