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
