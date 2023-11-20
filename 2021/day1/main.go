package main

import (
	"fmt"
	"strconv"
	"strings"

	helpers "github.com/merricatbw/AdventOfCode/helpers"
)

func didIncrease(cur, prev int) bool {
	return cur > prev
}

func countIncreases(depths []int) int {
	var increases int
	for i, depth := range depths {
		currentDepth := depth
		if i > 0 {
			prevDepth := depths[i-1]
			if didIncrease(currentDepth, prevDepth) {
				increases++
			}
		}
	}
	return increases
}

func main() {
	data := helpers.GetInput("input")
	data = strings.Trim(data, "\n")

	var depths []int
	for _, depthString := range strings.Split(data, "\n") {
		depthInt, err := strconv.Atoi(depthString)
		if err != nil {
			panic(err)
		}
		depths = append(depths, depthInt)
	}

	var depthChunks [][]int
	for i, _ := range depths {
		if i+2 >= len(depths) {
			break
		}
		depthChunk := []int{depths[i], depths[i+1], depths[i+2]}
		depthChunks = append(depthChunks, depthChunk)
	}

	var depthChunksSummed []int
	for _, chunk := range depthChunks {
		depthChunksSummed = append(depthChunksSummed, helpers.Sum(chunk))
	}

	fmt.Printf("Part 1: %d\n", countIncreases(depths))
	fmt.Printf("Part 2: %d\n", countIncreases(depthChunksSummed))
}
