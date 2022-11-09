package logic

import (
	"fmt"
)


// Loops through board lines and prints cells
func PrintBoard(board [3][3]string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s %s %s", board[i][0], board[i][1], board[i][2])
		fmt.Printf("      ") // Empty space

		// Prints controls label
		if i == 0 {
			fmt.Printf("1 2 3\n")
		} else if i == 1 {
			fmt.Printf("4 5 6\n")
		} else {
			fmt.Printf("7 8 9\n")
		}
	}
}


// Sets board to empty state
func ResetBoard() [3][3]string {
	var board = [3][3]string{{"_", "_", "_"},{"_","_","_"},{"_","_","_"}}
	return board
}


// Determine if we have a winner
// Returns 0 when game is still going,
// 50 when the game is a draw,
// 100 when X wins,
// 200 when O wins
func CheckGameState(board [3][3]string) int {
	var player string
	var underscore_count int = 0

	// Checks for horizontal wins
	for i := 0; i < 3; i++ {
		if board[i][0] != "_" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			player = board[i][0]
		}
	}

	
	// Checks for vertical wins
	for i := 0; i < 3; i++ {
		if board[0][i] != "_" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			player = board[0][i]
		}
	}


	// Checks for diagonal wins
	if board[1][1] != "_" {
		if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
			player = board[1][1]
		} else if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
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
	if player == "X" { 
		return 100
	} else if player == "O" {
		return 200
	} else if underscore_count == 9 {
		fmt.Printf("Draw!")
		return 50
	} else {
		return 0
	}
}


// If player is X, player becomes O, and vice versa
func SwapPlayer(player string) string {
	switch player {
	case "X":
		return "O"
	case "O":
		return "X"
	default:
		return "X"
	}
}


// Checks if move is valid
func CheckMoveValidity(board [3][3]string, position int) int {	
	var i, j int // Play coordinates
	// Only do anything if position is correct first
	if position < 1 || position > 9 {
		return -100
	} else {
		// Determines if intended position is free
		i, j = PositionMap(position)
		
		if board[i][j] != "_" {
			return -200
		} else {
			return 100
		}
	}
}


// Asks player where they want to place marker
func MarkBoard(board [3][3]string, player string, position int) [3][3]string {
	var i, j int // Cell coordinates
	i, j = PositionMap(position)
	board[i][j] = player
	return board
}


// Receives player input (between 1 and 9) and returns x y coordinates for board
func PositionMap(player_input int) (int, int) {
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
