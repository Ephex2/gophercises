package blackjack_test

import (
	"fmt"

	"github.com/Ephex2/gophercises/10.DeckOfCards/blackjack"
)

type TestPlayer struct {
	id         int
	points     int
	hands      []blackjack.Hand
	currentBet int
	sideBet    int
}

func (tp *TestPlayer) UpdateInfo(pi blackjack.PlayerInfo) {
	tp.id = pi.GetId()
	tp.points = pi.GetPoints()
	tp.hands = pi.GetHands()
	tp.currentBet = pi.GetCurrentBet()
	tp.sideBet = pi.GetSideBet()
}

// This test player is all-in; offer all possible points
func (tp *TestPlayer) OfferBet() int {
	return tp.points
}

// This player does not double-down
func (tp *TestPlayer) OfferDoubleDown(d *blackjack.Dealer, h blackjack.Hand) bool {
	return false
}

// Take a hit if we don't have blackjack.
// This is dumb, but should allow us to test players losing more easily.
func (tp *TestPlayer) OfferHit(d *blackjack.Dealer, h blackjack.Hand) bool {
	return h.Evaluate() != 21
}

// Side bets are not worth it to this player.
func (tp *TestPlayer) OfferSideBet(d *blackjack.Dealer, h blackjack.Hand) int {
	return 0
}

// This player always splits if the card values are < 9.
// Splits are always on pairs; if evaluate of the hand < 18 we should split given the above condition.
func (tp *TestPlayer) OfferSplit(d *blackjack.Dealer, h blackjack.Hand) bool {
	return h.Evaluate() < 18
}

// Let our test player know that the game is over and also why.
func (tp *TestPlayer) Done(doneMsg string) {
	fmt.Printf("Message for player with id: %v -> Game is over -- message from game: %v\n", tp.id, doneMsg)
}
