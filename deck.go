package playroom

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck []card

type card struct {
	name   string
	suit   string
	rank   string
	amount int
}

var amountMap = map[string]int{
	"Ace":   11, //fix later
	"Two":   2,
	"Three": 3,
	"Four":  4,
	"Five":  5,
	"Six":   6,
	"Seven": 7,
	"Eight": 8,
	"Nine":  9,
	"Ten":   10,
	"Jack":  10,
	"Queen": 10,
	"King":  10,
}

var suitUnicodeMap = map[string]string{
	"Spades":   "\u2660",
	"Hearts":   "\u2665",
	"Clubs":    "\u2663",
	"Diamonds": "\u2666",
}

func newDeck() Deck {

	d := Deck{}
	cardSuits := []string{"Clubs", "Hearts", "Diamonds", "Spades"}
	cardRanks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, s := range cardSuits {
		for _, r := range cardRanks {
			card := card{
				name:   r + " " + suitUnicodeMap[s],
				suit:   s,
				rank:   r,
				amount: amountMap[r],
			}
			d = append(d, card)
		}
	}

	return d
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card.name)
	}
}

func (d *Deck) deal(numhands int) [][]card {

	hands := [][]card{}

	for i := 0; i < numhands; i++ {
		hand := []card{(*d).dealCard(), (*d).dealCard()}
		hands = append(hands, hand)
	}

	return hands
}

func (d *Deck) dealCard() card {
	card := (*d)[0]
	(*d) = (*d)[1:]
	return card
}
