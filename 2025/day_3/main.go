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
	// read each line as array of integers
	arrayOfBanks := [][]int64{}
	// normalize newlines and split into lines
	lines := strings.SplitSeq(strings.ReplaceAll(string(f), "\r\n", "\n"), "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// nums are 1-9 separated by nothing
		bankStrs := strings.Split(line, "")
		bank := []int64{}
		for _, bs := range bankStrs {
			num, err := strconv.ParseInt(bs, 10, 64)
			if err != nil {
				panic(err)
			}
			bank = append(bank, num)
		}
		arrayOfBanks = append(arrayOfBanks, bank)
	}
	sumOfMaxBatteries := int64(0)
	for _, bank := range arrayOfBanks {
		if len(bank) == 0 {
			continue
		}
		// Choose up to 12 digits from the bank (preserving order) to form the largest possible number.
		k := min(len(bank), 12)
		pos := 0
		var maxBattery int64 = 0
		for i := 0; i < k; i++ {
			// end index (inclusive) for this selection
			end := len(bank) - (k - i)
			if end < pos {
				end = pos
			}
			// find max digit in bank[pos..end]
			best := bank[pos]
			bestIdx := pos
			for j := pos; j <= end; j++ {
				if bank[j] > best {
					best = bank[j]
					bestIdx = j
				}
			}
			maxBattery = maxBattery*10 + best
			pos = bestIdx + 1
		}
		fmt.Fprintf(out, "Max battery: %d from bank %v\n", maxBattery, bank)
		sumOfMaxBatteries += maxBattery
	}
	fmt.Fprintf(out, "Sum of max batteries: %d\n", sumOfMaxBatteries)

}
