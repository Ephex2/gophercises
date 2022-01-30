package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// A deck defines methods applicable to a given []Card.
type Deck []Card

// Decks should be sorted by suit and then by value.
// Generates a new 52 card deck already sorted.
func NewDeck() Deck {
	deck := newDeckOfCards()
	return deck
}

// Returns a deck consisting of *count* 52-card decks.
// The decks are appended in order already sorted. However, the multiple card deck is not already sorted.
func NewMultipleDeck(count int) (deck Deck) {
	for i := 0; i < count; i++ {
		deck = append(deck, newDeckOfCards()...)
	}

	return
}

// Returns a sorted deck of 52 cards.
func newDeckOfCards() Deck {
	var cards []Card

	for _, val := range []Suit{Spades, Diamonds, Clubs, Hearts} {
		for i := 0; i < 13; i++ {
			color := getCardColor(val)
			cards = append(cards, Card{Value: i + 1, Suit: val, Color: color})
		}
	}

	return cards
}

// Returns a []string containing the ordered string representation of each card in the deck.
func (d *Deck) String() []string {
	var out []string

	for _, card := range *d {
		out = append(out, card.String())
	}

	return out
}

// Prints the string representation of each ordered card in the deck on a new line.
func (d *Deck) Print() {
	for _, card := range d.String() {
		fmt.Println(card)
	}
}

// Sorts the deck based on their suit and then their value.
// Suits will be shown in the order they are written in the constant section.
// For reference (at time of writing): Spades, Diamonds, Clubs, Hearts, Joker
func (d *Deck) Sort() {
	sort.Slice(*d, func(i, j int) bool {
		d2 := *d
		if d2[i].Suit < d2[j].Suit {
			return true
		} else if d2[i].Suit > d2[j].Suit {
			return false
		}

		return d2[i].Value < d2[j].Value
	})
}

// Given a custom sorting function, return the deck sorted by this function.
func (d *Deck) CustomSort(less func(i, j int) bool) {
	sort.Slice(*d, less)
}

// Shuffle deck, using current time as seed.
// Based on: https://yourbasic.org/golang/shuffle-slice-array/
func (d *Deck) Shuffle() {
	d2 := *d
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d2), func(i, j int) { d2[i], d2[j] = d2[j], d2[i] })
}

// Adds *count* cards with the Joker suit to the deck.
func (d *Deck) AddJoker(count int) {
	for i := 0; i < count; i++ {
		joker := Card{Value: 0, Suit: 5}
		joker.Color = getCardColor(joker.Suit)
		*d = append(*d, joker)
	}
}

// Removes all cards of the input *value* from the deck. Some games may omit the use of 2s or 3s, for example.
func (d *Deck) RemoveCard(value int) {
	d2 := *d
	for i, card := range d2 {
		if card.Value == value {
			// without len(d2) - 1, the last element gets added to the deck each time this if is entered.
			d2 = append(d2[:i], d2[i+1:]...)
		}
	}

	*d = d2
}
