package blackjack

import (
	"fmt"

	"github.com/Ephex2/gophercises/10.BlackJackGame/deck"
)

type Dealer struct {
	faceUp   []deck.Card
	faceDown deck.Card
}

// Returns the hand that the dealer is currently facing showing.
// Allows us to make the faceup []card property unexported, allowing us to keep players from editing it.
func (d *Dealer) FaceUp() Hand {
	publicHand := Hand{d.faceUp}
	return publicHand
}

func (d *Dealer) PrintFaceUp() {
	fmt.Println("Dealer is showing:")
	for _, c := range d.faceUp {
		fmt.Println(c.String())
	}
}

// Draws the dealers initial hand
func (d *Dealer) setup(twoCards []deck.Card) {
	d.faceDown = twoCards[0]
	d.faceUp = []deck.Card{twoCards[1]}
}

// Get the current value of the dealer's hand, from the game's perspective.
// Method cannot be exported otherwise Players could evaluate the hidden value of a dealer's hand.
func (d *Dealer) evaluateHand() int {
	cards := append(d.faceUp, d.faceDown)
	h := Hand{Cards: cards}

	return h.Evaluate()
}
