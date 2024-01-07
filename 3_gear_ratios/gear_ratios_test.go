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

func TestGearRatiosPart2(t *testing.T) {
	input, _ := utils.ReadLinesToSlice("input.txt")

	var allNums []*EnhancedNum

	for lineIndex, line := range input {

		numbers := ExtractNumbers(line)
		for _, number := range numbers {
			number.lineIndex = lineIndex
		}

		allNums = append(allNums, numbers...)
	}

	totalSum := 0

	for lineIndexOfStar, line := range input {

		starPositions := FindCharacterPositions(line, '*')

		for _, starPosition := range starPositions {
			var gearNumbers []*EnhancedNum
			for _, number := range allNums {
				if IsNumberTouchingCharacter(number, lineIndexOfStar, starPosition, len(input)) {
					gearNumbers = append(gearNumbers, number)
				}
			}
			if len(gearNumbers) == 2 {
				totalSum += gearNumbers[0].value * gearNumbers[1].value
			}
		}
	}

	fmt.Printf("Sum of gear ratios: %d\n", totalSum)
}
