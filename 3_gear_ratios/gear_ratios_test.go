package main

import (
	"fmt"
	"testing"
	"utils"
)

func TestGearRatios(t *testing.T) {
	input, _ := utils.ReadLinesToSlice("input.txt")

	var allNums []*EnhancedNum

	for lineIndex, line := range input {

		numbers := ExtractNumbers(line)
		for _, number := range numbers {
			number.lineIndex = lineIndex
		}

		allNums = append(allNums, numbers...)
	}

	sumGearRaios := 0

	for _, number := range allNums {
		if isSymbolAround(number, input) {
			sumGearRaios += number.value
		}
	}

	fmt.Printf("Sum of gear ratios: %d\n", sumGearRaios)
}
