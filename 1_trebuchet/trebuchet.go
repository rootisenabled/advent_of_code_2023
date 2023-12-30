package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"utils"
)

func Substring(str string, start, end int) string {
	return strings.TrimSpace(str[start:end])
}

func findFirstAndLastDigit(input string) []string {
	digitsMap := map[string]int{
		"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4,
		"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	firstMatch, lastMatch := "", ""

	for i, char := range input {
		if char >= '0' && char <= '9' {
			digit := string(char)
			if firstMatch == "" {
				firstMatch = digit
			} else {
				lastMatch = digit
			}
		} else {
			for word, dig := range digitsMap {
				if strings.HasPrefix(input[i:], word) {
					if firstMatch == "" {
						firstMatch = fmt.Sprint(dig)
					} else {
						lastMatch = fmt.Sprint(dig)
					}
				}
			}
		}
	}

	if (lastMatch == "") && (firstMatch != "") {
		lastMatch = firstMatch
	}

	return []string{firstMatch, lastMatch}
}

func calculate(document []string) int {

	var sum int
	for _, line := range document {
		digits := findFirstAndLastDigit(line)

		var first string = digits[0]
		var last string = digits[len(digits)-1]
		num, _ := strconv.Atoi((first + last))
		sum += num
	}
	return sum
}

func main() {
	lines, err := utils.ReadLinesToSlice("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	sum := calculate(lines)
	fmt.Println(sum)
}
