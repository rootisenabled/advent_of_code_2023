package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type Game struct {
	ID    int
	Cubes []Cube
}

type Cube map[string]int

func parseGames(input string) []Game {
	games := []Game{}
	gameStrings := strings.Split(input, "Game ")

	for _, gameString := range gameStrings[1:] {
		parts := strings.SplitN(gameString, ":", 2)
		gameID, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		cubes := []Cube{}

		subsets := strings.Split(parts[1], ";")
		for _, subset := range subsets {
			cubeStrings := strings.Split(subset, ",")
			for _, cubeString := range cubeStrings {
				parts := strings.Fields(cubeString)
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]
				cubes = append(cubes, Cube{color: count})
			}
		}

		games = append(games, Game{ID: gameID, Cubes: cubes})
	}

	return games
}

func getValidGameIds(games []Game) []int {
	RED := 12
	GREEN := 13
	BLUE := 14

	valid_ids := []int{}

	for _, game := range games {
		isValid := true
		for _, cube := range game.Cubes {
			if cube["red"] > RED || cube["green"] > GREEN || cube["blue"] > BLUE {
				isValid = false
				break
			}
		}
		if isValid {
			valid_ids = append(valid_ids, game.ID)
		}
	}

	return valid_ids
}

func getMinimumSetOfCubes(game Game) Cube {
	var (
		maxRed   = 0
		maxGreen = 0
		maxBlue  = 0
	)

	for _, cube := range game.Cubes {
		if cube["red"] > maxRed {
			maxRed = cube["red"]
		}
		if cube["green"] > maxGreen {
			maxGreen = cube["green"]
		}
		if cube["blue"] > maxBlue {
			maxBlue = cube["blue"]
		}
	}

	return Cube{
		"red":   maxRed,
		"green": maxGreen,
		"blue":  maxBlue,
	}
}

func getSumOfPoweredCubes(games []Game) int {
	sum := 0
	for _, game := range games {
		min := getMinimumSetOfCubes(game)
		sum += min["red"] * min["blue"] * min["green"]
	}
	return sum
}

func sumSlice(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + sumSlice(numbers[1:])
}

func main() {
	input, _ := utils.ReadLines("input.txt")
	games := parseGames(input)
	valid_ids := getValidGameIds(games)

	fmt.Println(sumSlice(valid_ids))
	fmt.Println(getSumOfPoweredCubes(games))
}
