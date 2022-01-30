package deck_test

import (
	"testing"

	"github.com/Ephex2/gophercises/9.DeckOfCards/deck"
)

func TestSuitString(t *testing.T) {
	var c deck.Card
	suitMap := make(map[int]string)
	suitMap[0] = "Spades"
	suitMap[1] = "Diamonds"
	suitMap[2] = "Clubs"
	suitMap[3] = "Hearts"
	suitMap[4] = "Joker"

	for i := 0; i < 5; i++ {
		c.Suit = deck.Suit(i)

		if c.Suit.String() != suitMap[i] {
			testError(
				t,
				"String()",
				"Suit",
				suitMap[i],
				c.Suit.String())
		}
	}
}
