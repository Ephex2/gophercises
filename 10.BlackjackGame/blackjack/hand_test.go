package blackjack_test

import (
	"testing"

	"github.com/Ephex2/gophercises/10.DeckOfCards/blackjack"
	"github.com/Ephex2/gophercises/10.DeckOfCards/deck"
)

func TestEvaluate(t *testing.T) {
	valueHandMap := make(map[int]blackjack.Hand)

	// Bust! Should return -1
	valueHandMap[-1] = blackjack.Hand{
		[]deck.Card{
			{Value: 13},
			{Value: 13},
			{Value: 13},
		},
	}

	// Aces Should be high if you do not bust
	valueHandMap[17] = blackjack.Hand{
		[]deck.Card{
			{Value: 1},
			{Value: 10},
			{Value: 6},
		},
	}

	// Aces Should be low if them being high causes a bust, blackjack should not bust.
	valueHandMap[21] = blackjack.Hand{
		[]deck.Card{
			{Value: 1},
			{Value: 10},
			{Value: 10},
		},
	}

	// Blackjack should not bust without an Ace.
	valueHandMap[21] = blackjack.Hand{
		[]deck.Card{
			{Value: 2},
			{Value: 9},
			{Value: 10},
		},
	}

	// Many Aces!
	valueHandMap[14] = blackjack.Hand{
		[]deck.Card{
			{Value: 1},
			{Value: 1},
			{Value: 1},
			{Value: 1},
		},
	}

	// Key of map is expected result. Value of map is a given hand.
	for k, v := range valueHandMap {
		testOutput := v.Evaluate()

		if testOutput != k {
			t.Errorf("When evaluating hands, expected %v, got %v", k, testOutput)
		}
	}
}
