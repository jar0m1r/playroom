package blackjack

import "fmt"

type Game struct {
	Deck Deck
}

func makeGame() Game {

	g := Game{
		Deck: newDeck(),
	}

	g.Deck.shuffle()

	return g
}

func (g *Game) deal(numhands int) [][]card {
	return (*g).Deck.deal(numhands)
}

func (g *Game) play(hand []card) ([]card, string) {
	for {
		if playerStands() {
			return hand, "stands"
		}

		card := (*g).Deck.dealCard()
		hand = append(hand, card)
		fmt.Println("newhand:", hand)

		if playerBusted(hand) {
			return hand, "busted"
		}
	}
}

func playerStands() bool {
	fmt.Println("Do you want to hit(h) or stand(s)?")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "h" || choice == "hit" {
		return false
	}
	return true
}

func playerBusted(cards []card) bool {
	return cardsTotal(cards) > 21
}

func cardsTotal(cards []card) int {
	total := 0
	for _, card := range cards {
		total += amountMap[card.rank]
	}

	return total
}
