package playroom

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

//Table struct of internal Table type
type Table struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Gametype string           `json:"string"`
	Numseats int              `json:"numseats"`
	UserMap  map[string]*User `json:"users"`
	Seats    []Seat           `json:"seats"`
}

//Seat struct combines Player with a seat id and dealer boolean
type Seat struct {
	ID     string `json:"id"`
	Player Player `json:"player"`
	Dealer bool   `json:"dealer"`
}

//NewTable initializes a new internal Table
func NewTable(name, gametype string, numseats int) (Table, error) {

	id := uuid.New().String()

	seats := []Seat{}
	for i := 0; i < numseats; i++ {
		s := Seat{
			ID:     uuid.New().String(),
			Dealer: false,
		}
		seats = append(seats, s)
	}

	table := Table{
		ID:       id,
		Name:     name,
		Numseats: numseats,
		Gametype: gametype,
		UserMap:  map[string]*User{},
		Seats:    seats,
	}

	return table, nil //fix
}

//AddUser adds user to the Users slice in Table struct
func (t *Table) AddUser(u *User) {

	(*t).UserMap[(*u).ID] = u
}

//BroadcastTable broadcasts to all subscribed as user on table
func (t *Table) BroadcastTable() {
	for {
		time.Sleep(time.Second * 10)
		table, err := json.Marshal(t)

		if err != nil {
			fmt.Println("Error marshalling table for broadcast", err)
		}

		for _, v := range t.UserMap {
			if (*v).Wstableconn != nil {
				err := (*v).Wstableconn.WriteMessage(websocket.TextMessage, table)

				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

//AddPlayer adds a player struct constructed from User and table specific fields
func (t *Table) AddPlayer(u *User) {

	//id := uuid.New().String()

}

/* func (t *Table) RemovePlayer(playerId string) {
	//remove player from table.Players
}

func (t Table) GetPlayer(playerId string) Player {
	//return table.player
}
*/
func (t *Table) NewGame() {
	/* 	(*t).Game = makeGame() */
}

func (t *Table) Deal() {
	/* 	hands := (*t).Game.deal(len((*t).Players) + 1)

	   	for i := range (*t).Players {
	   		(*t).Players[i].Hand = hands[i]
	   	}

	   	(*t).Dealer.Hand = hands[len(hands)-1] */
}

func (t *Table) Play(playerindex int) {
	/* 	var result string
	   	fmt.Println("\nPlayer", t.Players[playerindex].Name, "is up. Good luck")
	   	(*t).Players[playerindex].Hand, result = (*t).Game.play((*t).Players[playerindex].Hand)
	   	fmt.Println("\nplayer", t.Players[playerindex].Name, result) */
}

func (t *Table) FinishGame() {

	/* 	scoreToBeat := 0
	   	var winner Player

	   	for _, player := range t.Players {
	   		total := cardsTotal(player.Hand)
	   		if total < 22 && total > scoreToBeat {
	   			scoreToBeat = total
	   			winner = player
	   		}
	   	}

	   	if winner.Name != "" {
	   		//play until > score to beat or busted
	   		if cardsTotal(t.Dealer.Hand) > scoreToBeat {
	   			winner = (*t).Dealer
	   		} else {
	   			for {
	   				(*t).Dealer.Hand = append((*t).Dealer.Hand, (*t).Game.Deck.dealCard())
	   				fmt.Println("\nDealer hand:", t.Dealer.Hand)
	   				if playerBusted(t.Dealer.Hand) {
	   					fmt.Println("\nDealer has busted!")
	   					break
	   				}
	   				if cardsTotal(t.Dealer.Hand) > scoreToBeat {
	   					winner = (*t).Dealer
	   					break
	   				}
	   			}
	   		}
	   	} else {
	   		winner = (*t).Dealer
	   	}

	   	//declare winner
	   	fmt.Println("\n\nThe winner is", winner.Name, "!!\n.....") */
}

func (t Table) PrintHands() {
	/* 	for _, player := range t.Players {
	   		player.printHand()
	   	}
	   	t.Dealer.printHand() */
}
