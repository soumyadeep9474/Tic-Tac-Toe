package main

import "utils"

func main() {
	board, player, player2 := utils.InitGame()
	utils.GameLoop(board, player, player2)
}
