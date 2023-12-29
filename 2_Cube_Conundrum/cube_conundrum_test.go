package main

import (
	"testing"
)

func Test(t *testing.T) {
	input, _ := readLines("input_test.txt")

	sum := sumSlice(getValidGameIds(parseGames(input)))
	expected := 8

	if sum != expected {
		t.Errorf("Expected %d, got %d", expected, sum)
	}
}
