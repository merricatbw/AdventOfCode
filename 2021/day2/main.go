package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/merricatbw/AdventOfCode/helpers"
)

type Position struct {
	depth int
	horiz int
}

type Submarine struct {
	position Position
	aim      int
}

func (sub *Submarine) move(instruction string, units int) {
	switch strings.ToUpper(instruction) {
	case "UP":
		sub.position.depth -= units
	case "DOWN":
		sub.position.depth += units
	case "FORWARD":
		sub.position.horiz += units
	}
}

func (sub *Submarine) correctedMove(instruction string, units int) {
	switch strings.ToUpper(instruction) {
	case "UP":
		sub.aim -= units
	case "DOWN":
		sub.aim += units
	case "FORWARD":
		sub.position.horiz += units
		sub.position.depth += sub.aim * units
	}
}

func (sub *Submarine) reset() {
	sub.position.depth = 0
	sub.position.horiz = 0
	sub.aim = 0
}

func (sub Submarine) readOut() string {
	return fmt.Sprintf("Depth: %d\nHorizontal Position: %d\n", sub.position.depth, sub.position.horiz)
}

func main() {
	data := helpers.GetInput("input")
	data = strings.Trim(data, "\n")

	sub := Submarine{position: Position{depth: 0, horiz: 0}, aim: 0}

	for _, instruction := range strings.Split(data, "\n") {
		splitInstruction := strings.Split(instruction, " ")
		direction := splitInstruction[0]
		unit, err := strconv.Atoi(splitInstruction[1])
		if err != nil {
			panic(err)
		}
		sub.move(direction, unit)
	}

	fmt.Printf("Part 1: %d\n", (sub.position.depth * sub.position.horiz))
	sub.reset()

	for _, instruction := range strings.Split(data, "\n") {
		splitInstruction := strings.Split(instruction, " ")
		direction := splitInstruction[0]
		unit, err := strconv.Atoi(splitInstruction[1])
		if err != nil {
			panic(err)
		}
		sub.correctedMove(direction, unit)
	}
	fmt.Printf("Part 2: %d\n", (sub.position.depth * sub.position.horiz))
}
