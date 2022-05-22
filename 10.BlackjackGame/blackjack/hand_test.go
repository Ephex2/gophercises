package blackjack_test

import (
	"testing"

	"github.com/Ephex2/gophercises/10.BlackJackGame/blackjack"
	"github.com/Ephex2/gophercises/10.BlackJackGame/deck"
)

func TestEvaluate(t *testing.T) {
	// using a map with value as key doesn't allow us to reuse different hands for the same value.
	type valueHand struct {
		value int
		hand  blackjack.Hand
	}

	var valueHands = []valueHand{{
		value: -1,
		hand: blackjack.Hand{
			[]deck.Card{
				{Value: 13},
				{Value: 13},
				{Value: 13},
			},
		},
	}}

	valueHands = append(valueHands, valueHand{
		value: 17,
		hand: blackjack.Hand{
			[]deck.Card{
				{Value: 1},
				{Value: 10},
				{Value: 6},
			},
		},
	},
	)

	valueHands = append(valueHands, valueHand{
		value: 21,
		hand: blackjack.Hand{
			[]deck.Card{
				{Value: 1},
				{Value: 10},
				{Value: 10},
			},
		},
	},
	)

	valueHands = append(valueHands, valueHand{
		value: 14,
		hand: blackjack.Hand{
			[]deck.Card{
				{Value: 1},
				{Value: 1},
				{Value: 1},
				{Value: 1},
			},
		},
	},
	)

	valueHands = append(valueHands, valueHand{
		value: 21,
		hand: blackjack.Hand{
			[]deck.Card{
				{Value: 2},
				{Value: 9},
				{Value: 10},
			},
		},
	},
	)

	// Key of map is expected result. Value of map is a given hand.
	for _, valueHand := range valueHands {
		testOutput := valueHand.hand.Evaluate()

		if testOutput != valueHand.value {
			t.Errorf("When evaluating hands, expected %v, got %v", valueHand.value, testOutput)
		}
	}
}
