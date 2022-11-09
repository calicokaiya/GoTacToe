package logic

import (
	"fmt"
)

// THIS IS STILL A WIP!
// AI will determine if it is in a position to win with 1 placement, and take it accordingly.
// Else, AI will determine if the player is in a position to win next game and prevent it accordingly.
// Else, AI will determine which position would give it the biggest advantage, and take it (2 possible win conditions > 1 win condition) (OPTIONAL)
// In the case that I am too bad to code biggest advantage last-resort, AI should choose a random move.

// Checks if AI can win in one move and returns winning position
func Winnable(gameboard [3][3]string) {
	var testboard [3][3]string
	testboard = gameboard

	// THIS DOES NOT YET WORK - Redo this whole thing or smth idk
	for i := 1; i < 10; i++ {
		if CheckMoveValidity(gameboard, i) == 100 {
			var x, y int = PositionMap(i)
			testboard[x][y] = "O"
			if CheckGameState(testboard) == 200 {
				fmt.Println("AI can win")
			}
			testboard = gameboard
		}
	}
}

// Checks if AI will lose in one move
func Losable() {

}

