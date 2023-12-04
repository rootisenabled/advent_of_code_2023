package main

import (
	"testing"
)

func TestTrebuchet(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	sum := calculate(input)
	if sum != 142 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 142)
	}
}
