package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data), nil
}

type Game struct {
	ID    int
	Cubes []Cube
}

type Cube map[string]int

func parseGames(input string) []Game {
	games := []Game{}
	gameStrings := strings.Split(input, "Game ")

	for _, gameString := range gameStrings[1:] {
		game := Game{
			Cubes: []Cube{
				{"red": 0},
				{"green": 0},
				{"blue": 0},
			},
		}

		parts := strings.SplitN(gameString, ":", 2)
		game.ID, _ = strconv.Atoi(strings.TrimSpace(parts[0]))
		var cubes1 []Cube

		subsets := strings.Split(parts[1], ";")
		for _, subset := range subsets {
			cubes := strings.Split(subset, ",")
			for _, cube := range cubes {
				parts := strings.Fields(cube)
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]
				newCube := Cube{
					color: count,
				}
				cubes1 = append(cubes1, newCube)
			}
		}
		game.Cubes = cubes1

		games = append(games, game)
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

func sumSlice(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + sumSlice(numbers[1:])
}

func main() {
	input, _ := readLines("input.txt")
	games := parseGames(input)
	valid_ids := getValidGameIds(games)

	fmt.Println(sumSlice(valid_ids))
}
