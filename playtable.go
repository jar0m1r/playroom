package main

import (
	"fmt"
)

func startTable() {

	t := NewTable()

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

func startGame(t *table) {
	(*t).NewGame()

	(*t).Deal()

	(*t).printHands() //one dealer hand blind (later)

	for playerindex := range (*t).players {
		(*t).Play(playerindex)
	}

	(*t).printHands() //one dealer hand blind (later)

	//check if deck actually schrunk

	(*t).FinishGame()

}
