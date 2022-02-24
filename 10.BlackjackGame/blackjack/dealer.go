package blackjack

import "github.com/Ephex2/gophercises/10.DeckOfCards/deck"

type Dealer struct {
	faceUp   []deck.Card
	faceDown deck.Card
}

// Returns the int value of the card that the dealer is currently showing
// Allows us to make the faceup []card property unexported, allowing us to keep players from editing it.
func (d *Dealer) FaceUp() int {
	publicHand := Hand{d.faceUp}
	return publicHand.Evaluate()
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
