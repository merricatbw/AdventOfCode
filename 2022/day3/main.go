package main

import (
  helpers "github.com/merricatbw/AdventOfCode/helpers"
  "fmt"
  "strings"
)

type Backpack struct {
  compartments map[int][]Item
}

func newBackpack(items []Item) Backpack {
  compartments := make(map[int][]Item)
  compartments[0] = items[0:len(items)/2]
  compartments[1] = items[len(items)/2:]
  return Backpack{compartments: compartments}
}

func (b Backpack)findDupe() (isDupe bool, dupe Item) { 
  for _, x := range b.compartments[0] {
    for _, y := range b.compartments[1] {
      if x == y {
        return true, x
      }
    }
  }
  return false, Item{} 
}

type Item struct {
  _type rune;
  priority int
}

func newItem(_type rune) Item {
  priority := int(_type)
  if priority >= 65 && priority <= 90 {
    priority -= 38
  } else if priority >= 97 && priority <= 122 {
    priority -= 96
  }

  return Item{_type: _type, priority: priority}
}

func main() {
  data := helpers.GetInput("input");
  data = strings.Trim(data, "\n")

  priorityTotal := 0
  var backpacks []Backpack

  for _, itemString := range strings.Split(data, "\n") { 
    var items []Item
    for _, item := range itemString {
      items = append(items, newItem(item))
    }
    backpacks = append(backpacks, newBackpack(items))
  }

  for _, backpack := range backpacks {
    isDupe, dupe := backpack.findDupe()
    if isDupe {
      priorityTotal += dupe.priority
    }
  }

  fmt.Printf("Part 1: %d\n", priorityTotal)
}
