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
	rawLines := strings.Split(content, "\n")
	arrayOfLists := [][]string{}

	// Process each line
	for _, line := range rawLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Fprintf(out, "Processing line: %s\n", line)
		// Split the line into words and sort them
		words := strings.Fields(line)
		arrayOfLists = append(arrayOfLists, words)
	}
	m := len(arrayOfLists)
	n := len(arrayOfLists[0])
	fmt.Fprintf(out, "The array of Lists if of %d-length and %d-breadth", n, m)
	// part_1(arrayOfLists, out)
	part_2(arrayOfLists, out)

}
func part_1(arrayOfLists [][]string, out *os.File) {
	// we have to compute the answers vertically, column-wise
	m := len(arrayOfLists)
	n := len(arrayOfLists[0])
	answerOfLists := []int{}
	for col := 0; col < n; col++ {
		fmt.Fprintf(out, "\nProcessing column %d", col+1)
		sign := arrayOfLists[m-1][col]
		if sign == "+" {
			sum := 0
			for row := 0; row < m-1; row++ {
				fmt.Fprintf(out, "\nAdding element %d: %s", row, arrayOfLists[row][col])
				num := 0
				_, err := fmt.Sscanf(arrayOfLists[row][col], "%d", &num)
				if err != nil {
					panic(err)
				}
				sum += num
			}
			answerOfLists = append(answerOfLists, sum)
		} else {
			product := 1
			for row := 0; row < m-1; row++ {
				fmt.Fprintf(out, "\nMultiplying element %d: %s", row, arrayOfLists[row][col])
				num := 0
				_, err := fmt.Sscanf(arrayOfLists[row][col], "%d", &num)
				if err != nil {
					panic(err)
				}
				product *= num
			}
			answerOfLists = append(answerOfLists, product)
		}
	}

	finalSum := 0
	for _, ans := range answerOfLists {
		finalSum += ans
	}
	fmt.Fprintf(out, "\nThe final sum of answers is: %d\n", finalSum)
}
func part_2(arrayOfLists [][]string, out *os.File) {
	m := len(arrayOfLists)
	n := len(arrayOfLists[0])
	finalArrayOfLists := []int64{}
	for col := 0; col < n; col++ {
		fmt.Fprintf(out, "\nProcessing column %d in part 2", col+1)
		sign := arrayOfLists[m-1][col]
		numbers := []int{}
		for row := 0; row < m-1; row++ {
			var num int
			_, _ = fmt.Sscanf(arrayOfLists[row][col], "%d", &num)
			numbers = append(numbers, num)
		}
		// Convert numbers to strings to easily access digits by position
		numStrs := make([]string, len(numbers))
		maxDigits := 0
		for i, num := range numbers {
			s := fmt.Sprintf("%d", num)
			numStrs[i] = s
			if len(s) > maxDigits {
				maxDigits = len(s)
			}
		}

		fmt.Fprintf(out, "\nMax digits in column %d: %d", col+1, maxDigits)
		fmt.Fprintf(out, "\nNumbers in column %d: %v", col+1, numbers)

		newNumbers := []int64{}

		// Construct new numbers vertically based on alignment
		if sign == "+" {
			// Left alignment
			for digitIdx := 0; digitIdx < maxDigits; digitIdx++ {
				currentVal := int64(0)
				hasDigits := false
				for _, s := range numStrs {
					if digitIdx < len(s) {
						digit := int64(s[digitIdx] - '0')
						currentVal = currentVal*10 + digit
						hasDigits = true
					}
				}
				if hasDigits {
					newNumbers = append(newNumbers, currentVal)
				}
			}
		} else {
			// Right alignment
			for digitIdx := 0; digitIdx < maxDigits; digitIdx++ {
				currentVal := int64(0)
				hasDigits := false
				for _, s := range numStrs {
					charIdx := len(s) - 1 - digitIdx
					if charIdx >= 0 {
						digit := int64(s[charIdx] - '0')
						currentVal = currentVal*10 + digit
						hasDigits = true
					}
				}
				if hasDigits {
					newNumbers = append(newNumbers, currentVal)
				}
			}
		}

		fmt.Fprintf(out, "\n New Numbers in column %d: %v", col+1, newNumbers)
		fmt.Printf("Column %d: %v -> ", col+1, newNumbers)

		if sign == "+" {
			sum := int64(0)
			for _, v := range newNumbers {
				sum += v
			}
			fmt.Fprintf(out, "\nSum of reversed numbers in column %d: %d", col+1, sum)
			fmt.Printf("Sum: %d\n", sum)
			finalArrayOfLists = append(finalArrayOfLists, sum)
		} else {
			product := int64(1)
			for _, v := range newNumbers {
				product *= v
			}
			fmt.Fprintf(out, "\nProduct of reversed numbers in column %d: %d", col+1, product)
			fmt.Printf("Product: %d\n", product)
			finalArrayOfLists = append(finalArrayOfLists, product)
		}
	}
	totalSum := int64(0)
	for _, val := range finalArrayOfLists {
		totalSum += val
	}
	fmt.Fprintf(out, "\nTotal sum of all columns in part 2: %d\n", totalSum)
}
