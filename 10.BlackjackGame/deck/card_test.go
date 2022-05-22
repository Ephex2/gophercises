package deck_test

import (
	"testing"

	"github.com/Ephex2/gophercises/10.BlackJackGame/deck"
)

// Helper functions are in helpers_test.go.

// First card should have Suit == 0 (spades) and Value == 1 (Ace).
// Last card should have Suit == 5 and Value == 0 (Joker)
func TestCardString(t *testing.T) {
	d := deck.NewDeck()
	d.AddJoker(1)
	d.Sort()

	firstCard := d.Cards[0]
	lastCard := d.Cards[len(d.Cards)-1]

	firstCardTestString := firstCard.String()
	firstCardShouldString := "Ace of Spades with color Black"

	lastCardTestString := lastCard.String()
	lastCardShouldString := "Joker with color Red"

	if firstCardTestString != firstCardShouldString {
		testError(t, "String()", "Card", firstCardShouldString, firstCardTestString)
	}

	if lastCardTestString != lastCardShouldString {
		testError(t, "String()", "Card", firstCardShouldString, firstCardTestString)
	}
}
