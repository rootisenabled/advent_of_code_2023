package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// read line by line into memory
// all file contents is stores in lines[]
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func calculate(document []string) int {
	var sum int
	for _, line := range document {
		var digits []string
		for _, symbol := range line {
			if unicode.IsDigit(symbol) {
				digits = append(digits, string(symbol))

			}
		}
		var first string = digits[0]
		var last string = digits[len(digits)-1]
		num, _ := strconv.Atoi((first + last))
		sum += num
	}
	return sum
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	sum := calculate(lines)
	fmt.Println(sum)
}
