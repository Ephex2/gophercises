package blackjack

import "github.com/Ephex2/gophercises/10.DeckOfCards/deck"

type turn struct {
	deck     deck.Deck
	turnOver bool
}

func newTurn(g *Game) *turn {
	return &turn{}
}

func (t *turn) Hit()
