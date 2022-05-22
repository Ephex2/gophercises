package deck_test

import (
	"testing"

	"github.com/Ephex2/gophercises/10.BlackJackGame/deck"
)

func TestCardColorString(t *testing.T) {
	var c deck.Card
	colorMap := make(map[int]string)
	colorMap[0] = "Red"
	colorMap[1] = "Black"

	for i := 0; i < len(colorMap); i++ {
		c.Color = deck.CardColor(i)

		if c.Color.String() != colorMap[i] {
			testError(
				t,
				"String()",
				"CardColor",
				colorMap[i],
				c.Suit.String())
		}
	}
}
