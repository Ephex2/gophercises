package deck_test

import (
	"fmt"
	"testing"

	"github.com/Ephex2/gophercises/10.DeckOfCards/deck"
)

func testError(t *testing.T, method string, typ string, expected string, received string) {
	t.Errorf("Error with %v method on %v type. Received: %q but expected: %q", method, typ, received, expected)
}

// Each deck must have 52 cards, 4 of every value from 1 to 13. Jokers also = to the number of jokers specified.
func testDeckIntegrity(t *testing.T, numDecks int, d deck.Deck, numJokers int) {
	if numDecks < 1 {
		t.Errorf("Incorrect number of decks sent to testDeckIntegrity: %v", numDecks)
	}

	for i := 1; i < 14; i++ {
		// Standard deck test section.
		// Loop through deck to evaluate values, if 4 do not match, fail test.
		var testCount int
		var shouldCount = 4 * numDecks
		var shouldText = fmt.Sprintf("%v, for value: %v", shouldCount, i)
		for _, card := range d {
			if card.Value == i {
				testCount++
			}
		}

		if testCount != shouldCount {
			testError(t, "NewDeck()", "Deck", shouldText, fmt.Sprint(testCount))
		}
	}

	if numJokers != 0 {
		var jokerTestCount int
		for _, card := range d {
			if card.Value == 0 {
				jokerTestCount++
			}
		}

		if numJokers != jokerTestCount {
			testError(t, "NewDeck()", "Deck", fmt.Sprint(numJokers), fmt.Sprint(jokerTestCount))
		}
	}
}
