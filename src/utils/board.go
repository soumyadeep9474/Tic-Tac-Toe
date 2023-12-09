package utils

import (
	"fmt"
	"strings"
)

// CreateBoard -> A function that creates the tic-tac-toe board
func CreateBoard() [][]string {
	board := [][]string{
		[]string{EmptyCel, EmptyCel, EmptyCel},
		[]string{EmptyCel, EmptyCel, EmptyCel},
		[]string{EmptyCel, EmptyCel, EmptyCel},
	}

	return board
}

// DisplayBoard -> A function that displays the game board
func DisplayBoard(board *[][]string) {
	for i := 0; i < len(*board); i++ {
		fmt.Println(strings.Join((*board)[i], " "))
	}

	fmt.Println()
}
