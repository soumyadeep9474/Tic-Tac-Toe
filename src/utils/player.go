package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Player -> Struct that represents a player
type Player struct {
	Nickname, Option string
}

// CreatePlayer -> A function that creates a single player
func CreatePlayer() (Playerr Player) {
	var nickname, option string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your nickname and option, eg (Foo X): ")
	_, err := fmt.Fscan(reader, &nickname, &option)

	if err != nil {
		log.Fatal(err)
	}

	reader.ReadString('\n')

	return Player{nickname, option}
}
