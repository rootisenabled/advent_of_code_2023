package main

import (
	"regexp"
	"strconv"
	"strings"
)

type EnhancedNum struct {
	value         int
	startPosition int
	endPosition   int
	lineIndex     int
}

func isSymbolAround(number *EnhancedNum, lines []string) bool {
	fromIndex := number.startPosition - 1
	if fromIndex < 0 {
		fromIndex = 0
	}
	toIndex := number.endPosition + 1
	if toIndex > len(lines) {
		toIndex = len(lines)
	}
	for index := number.lineIndex - 1; index <= number.lineIndex+1; index++ {
		if index < 0 || index >= len(lines) {
			continue
		}

		found := strings.IndexAny(lines[index][fromIndex:toIndex], "+#$*@/=%-&")
		if found > -1 {
			return true
		}
	}
	return false

}

func ExtractNumbers(input string) []*EnhancedNum {
	re := regexp.MustCompile(`\d+`)

	match := re.FindAllStringSubmatchIndex(input, -1)

	var result []*EnhancedNum

	for _, m := range match {
		value, _ := strconv.Atoi(input[m[0]:m[1]])
		result = append(result, &EnhancedNum{
			value:         value,
			startPosition: m[0],
			endPosition:   m[1],
		})
	}

	return result
}

func IsNumberTouchingCharacter(number *EnhancedNum, lineIndex int, charPostion int, numberOfLines int) bool {
	for index := lineIndex - 1; index <= lineIndex+1; index++ {
		if index < 0 || index >= numberOfLines {
			continue
		}

		if number.lineIndex != index {
			continue
		}

		if number.startPosition == charPostion || number.startPosition == charPostion+1 {
			return true
		}

		if number.endPosition == charPostion || number.endPosition-1 == charPostion {
			return true
		}

		if number.lineIndex != lineIndex {
			if number.startPosition <= charPostion && number.endPosition >= charPostion {
				return true
			}
		}

	}
	return false
}

func FindCharacterPositions(input string, character rune) []int {
	var result []int

	for index, char := range input {
		if char == character {
			result = append(result, index)
		}
	}

	return result
}
