package blackjack

import "github.com/Ephex2/gophercises/10.DeckOfCards/deck"

// A hand held by a player. Each hand has an individual bet associated with it to support splitting.
type Hand struct {
	Cards []deck.Card // The set of cards in a given hand, at a given moment in the game.
}

// Busting has a negative value (-1).
// Return value of hand, evaluating aces as either 1 or 11.
// Always returns highest possible legal value (so, A hand Ace - Four - Four would return 19).
func (h *Hand) Evaluate() int {
	var sum int
	var sum2 int

	for i, card := range h.Cards {
		switch card.Value {
		case 1:
			newHand := []Hand{{h.Cards}}
			newHand[0].Cards[i].Value = 0
			sum += 1
			sum2 += newHand[0].Evaluate()
		case 0:
			sum += 11
			sum2 += 11
		case 11:
			sum += 10
			sum2 += 10
		case 12:
			sum += 10
			sum2 += 10
		case 13:
			sum += 10
			sum2 += 10
		default:
			sum += card.Value
			sum2 += card.Value
		}
	}

	if sum > 21 && sum2 > 21 {
		return -1
	} else if sum >= sum2 {
		return sum
	} else {
		return sum2
	}
}
