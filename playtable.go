package main

import (
	"fmt"

	"github.com/jar0m1r/blackjack"
)

func startTable() {

	t := blackjack.NewTable()

	numplayers := 1
	fmt.Println("How many players?")
	fmt.Scanln(&numplayers)

	for i := 0; i < numplayers; i++ {
		name := ""
		fmt.Println("Please enter user", i+1, "name")
		fmt.Scanln(&name)
		t.AddPlayer(name)
	}

	for {
		startGame(&t)
		newgame := ""
		fmt.Println("Do you want a new game y/n?")
		fmt.Scanln(&newgame)
		if newgame == "n" {
			break
		}
	}

}

func startGame(t *blackjack.Table) {
	(*t).NewGame()

	(*t).Deal()

	(*t).PrintHands() //one dealer hand blind (later)

	for playerindex := range (*t).Players {
		(*t).Play(playerindex)
	}

	(*t).PrintHands() //one dealer hand blind (later)

	//check if deck actually schrunk

	(*t).FinishGame()

}
