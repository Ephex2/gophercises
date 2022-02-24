package main

import (
	"github.com/Ephex2/gophercises/10.DeckOfCards/blackjack"
	"github.com/Ephex2/gophercises/10.DeckOfCards/blackjackai"
)

func main() {
	opts := blackjack.Options{
		Decks:  10,
		Rounds: 10,
	}

	game := blackjack.New(opts)
	ai := &blackjackai.AiPlayer{}
	game.Play([]blackjack.Player{ai})
}
