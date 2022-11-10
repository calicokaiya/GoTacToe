package main

import (
	"fmt"
	"calicokaiya/tictactoe/logic"
	"strings"
	"math/rand"
	"time"
)


// Handles inputs and printing.
// Playertype 0 is bot,
// Playertype 1 is human.
func play(board [3][3]string, player string, playertype int) [3][3]string {
	var validpos int = -1
	var position int

	fmt.Printf("%s's turn: ", player)

	// 100 is a valid position, anything else is an error
	// If player is human, ask them for their input
	// If player is a bot, do everything automatically
	for validpos != 100 {
		if playertype == 1 {
			// Player logic
			fmt.Scanln(&position)
			validpos = logic.CheckMoveValidity(board, position)
			if validpos == -100 {
				fmt.Printf("Number should be between 1 and 9!")
			} else if validpos == -200 {
				fmt.Printf("This cell is already occupied!")
			} else {
				board = logic.MarkBoard(board, player, position)
			}
		} else {
			// Bot logic
		
			// See if AI can win with immediate move
			// Will try every single move and determine if AI would win
			
			position = logic.AiWinnable(board, player) // This should be 100 if AI can win
			if position != -100 {
				validpos = 100
			} else { 
				// If the AI can't win, it'll determine if it can lose
				position = logic.AiLosable(board, player)
				fmt.Printf("Set position to %d\n", position)
				if position != -100 {
					fmt.Printf("AI will lose if they don't play %d", position)
					validpos = 100
				} else { 
					// Set random seed
					fmt.Printf("Going into random")
					s1 := rand.NewSource(time.Now().UnixNano())
					r1 := rand.New(s1)

					position = r1.Intn(9) + 1
					if logic.CheckMoveValidity(board, position) == 100 {

						validpos = 100
					}
				}
			}
		}
	}
	fmt.Println("Trying to mark board")
	board = logic.MarkBoard(board, player, position)
	fmt.Printf("\n")
	return board
}


func main() {
	var gameover int = 0
	var player string = "X"
	var replay string
	var board = [3][3]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}



	// Game ends when gameover is different than 0
	for gameover == 0 {
		fmt.Printf("\n")
		logic.PrintBoard(board)
		if player == "X" {
			board = play(board, player, 1)
		} else {
			board = play(board, player, 0)
		}
		gameover = logic.CheckGameState(board)

		// If player wants to play again, reset the game
		if gameover != 0 {
			logic.PrintBoard(board)
			logic.PrintResults(gameover, player)
			
			// Get user input
			fmt.Printf("Play again? [Y/n] ")
			fmt.Scanln(&replay)
			replay = strings.ToLower(replay)
			switch replay {
			case "y":
				gameover = 0
				player = "O"
				board = logic.ResetBoard()
				
			}
		}

		// Once a player is done with their turn, let other player play
		player = logic.SwapPlayer(player)
	}
}
