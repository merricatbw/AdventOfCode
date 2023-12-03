package main

import (
	"fmt"
	"strconv"
	"strings"

	helpers "github.com/merricatbw/AdventOfCode/helpers"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func main() {
	data := strings.Split(strings.Trim(helpers.GetInput("input"), "\n"), "\n")
	var totalValid int
	var totalPower int
	for _, game := range data {
		valid := true
		id, err := getGameId(game)
		if err != nil {
			panic(err)
		}
		game, _ = strings.CutPrefix(game, fmt.Sprintf("Game %d:", id))
		totalPower += getPower(game)
		for _, blocksPulled := range strings.Split(game, ";") {
			for _, pull := range strings.Split(blocksPulled, ",") {
				color, amount := getBlockInfo(pull)
				switch color {
				case "blue":
					if amount > maxBlue {
						valid = false
					}
				case "red":
					if amount > maxRed {
						valid = false
					}
				case "green":
					if amount > maxGreen {
						valid = false
					}
				}
			}
		}
		if valid {
			totalValid += id
		}
	}
	fmt.Printf("Part 1: %d\n", totalValid)
	fmt.Printf("Part 2: %d\n", totalPower)

}

func getPower(game string) (power int) {
	minCubes := findFewestPossibleCubes(game)
	power = minCubes["red"] * minCubes["blue"] * minCubes["green"]
	return
}

func findFewestPossibleCubes(game string) map[string]int {
	lowestPossible := make(map[string]int)
	lowestPossible["red"] = 0
	lowestPossible["green"] = 0
	lowestPossible["blue"] = 0
	for _, cubesSet := range strings.Split(game, ";") {
		for _, cube := range strings.Split(cubesSet, ",") {
			color, amount := getBlockInfo(cube)
			if lowestPossible[color] < amount {
				lowestPossible[color] = amount
			}
		}
	}
	return lowestPossible
}

func getBlockInfo(pull string) (color string, amount int) {
	split := strings.Split(strings.Trim(pull, " "), " ")
	color = split[1]
	amount, _ = strconv.Atoi(split[0])
	return
}

func getGameId(game string) (id int, err error) {
	game, _ = strings.CutPrefix(game, "Game ")
	game = strings.Split(game, ":")[0]
	id, err = strconv.Atoi(game)
	return
}
