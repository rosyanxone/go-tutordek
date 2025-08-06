package basic

import "fmt"

type Gamer struct {
	name  string
	games []string
}

func (gamer *Gamer) addGame(gameName string) {
	gamer.games = append(gamer.games, gameName)
}

func base() {
	gamer := Gamer{name: "john"}

	gamer.addGame("BlackWing")
	gamer.addGame("RDR 2")
	gamer.addGame("Black Myth Wukong")

	for _, game := range gamer.games {
		fmt.Println(game)
	}
}
