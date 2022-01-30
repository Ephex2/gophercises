package deck

import (
	"fmt"
	"strconv"
)

type Card struct {
	// Representation of a playing card in a typical 52-card deck.
	// Jokers have Suit.String() == "Joker" and a value of 0.
	// The first Joker is red and each subsequent Joker changes color (Ex, in a deck with 4 Jokers, there are 2 red Jokers and 2 black Jokers)
	Value int // Represents the card's value, ranges from 1 to 13. Ace is one.
	Suit  Suit
	Color CardColor
}

//go:generate stringer -type=Suit
// Suit.String() one of: Spades, Diamonds, Clubs, Hearts, and Joker.
type Suit int

// CardColor.String() can be Red or Black.
type CardColor int

const (
	// ...A brand new deck of cards is typically sorted by suit, so that the 13 cards in the deck are all spades,
	// the next 13 diamonds, the next 13 clubs, and the last 13 hearts...
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
	Joker
	Red CardColor = iota
	Black
)

var lastJokerColor bool

// Determine card color based on suit. Jokers get a color based on flipping the last Joker's color. First Joker is red.
func getCardColor(suit Suit) CardColor {
	var returnVal int
	if suit == 0 || suit == 2 {
		returnVal = 1 // Black
	}

	if suit == 1 || suit == 3 {
		returnVal = 0 // Red
	}

	if suit == 5 {
		// Joker case
		if !lastJokerColor {
			returnVal = 0 // Red
			lastJokerColor = true
		} else if lastJokerColor {
			returnVal = 1 // Black
			lastJokerColor = false
		}
	}

	if returnVal == 0 {
		return 0 // Red
	}

	return 1 // Black
}

// Output a string representation of the Card type, with the format:
// ("%v of %v with color %v", valueString, suitString, colorString) for regular cards
// ("Joker with color %v", colorString) for Jokers
func (c *Card) String() string {
	var valueString string
	suitString := c.Suit.String()

	switch c.Value {
	case 1:
		valueString = "Ace"
	case 11:
		valueString = "Jack"
	case 12:
		valueString = "Queen"
	case 13:
		valueString = "King"
	case 0:
		valueString = "Joker"
	default:
		valueString = strconv.Itoa(c.Value)
	}

	colorString := c.Color.String()
	if valueString == "Joker" {
		return fmt.Sprintf("Joker with color %v", colorString)
	}

	return fmt.Sprintf("%v of %v with color %v", valueString, suitString, colorString)
}
