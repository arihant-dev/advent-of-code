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
	arrayOfBanks := [][]int{}
	lines := strings.Split(strings.ReplaceAll(string(f), "\r\n", "\n"), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// nums are 1-9 seperated by nothing
		bankStrs := strings.Split(line, "")
		bank := []int{}
		for _, bs := range bankStrs {
			num, err := strconv.Atoi(bs)
			if err != nil {
				panic(err)
			}
			bank = append(bank, num)
		}
		arrayOfBanks = append(arrayOfBanks, bank)
	}
	sumOfMaxBatteries := 0
	for _, bank := range arrayOfBanks {
		maxBattery := 0
		for i := range bank {
			for j := i + 1; j < len(bank); j++ {
				if bank[i]*10+bank[j] > maxBattery {
					maxBattery = bank[i]*10 + bank[j]
				}
			}
		}
		fmt.Fprintf(out, "Max battery: %d from bank %v\n", maxBattery, bank)
		sumOfMaxBatteries += maxBattery
	}
	fmt.Fprintf(out, "Sum of max batteries: %d\n", sumOfMaxBatteries)

}
