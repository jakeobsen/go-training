// Copyright 2022 Morten Jakobsen. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was rewritten from 01-tictactoe/main.go by postmannen
// Check out his GitHub here https://github.com/postmannen
// Thank you Bjørn!

package main

import (
	"fmt"
	"strconv"
)

type game struct {
	position      []string
	currentPlayer string
	nextMove      string
	winningPlayer string
	totalPlays    int
}

func newGame() *game {
	g := game{
		position:      []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		currentPlayer: "x",
		nextMove:      "",
		winningPlayer: "",
		totalPlays:    0,
	}

	return &g
}

// Main exported game function

func main() {
	g := newGame()

	// NB; split up the initialization, and the play code.
	g.play()
}

// Game logic below

func (g *game) play() error {
	g.printField()
	for {
		g.askPlayerForMove()
		g.tryPerformPlayerMove()
		g.printField()
		g.checkForWinner()

		// When a winner is found, break the loop and return winner
		if g.winningPlayer != "" {
			// NB: Added a break which removed the need for the else.
			if g.winningPlayer == "draw" {
				fmt.Println("It's a draw!")
				break
			}

			fmt.Println("Winner is " + g.winningPlayer + "!")
			break
		}
	}

	return nil
}

func (g *game) printField() {
	fmt.Println("┌───┬───┬───┐")
	fmt.Println("│ " + g.position[7] + " │ " + g.position[8] + " │ " + g.position[9] + " │")
	fmt.Println("├───┼───┼───┤")
	fmt.Println("│ " + g.position[4] + " │ " + g.position[5] + " │ " + g.position[6] + " │")
	fmt.Println("├───┼───┼───┤")
	fmt.Println("│ " + g.position[1] + " │ " + g.position[2] + " │ " + g.position[3] + " │")
	fmt.Println("└───┴───┴───┘")
}

func (g *game) askPlayerForMove() {
	fmt.Print("Player " + g.currentPlayer + ", make a move: ")
	_, err := fmt.Scanln(&g.nextMove)
	if err != nil {
		// NB: Added printing of the actual error.
		fmt.Printf("Error : %v\n", err)
		return
	}
}

func (g *game) tryPerformPlayerMove() {
	// The move selected by the player in askPlayerForMove() is validated
	// and the positions are updated if available.
	// The validation fails if the input is not parseable as an integer between 1 and 9

	// NB: Change to use strconv.Atoi.
	move, err := strconv.Atoi(g.nextMove)
	// NB: We should check the error first, and exit early to save processing.
	// We can then also completely remove the need for the else at the bottom.
	if err != nil {
		fmt.Printf("error: conversion from string to int failed: %v\n", err)
		return
	}

	// NB: Moved out the checking of legal characters so we can exit at once
	// and not do any more checks which is not needed.
	if move < 1 && move > 9 {
		fmt.Printf("bad move, enter a number between 1-9\n")
		return
	}

	// NB: Check if play is valid, exit if not.
	if g.position[move] != g.nextMove {
		fmt.Printf("not valid play, try again\n")
		return
	}

	// NB: We know that all checks have passed now, so we can update positions.
	//
	// This is a valid play, update positions to reflect which selection the player made
	g.position[move] = g.currentPlayer
	// And update total plays count for draw detection
	g.totalPlays++
	// Then switch to the other player
	g.switchPlayer()

}

func (g *game) switchPlayer() {
	// Toggle between player x and o
	if g.currentPlayer == "x" {
		g.currentPlayer = "o"
		return
	}

	// NB: The added return in the if clause above let's us remove the need for an else,  and
	// makes the code more readable since we know that the current player should be set to x.
	g.currentPlayer = "x"

}

func (g *game) checkForWinner() {
	// There is only 8 winning moves in tic-tac-toe, so I chose to just check for every combination
	// A draw happens when all 9 positions have been played (total_plays>=9), and no winner has been found

	// NB: Replaced all the else if's with a switch/case setup for readability.
	switch {
	case g.position[1] == g.position[2] && g.position[2] == g.position[3]:
		g.winningPlayer = g.position[1]
	case g.position[4] == g.position[5] && g.position[5] == g.position[6]:
		g.winningPlayer = g.position[4]
	case g.position[7] == g.position[8] && g.position[8] == g.position[9]:
		g.winningPlayer = g.position[7]
	case g.position[7] == g.position[4] && g.position[4] == g.position[1]:
		g.winningPlayer = g.position[7]
	case g.position[8] == g.position[5] && g.position[5] == g.position[2]:
		g.winningPlayer = g.position[8]
	case g.position[9] == g.position[6] && g.position[6] == g.position[3]:
		g.winningPlayer = g.position[9]
	case g.position[7] == g.position[5] && g.position[5] == g.position[3]:
		g.winningPlayer = g.position[7]
	case g.position[9] == g.position[5] && g.position[5] == g.position[1]:
		g.winningPlayer = g.position[9]
	case g.totalPlays >= 9:
		g.winningPlayer = "draw"
	default:
		g.winningPlayer = ""
	}
}
