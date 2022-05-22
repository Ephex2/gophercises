package blackjack

import (
	"github.com/Ephex2/gophercises/10.BlackJackGame/deck"
)

// A hand held by a player.
type Hand struct {
	Cards []deck.Card // The set of cards in a given hand, at a given moment in the game.
}

// Returns the  value of a given hand, evaluating aces as either 1 or 11.
// Busting returns a negative value of -1.
// Always returns the highest possible legal value of a hand (so, A hand of Ace, Four, and Four would return 19).
func (h *Hand) Evaluate() int {
	var sum int
	var aceCounter int

	for _, card := range h.Cards {
		// Your hand will always bust with >1 Ace, stop caring about Aces after the first one.
		if card.Value == 1 && aceCounter < 1 {
			sum += 11
			aceCounter++
			continue
		}

		switch card.Value {
		// Next three cases are giving the face cards their appropriate 10 value.
		case 11, 12, 13:
			sum += 10
		default:
			sum += card.Value
		}
	}

	if sum > 21 && aceCounter == 1 {
		// Build alternate score, in case our first Ace at 11 busts us. 11 - 10 = 1 ; current - 10 = alternate hand score.
		alternate := sum - 10
		if alternate <= 21 {
			return alternate
		}
	} else if sum <= 21 {
		// Sum always higher than alternate if not bust.
		return sum
	}

	// Either alternate > 21 or aceCounter is 0 and sum is <> 21
	return -1
}
