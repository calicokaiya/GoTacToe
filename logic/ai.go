package logic

import (
	"fmt"
)

// THIS IS STILL A WIP!
// AI will determine if it is in a position to win with 1 placement, and take it accordingly.
// Else, AI will determine if the player is in a position to win next game and prevent it accordingly.
// Else, AI will determine which position would give it the biggest advantage, and take it (2 possible win conditions > 1 win condition) (OPTIONAL) (NOT IMPLEMENTED YET)
// In the case that I am too bad to code biggest advantage last-resort, AI should choose a random move.

// Checks if AI can win in one move and returns winning position
// This function contains hardcoded lines that may or may not make break things in the future.
func AiWinnable(gameboard [3][3]string, player string) int {
	var testboard [3][3]string
	testboard = gameboard
	var a, b int
	
	// This should cycle through all valid moves and attempt
	fmt.Printf("\n")
	for i := 1; i < 10; i++ {
		if CheckMoveValidity(gameboard, i) == 100 {
			a, b = PositionMap(i)
			testboard[a][b] = player
			if CheckGameState(testboard) == 200 {
				//fmt.Println(testboard)
				return i
			}
			testboard = gameboard
		}
	}
	return -100
}


// Checks if AI will lose in one move
// This function contains hardcoded lines that may or may not make break things in the future.
func AiLosable(gameboard [3][3] string, player string) int {
	var testboard [3][3]string
	testboard = gameboard
	var a, b int
	player = SwapPlayer(player)

	// This should cycle through all valid moves and attempt
	fmt.Printf("\n")
	for i := 1; i < 10; i++ {
		if CheckMoveValidity(gameboard, i) == 100 {
			a, b = PositionMap(i)
			testboard[a][b] = player
			if CheckGameState(testboard) == 100 {
				return i
			}
			testboard = gameboard
		}
	}
	return -100
}
