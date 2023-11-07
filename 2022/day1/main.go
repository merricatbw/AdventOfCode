package main

import (
  helpers "github.com/merricatbw/AdventOfCode/helpers"
  "strings"
  "strconv"
  "fmt"
  "slices"
)

type Elf struct {
  foodItems []int
}

func (e Elf) countCalories() int {
  return helpers.Sum(e.foodItems)
}


func convertToFoodItems(food string) (foodItems []int) {
  for _, item := range strings.Split(food, "\n") {
    convertedFoodItem, error := strconv.Atoi(item)
    if error != nil {
      panic(error)
    }
    foodItems = append(foodItems, convertedFoodItem) 
  }
  return
}

func main() { 
  f := helpers.GetInput("input")
  rawFoodItems := strings.Split(f, "\n\n")
  var elves []Elf 
  var calorieCounts []int

  for _, items := range rawFoodItems {
    convertedFood := convertToFoodItems(items) 
    elf := Elf{convertedFood}
    elves = append(elves, elf)
  }

  for _, elf := range elves {
    calorieCounts = append(calorieCounts, elf.countCalories())
  }

  slices.Sort(calorieCounts)
  part1 := calorieCounts[len(calorieCounts) -1]
  part2 := helpers.Sum(calorieCounts[len(calorieCounts) -3 :])

  fmt.Printf("Part 1: %d\n", part1)
  fmt.Printf("Part 2: %d\n", part2)
}
