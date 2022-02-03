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

func (d *Dealer) setup(deckPointer *deck.Deck) {
	d.faceDown = deckPointer.Draw(1)[0]
	d.faceUp = deckPointer.Draw(1)
}

// Get the current value of the dealer's hand, from the game's perspective.
// Method cannot be exported otherwise Players could evaluate the hidden value of a dealer's hand.
func (d *Dealer) evaluateHand() int {
	cards := append(d.faceUp, d.faceDown)
	h := Hand{Cards: cards}

	return h.Evaluate()
}
