package main

import (
	"fmt"
	"strings"
)


// Receives player input (between 1 and 9) and returns x y coordinates for board
func position_map(player_input int) (int, int) {
	switch player_input {
	case 1:
		return 0, 0	
	case 2:
		return 0, 1
	case 3:
		return 0, 2
	case 4:
		return 1, 0
	case 5:
		return 1, 1
	case 6:
		return 1, 2
	case 7:
		return 2, 0
	case 8:
		return 2, 1
	case 9:
		return 2, 2
	
	// If it errors, return -1.
	default:
		return -1, -1
	}
}


// Asks player where they want to place marker
func player_play(board [3][3]string, player string) [3][3]string {
	var position int
	var valid_pos bool = false
	var x, y int // Play coordinates

	// Checks if the position is valid in the first place
	// There might be a better way to do this ?
	// Feel free to submit a PR
	for !valid_pos {
		fmt.Printf("%s's turn: ", player)
		fmt.Scanln(&position)
		
		// Only do anything if position is correct first
		if position < 1 || position > 9 {
			fmt.Println("Value must be between 1 and 9!")
		} else {
			// Determines intended position
			x, y = position_map(position)
			
			if board[x][y] != "_" {
				fmt.Println("Player has already played here!")
			} else {
				valid_pos = true
			}
		}
	}

	board[x][y] = player
	return board
}


// If player is X, player becomes O, and vice versa
func swap_player(player string) string {
	switch player {
	case "X":
		return "O"
	case "O":
		return "X"
	default:
		return "X"
	}
}


// Prints contents of board
func print_board(board [3][3]string) {
	fmt.Printf("\n")
	// Loops through board in order to print it
	for i := 0; i < 3; i++ {
		fmt.Printf("%s %s %s", board[i][0], board[i][1], board[i][2])
		fmt.Printf("     ") // Whitespace

		// Prints control label
		if i == 0 {
			fmt.Printf("1 2 3\n")
		} else if i == 1 {
			fmt.Printf("4 5 6\n")
		} else {
			fmt.Printf("7 8 9\n")
		}
	}
}


// Determine if we have a winner (returns true if game is over)
func check_game_state(board [3][3]string) bool {
	var player string
	var gameover bool
	var underscore_count int = 0

	// X X X   _ _ _   _ _ _
	// _ _ _   X X X   _ _ _
	// _ _ _   _ _ _   X X X
	for i := 0; i < 3; i++ {
		if board[i][0] != "_" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			player = board[i][0]
			gameover = true
		}
	}


	// X _ _   _ X _   _ _ X
	// X _ _   _ X _   _ _ X
	// X _ _   _ X _   _ _ X
	for i := 0; i < 3; i++ {
		if board[0][i] != "_" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			player = board[i][0]
			gameover = true
		}
	}


	// X _ _   _ _ X
	// _ X _   _ X _
	// _ _ X   X _ _
	
	if board[1][1] != "_" {
		if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
			gameover = true
			player = board[1][1]
		} else if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
			gameover = true
			player = board[1][1]
		}
	}


	// Determine if the game is a draw 
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != "_" {
				underscore_count++
			}
		}
	}


	// Returns correct value
	if underscore_count == 9 {
		fmt.Printf("Draw!")
		return true
	} else if gameover {
		fmt.Printf("%s is the winner!", player)
		return true
	} else {
		return false
	}
}


// Sets board to an empty state
func reset_board(board [3][3]string) [3][3]string {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "_"
		}
	}
	return board
}

func main() {
	var gameover bool = false
	var player string = "X"
	var replay string
	var board = [3][3]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}


	// Game ends when gameover is false
	for !gameover {
		print_board(board)
		board = player_play(board, player)
		player = swap_player(player)
		gameover = check_game_state(board)

		// If player wants to play again, reset the game
		if gameover {
			fmt.Printf("\nPlay again? (Y/n) ")
			fmt.Scanln(&replay)
			replay = strings.ToLower(replay)
			switch replay {
			case "y":
				gameover = false
				player = "X"
				board = reset_board(board) 
			}
		}
	}
}
