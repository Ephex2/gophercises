package deck_test

import (
	"fmt"
	"testing"

	"github.com/Ephex2/gophercises/10.DeckOfCards/deck"
)

// Helper functions are in helpers_test.go.

func TestNewDeck(t *testing.T) {
	d := deck.NewDeck()
	testDeckIntegrity(t, 1, d, 0)

}

func TestNewMultipleDeck(t *testing.T) {
	// Test for a variety of deck sizes
	sizes := []int{4, 11, 100}
	for _, val := range sizes {
		d := deck.NewMultipleDeck(val)
		testDeckIntegrity(t, val, d, 0)
	}
}

func TestDeckSorting(t *testing.T) {
	d := deck.NewDeck()
	d.Shuffle()
	d.Sort()
	// First three cards should be Ace of Spades, Two of Spades, and Three of Spades.
	for i := 0; i < 4; i++ {
		if d[i].Value != i+1 && d[i].Suit != 0 {
			expectedString := fmt.Sprintf("%v for card in position %v in sorted deck is incorrect.", i+1, i)
			testError(t, "DeckSorting()", "Deck", expectedString, fmt.Sprint(d[i].Value))
		}
	}
}

func TestDeckCustomSortings(t *testing.T) {
	// Custom sort using a value first, suit second function
	// First four cards should have value 1 ( each be Aces ).
	d := deck.NewDeck()
	less := func(i, j int) bool {
		if d[i].Value < d[j].Value {
			return true
		} else if d[i].Value > d[j].Value {
			return false
		}

		return d[i].Suit < d[j].Suit
	}

	d.CustomSort(less)

	for i := 0; i < 4; i++ {
		if d[i].Value != 1 {
			// One of the first four cards is not an Ace, error:
			testError(t, "CustomSort()", "Deck", fmt.Sprint(1), fmt.Sprint(d[i].Value))
		}
	}
}

func TestShuffle(t *testing.T) {
	// Will make sure first three cards ARE NOT the Ace, Two, and Three of Spades after shuffling, three times in a row.
	// It is possible that this happens by chance but should be outlandishly rare.
	d := deck.NewDeck()
	d.Sort()

	var unshuffledHitCount int
	for i := 0; i < 3; i++ {
		d.Shuffle()
		var internalHitCount int

		for j := 0; j < 3; j++ {
			if d[i].Value == i && d[i].Suit == 0 {
				internalHitCount++
			}
		}

		if internalHitCount == 3 {
			unshuffledHitCount++
		}
	}

	if unshuffledHitCount == 3 {
		testError(t,
			"Shuffle()",
			"Deck",
			"To not have the first three cards be the same after shuffling three times.",
			"The first three cards were the same each shuffle, after shuffling three times.",
		)
	}
}

func TestAddJoker(t *testing.T) {
	// Test deck integrity after adding, 1, 5, and 8000 jokers.
	for _, val := range []int{1, 5, 8000} {
		d := deck.NewDeck()
		d.AddJoker(val)
		testDeckIntegrity(t, 1, d, val)
	}
}

func TestRemoveCard(t *testing.T) {
	// Remove all 7s and 6s, expect to hit no 7s or 6s.
	// Also expect deck sice to be 52-8 = 44
	d := deck.NewDeck()
	d.RemoveCard(6)
	d.RemoveCard(7)

	if len(d) != 44 {
		testError(t, "RemoveCard()", "Deck", "44 cards in deck", fmt.Sprint(len(d)))
	}

	var sixCount int
	var sevenCount int
	for _, card := range d {
		if card.Value == 6 {
			sixCount++
		}
		if card.Value == 7 {
			sevenCount++
		}
	}

	if sixCount != 0 {
		testError(t, "RemoveCard()", "Deck", "0 cards with value 6 in deck", fmt.Sprint(sixCount))
	}

	if sevenCount != 0 {
		testError(t, "RemoveCard()", "Deck", "0 cards with value 7 in deck", fmt.Sprint(sevenCount))
	}
}
