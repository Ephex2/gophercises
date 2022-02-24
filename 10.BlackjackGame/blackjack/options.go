package blackjack

// Define options for the game of blackjack about to be played
type Options struct {
	Decks          int // The number of decks in the aggregate blackjack deck.
	StartingPoints int // The number of points each player should start with.
	Rounds         int // The number of rounds to be played. If blank, game continues until points are 0 or players quit.
}
