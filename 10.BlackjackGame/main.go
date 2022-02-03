package main

import (
	"github.com/Ephex2/gophercises/10.DeckOfCards/blackjack"
)

// This type should implement the blackjack.AI interface
type AI struct{}

func main() {
	//var ai AI
	// setup ai if you need to...

	opts := blackjack.Options{
		Decks: 3,
	}

	blackjack.New(opts)
	//winnings := game.Play(ai)
	//fmt.Println("Our AI won/lost:", winnings)
}
