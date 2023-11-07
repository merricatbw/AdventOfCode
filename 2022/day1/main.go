package main

import (
  helpers "github.com/merricatbw/AdventOfCode/helpers"
  "strings"
  "strconv"
  "fmt"
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

  for _, items := range rawFoodItems {
    convertedFood := convertToFoodItems(items) 
    elf := Elf{convertedFood}
    elves = append(elves, elf)
  }

  mostCalories := 0
  for _, elf := range elves {
    calories := elf.countCalories()
    if calories > mostCalories {
      mostCalories = calories
    }
  }

  fmt.Printf("Part 1: %d\n", mostCalories)

}
