package playroom

import "fmt"

//Player is the struct for Users actually playing on a table
type Player struct {
	ID    string `json:"id"`
	User  *User  `json:"user"`
	Name  string `json:"name"`
	Hand  []card `json:"hand"`
	Score int    `json:"score"`
}

func (p *Player) hit() {

}

func (p *Player) stand() {

}

func (p *Player) addToScore(amount int) int {
	(*p).Score += amount
	return (*p).Score
}

func (p Player) getScore() int {
	return p.Score
}

func (p Player) getHand() []card {
	return p.Hand
}

func (p Player) printHand() {
	fmt.Println("\nPlayer", p.Name, "hand:")
	for _, card := range p.Hand {
		fmt.Println(" ", card.name)
	}
}
