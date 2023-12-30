package main

import (
	"encoding/json"
	"testing"
	"utils"
)

func Test(t *testing.T) {
	input, _ := utils.ReadLines("input_test.txt")

	sum := sumSlice(getValidGameIds(parseGames(input)))
	expected := 8

	if sum != expected {
		t.Errorf("Expected %d, got %d", expected, sum)
	}
}

func TestMinimumSetOfCubes(t *testing.T) {
	input, _ := utils.ReadLines("input_test.txt")

	games := parseGames(input)
	minimumSet := getMinimumSetOfCubes(games[0])
	expected := Cube{
		"red":   4,
		"green": 2,
		"blue":  6,
	}

	minimumSetJSON, _ := json.Marshal(minimumSet)
	expectedJSON, _ := json.Marshal(expected)

	if string(minimumSetJSON) != string(expectedJSON) {
		t.Errorf("Expected %v, got %v", expected, minimumSet)
	}
}

func TestGetSumOfPoweredCubes(t *testing.T) {
	input, _ := utils.ReadLines("input_test.txt")

	games := parseGames(input)
	sumOfPowered := getSumOfPoweredCubes(games)
	expected := 2286

	if sumOfPowered != expected {
		t.Errorf("Expected %d, got %d", expected, sumOfPowered)
	}
}
