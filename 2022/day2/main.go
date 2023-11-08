/*
  This one kind of got away from me... -merricat
*/
package main

import (
  helpers "github.com/merricatbw/AdventOfCode/helpers"
  "fmt"
  "strings"
)

const (
  rockString    = "Rock"
  paperString   = "Paper"
  scissorString = "Scissors"

  winString  = "Win"
  loseString = "Lose"
  drawString = "Draw"

)

func playRound(hands string) (playerScore, opponentScore int) {
  playerScore = 0
  opponentScore = 0

  splitHands   := strings.Split(hands, " ")
  playerHand   := decryptHand(splitHands[1])
  opponentHand := decryptHand(splitHands[0])


  switch playerHand {
  case rockString:
    switch opponentHand {
      case rockString:
        playerScore = calculateScore(playerHand, drawString)
        opponentScore = calculateScore(opponentHand, drawString)
      case paperString:
        playerScore = calculateScore(playerHand, loseString)
        opponentScore = calculateScore(opponentHand, winString)
      case scissorString:
        playerScore = calculateScore(playerHand, winString)
        opponentScore = calculateScore(opponentHand, loseString)
    }
  case paperString:
    switch opponentHand {
      case paperString:
        playerScore = calculateScore(playerHand, drawString)
        opponentScore = calculateScore(opponentHand, drawString)
      case scissorString:
        playerScore = calculateScore(playerHand, loseString)
        opponentScore = calculateScore(opponentHand, winString)
      case rockString:
        playerScore = calculateScore(playerHand, winString)
        opponentScore = calculateScore(opponentHand, loseString)
    }
  case scissorString:
    switch opponentHand {
      case scissorString:
        playerScore = calculateScore(playerHand, drawString)
        opponentScore = calculateScore(opponentHand, drawString)
      case rockString:
        playerScore = calculateScore(playerHand, loseString)
        opponentScore = calculateScore(opponentHand, winString)
      case paperString:
        playerScore = calculateScore(playerHand, winString)
        opponentScore = calculateScore(opponentHand, loseString)
    }
  }

  return
}

func calculateScore(hand string, outcome string) (score int) {
  handScore := make(map[string]int)
  handScore[rockString]    = 1
  handScore[paperString]   = 2
  handScore[scissorString] = 3

  outcomeScore := make(map[string]int)
  outcomeScore[winString]  = 6
  outcomeScore[drawString] = 3
  outcomeScore[loseString] = 0

  score = handScore[hand] + outcomeScore[outcome]
  return
}

func calculateCorrectHands(hands string) string {
  split := strings.Split(hands, " ")
  opponentHand := decryptHand(split[0])
  outcome := decryptOutcome(split[1])
  var playerHand string

  switch outcome {
    case winString:
      switch opponentHand {
      case rockString:
          playerHand = encryptHand(paperString)
      case paperString:
          playerHand = encryptHand(scissorString)
      case scissorString:
          playerHand = encryptHand(rockString)
      }
    case drawString:
      playerHand = encryptHand(opponentHand)
    case loseString:
      switch opponentHand {
      case rockString:
        playerHand = encryptHand(scissorString)
      case paperString:
        playerHand = encryptHand(rockString)
      case scissorString:
        playerHand = encryptHand(paperString)
      }
  }

  return fmt.Sprintf("%s %s", encryptHand(opponentHand), playerHand)

}

func encryptHand(decryptedHand string) (encrypted string) {
  switch decryptedHand {
  case rockString:
    encrypted = "A"
  case paperString:
    encrypted = "B"
  case scissorString:
    encrypted = "C"
  }
  return
}

func decryptOutcome(outcome string) (decrypted string) {
  switch outcome {
  case "X":
    decrypted = loseString
  case "Y":
    decrypted = drawString
  case "Z":
    decrypted =  winString
  }

  return
}

func decryptHand(hand string) (decrypted string) {
  switch hand{
  case "A", "X":
    decrypted = rockString
  case "B", "Y":
    decrypted = paperString
  case "C", "Z":
    decrypted = scissorString
  }
  return
}

func main() { 
  data := helpers.GetInput("input")
  totalPlayerScore1 := 0
  totalPlayerScore2 := 0
  for _, round := range strings.Split(data, "\n") {
    playerScore, _ := playRound(round)
    totalPlayerScore1 += playerScore
  }

  for _, round := range strings.Split(data, "\n") {
    playerScore, _ := playRound(calculateCorrectHands(round))
    totalPlayerScore2 += playerScore
  }

  fmt.Printf("Part 1: %d\n", totalPlayerScore1)
  fmt.Printf("Part 2: %d\n", totalPlayerScore2)
}
