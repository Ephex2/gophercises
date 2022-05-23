package deck

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// A deck defines methods applicable to a given []Card.
type Deck struct {
	Cards []Card
}

// Decks should be sorted by suit and then by value.
// Generates a new 52 card deck already sorted.
func NewDeck() Deck {
	cards := newDeckOfCards()
	return Deck{Cards: cards}
}

// Returns a deck consisting of *count* 52-card decks.
// The decks are appended in order already sorted. However, the multiple card deck is not already sorted.
func NewMultipleDeck(count int) (deck Deck) {
	for i := 0; i < count; i++ {
		deck.Cards = append(deck.Cards, newDeckOfCards()...)
	}

	return
}

// Returns a sorted deck of 52 cards.
func newDeckOfCards() []Card {
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

	for _, card := range d.Cards {
		out = append(out, card.String())
	}

	return out
}

// Draws n cards from the deck, removing it.
// If attempting to draw more cards than are in the deck, draw all remaining cards.
func (d *Deck) Draw(count int) (cards []Card, err error) {
	if count == 0 {
		return nil, nil
	}

	if count > len(d.Cards) {
		errMsg := fmt.Sprintf("Overdraw: Attempted to draw %v cards from deck of size %v", count, len(d.Cards))
		err = errors.New(errMsg)
		return nil, err
	}

	cards = d.Cards[0:count]
	d.Cards = d.Cards[count:]

	return cards, nil
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
	sort.Slice(d.Cards, func(i, j int) bool {
		if d.Cards[i].Suit < d.Cards[j].Suit {
			return true
		} else if d.Cards[i].Suit > d.Cards[j].Suit {
			return false
		}

		return d.Cards[i].Value < d.Cards[j].Value
	})
}

// Given a custom sorting function, return the deck sorted by this function.
func (d *Deck) CustomSort(less func(i, j int) bool) {
	sort.Slice(d.Cards, less)
}

// Shuffle deck, using current time squared as seed.
// Based on: https://yourbasic.org/golang/shuffle-slice-array/
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano() * time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

// Adds *count* cards with the Joker suit to the deck.
func (d *Deck) AddJoker(count int) {
	for i := 0; i < count; i++ {
		joker := Card{Value: 0, Suit: 5}
		joker.Color = getCardColor(joker.Suit)
		d.Cards = append(d.Cards, joker)
	}
}

// Removes all cards of the input *value* from the deck. Some games may omit the use of 2s or 3s, for example.
func (d *Deck) RemoveCard(value int) {
	for i, card := range d.Cards {
		if card.Value == value {
			// without len(d2) - 1, the last element gets added to the deck each time this if is entered.
			d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
		}
	}
}
