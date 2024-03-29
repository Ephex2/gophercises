# Deck of Cards - Gophercise 9

We were meant to create a package that would allow us to interact with a deck of cards, with the following requirements:

- An option to sort the cards with a user-defined comparison function. The sort package in the standard library can be used here, and expects a less(i, j int) bool function.
- A default comparison function that can be used with the sorting option.
- An option to shuffle the cards.
- An option to add an arbitrary number of jokers to the deck.
- An option to filter out specific cards. Many card games are played without 2s and 3s, while others might filter out other cards. We can provide a generic way to handle this as an option.
- An option to construct a single deck composed of multiple decks. This is used often enough in games like blackjack that having an option to build a deck of cards with say 3 standard decks can be useful.

Also, we must be able to generate new decks.

*Note*: We were not meant to create a Deck type but I could really wrap my head around how it should work better with such a type present, so I chose to implement one for my personal attempt at the task.

The package has a fair amount of tests as well.

## Documentation
I made an effort to provide the comments necessary for go doc to be helpful. The documentation can be output into your terminal with the following command:
```
go doc .\deck
```
OR
```
go doc ./deck
```


Automatically generated documentation:
```
package deck // import "."


CONSTANTS

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

TYPES

type Card struct {
	// Representation of a playing card in a typical 52-card deck.
	// Jokers have Suit.String() == "Joker" and a value of 0.
	// The first Joker is red and each subsequent Joker changes color (Ex, in a deck with 4 Jokers, there are 2 red Jokers and 2 black Jokers)
	Value int // Represents the card's value, ranges from 1 to 13. Ace is one.
	Suit  Suit
	Color CardColor
}

func (c *Card) String() string
    Output a string representation of the Card type, with the format: ("%v of %v
    with color %v", valueString, suitString, colorString) for regular cards
    ("Joker with color %v", colorString) for Jokers

type CardColor int
    CardColor.String() can be Red or Black.

func (i CardColor) String() string

type Deck []Card
    A deck defines methods applicable to a given []Card.

func NewDeck() Deck
    Decks should be sorted by suit and then by value. Generates a new 52 card
    deck already sorted.

func NewMultipleDeck(count int) (deck Deck)
    Returns a deck consisting of *count* 52-card decks. The decks are appended
    in order already sorted. However, the multiple card deck is not already
    sorted.

func (d *Deck) AddJoker(count int)
    Adds *count* cards with the Joker suit to the deck.

func (d *Deck) CustomSort(less func(i, j int) bool)
    Given a custom sorting function, return the deck sorted by this function.

func (d *Deck) Print()
    Prints the string representation of each ordered card in the deck on a new
    line.

func (d *Deck) RemoveCard(value int)
    Removes all cards of the input *value* from the deck. Some games may omit
    the use of 2s or 3s, for example.

func (d *Deck) Shuffle()
    Shuffle deck, using current time as seed. Based on:
    https://yourbasic.org/golang/shuffle-slice-array/

func (d *Deck) Sort()
    Sorts the deck based on their suit and then their value. Suits will be shown
    in the order they are written in the constant section. For reference (at
    time of writing): Spades, Diamonds, Clubs, Hearts, Joker

func (d *Deck) String() []string
    Returns a []string containing the ordered string representation of each card
    in the deck.

type Suit int
    Suit.String() one of: Spades, Diamonds, Clubs, Hearts, and Joker.

func (i Suit) String() string

```
