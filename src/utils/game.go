package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// EmptyCel -> a simple empty cell
const EmptyCel = "_"

var counter int = 0

// InitGame -> Init the game resources
func InitGame() ([][]string, Player, Player) {
	fmt.Println("Welcome to the Tic-Tac-Toe game")
	board := CreateBoard()
	fmt.Println("The board was created...")
	player := CreatePlayer()
	fmt.Println("The player 1 was created...")
	player2 := CreatePlayer()
	fmt.Println("The player 2 was created...")
	return board, player, player2
}

func playerRound(board *[][]string, player *Player) {
	reader := bufio.NewReader(os.Stdin)
	var x, y int

	fmt.Printf("Player %s -> Enter a coordinate, eg (2 2):", player.Nickname)
	_, err := fmt.Fscan(reader, &x, &y)

	if err != nil {
		log.Fatal(err)
	}

	if (x > 2 || y > 2) || ((*board)[x][y] != EmptyCel) {
		fmt.Println("Invalid coord")
		for {
			fmt.Printf("Player %s -> Enter again a coordinate:", player.Nickname)
			_, err := fmt.Fscan(reader, &x, &y)

			if err != nil {
				log.Fatal(err)
			}

			if (x > 2 || y > 2) || ((*board)[x][y] != EmptyCel) {
				continue
			} else {
				(*board)[x][y] = player.Option
				break
			}
		}
	} else {
		(*board)[x][y] = player.Option
	}

	if verifyWinner(board) {
		DisplayBoard(board)
		fmt.Printf("The winner was %s\n", player.Nickname)
		os.Exit(0)
	}

	reader.ReadString('\n')
}

// GameLoop -> The core of the game
func GameLoop(board [][]string, player, player2 Player) {
	for {
		playerRound(&board, &player)
		counter++
		DisplayBoard(&board)

		if counter == 9 {
			if isFull(&board) {
				fmt.Println("The result was a draw")
				break
			}
		}

		playerRound(&board, &player2)
		counter++
		DisplayBoard(&board)
	}
}

func isFull(board *[][]string) bool {
	c := 0

	for _, row := range *board {
		for _, cell := range row {
			if cell == EmptyCel {
				c++
			}
		}
	}

	if c != 0 {
		return false
	}
	return true
}

func verifyWinner(board *[][]string) bool {
	for row := 0; row < len(*board); row++ {
		if (*board)[row][0] != EmptyCel && (*board)[row][0] == (*board)[row][1] && (*board)[row][0] == (*board)[row][2] {
			return true
		}

		for col := 0; col < 3; col++ {
			if (*board)[0][col] != EmptyCel && (*board)[0][col] == (*board)[1][col] && (*board)[0][col] == (*board)[2][col] {
				return true
			}

			if (*board)[1][1] != EmptyCel {
				if (*board)[1][1] == (*board)[0][0] && (*board)[1][1] == (*board)[2][2] {
					return true
				}

				if (*board)[1][1] == (*board)[0][2] && (*board)[1][1] == (*board)[2][0] {
					return true
				}
			}

		}
	}
	return false
}
