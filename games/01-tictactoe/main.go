// Copyright 2022 Morten Jakobsen. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
)

var (
	gamePosition  = []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	currentPlayer = "x"
	nextMove      = ""
	winningPlayer = ""
	totalPlays    = 0
)

func main() {
	printGameField()

	errorCount := 0
	for {
		if errorCount >= 3 {
			printGameField()
			errorCount = 0
		}

		askPlayerForMove()
		if tryPerformPlayerMove() {
			printGameField()
			checkForWinner()
		} else {
			errorCount++
		}

		// When a winner is found, break the loop and return winner
		if winningPlayer != "" {
			if winningPlayer == "draw" {
				fmt.Println("It's a draw!")
			} else {
				fmt.Println("Winner is " + winningPlayer + "!")
			}
			break
		}
	}
}

// Game logic below

func printGameField() {
	fmt.Println("┌───┬───┬───┐")
	fmt.Println("│ " + gamePosition[7] + " │ " + gamePosition[8] + " │ " + gamePosition[9] + " │")
	fmt.Println("├───┼───┼───┤")
	fmt.Println("│ " + gamePosition[4] + " │ " + gamePosition[5] + " │ " + gamePosition[6] + " │")
	fmt.Println("├───┼───┼───┤")
	fmt.Println("│ " + gamePosition[1] + " │ " + gamePosition[2] + " │ " + gamePosition[3] + " │")
	fmt.Println("└───┴───┴───┘")
}

func askPlayerForMove() {
	fmt.Print("Player " + currentPlayer + ", make a move: ")
	_, err := fmt.Scanln(&nextMove)
	if err != nil {
		return
	}
}

func tryPerformPlayerMove() bool {
	// The move selected by the player in askPlayerForMove() is validated
	// and the positions are updated if available.
	// The validation fails if the input is not parseable as an integer between 1 and 9

	move, err := strconv.ParseInt(nextMove, 0, 8)
	if err == nil && move > 0 && move < 10 && gamePosition[move] == nextMove {
		// This is a valid play, update positions to reflect which selection the player made
		gamePosition[move] = currentPlayer

		// And update total plays count for draw detection
		totalPlays++

		// Then switch to the other player
		switchPlayer()
		fmt.Println()
		return true
	} else {
		fmt.Println("Bad move, try again.\n")
		return false
	}
}

func switchPlayer() {
	// Toggle between player x and o
	if currentPlayer == "x" {
		currentPlayer = "o"
	} else {
		currentPlayer = "x"
	}
}

func checkForWinner() {
	// There is only 8 winning moves in tic-tac-toe, so I chose to just check for every combination
	// A draw happens when all 9 positions have been played (total_plays>=9), and no winner has been found

	if gamePosition[1] == gamePosition[2] && gamePosition[2] == gamePosition[3] {
		winningPlayer = gamePosition[1]
	} else if gamePosition[4] == gamePosition[5] && gamePosition[5] == gamePosition[6] {
		winningPlayer = gamePosition[4]
	} else if gamePosition[7] == gamePosition[8] && gamePosition[8] == gamePosition[9] {
		winningPlayer = gamePosition[7]
	} else if gamePosition[7] == gamePosition[4] && gamePosition[4] == gamePosition[1] {
		winningPlayer = gamePosition[7]
	} else if gamePosition[8] == gamePosition[5] && gamePosition[5] == gamePosition[2] {
		winningPlayer = gamePosition[8]
	} else if gamePosition[9] == gamePosition[6] && gamePosition[6] == gamePosition[3] {
		winningPlayer = gamePosition[9]
	} else if gamePosition[7] == gamePosition[5] && gamePosition[5] == gamePosition[3] {
		winningPlayer = gamePosition[7]
	} else if gamePosition[9] == gamePosition[5] && gamePosition[5] == gamePosition[1] {
		winningPlayer = gamePosition[9]
	} else if totalPlays >= 9 {
		winningPlayer = "draw"
	} else {
		winningPlayer = ""
	}
}
