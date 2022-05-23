package blackjack

import (
	"errors"
	"fmt"
)

// PlayerInfo represents all interesting information about a player during a game of blackjack.
// Fields are not exported but have getter methods defined. Game orchestrates values, players cannot modify them.
type PlayerInfo struct {
	id         int    // ID used to distinguish between players. Set by game engine.
	points     int    // Proxy for money in our game. All players start wil a pool of points and can make bets from this pool. Gets set to 0 when Done() is called or bet <= 0 is made.
	hands      []Hand // Represents the hands of cards a player can have, and the bet on each hand. Abstrated due to splitting of the hands into subhands.
	player     Player // Reference to the player, in order to call methods.
	currentBet int    // Value of bet for current round.
	sideBet    int    // Value of side bet, will be 0 by default, may not be set.
}

// Public getter - player interface will be sent PlayerInfo type in SendInfo()
// This set of public getters allows the player's interface implementation to query the game state for info.
func (pi *PlayerInfo) GetId() int {
	return pi.id
}

// Public getter
func (pi *PlayerInfo) GetPoints() int {
	return pi.points
}

// Public getter
func (pi *PlayerInfo) GetHands() []Hand {
	return pi.hands
}

// Public getter
func (pi *PlayerInfo) GetCurrentBet() int {
	return pi.currentBet
}

// Public getter
func (pi *PlayerInfo) GetSideBet() int {
	return pi.sideBet
}

// Sends current version of struct to players.
func (pi *PlayerInfo) updateInfo() {
	pi.player.UpdateInfo(*pi)
}

// Asks player to place bets. Bets larger than current points are set to current points.
// If bet is <= 0, player is Done()
func (pi *PlayerInfo) offerBet() {
	betValue := pi.player.OfferBet()

	if betValue > pi.GetPoints() {
		betValue = pi.GetPoints()
	}

	pi.currentBet = betValue
	pi.updateInfo()
}

// Asks player to place sidebet. Bets larger than current points are set to current points.
// If bet is <= 0, player is Done()
func (pi *PlayerInfo) offerSideBet(d Dealer, h Hand) {
	pi.updateInfo()
	sidebet := pi.player.OfferSideBet(d, h)
	totalCurrentBet := pi.currentBet * len(pi.hands)

	if totalCurrentBet <= pi.points {
		pi.sideBet = pi.GetPoints()
	}

	pi.sideBet = sidebet
	pi.updateInfo()
}

// Offer player a split if both his cards are the same value, iff he has 2 cards in hand.
func (pi *PlayerInfo) offerSplit(d Dealer, h *Hand) bool {
	if pi.canSplit(h) {
		pi.updateInfo()
		return pi.player.OfferSplit(d, *h)
	}

	return false
}

// Determine if the player is allowed to split.
func (pi *PlayerInfo) canSplit(h *Hand) bool {
	if len(h.Cards) == 2 && h.Cards[0] == h.Cards[1] {
		if pi.currentBet*(len(pi.hands)+1) <= pi.points {
			return true
		}
	}

	return false
}

// Asks player if they would like to hit or not.
func (pi *PlayerInfo) offerHit(d Dealer, h Hand) bool {
	pi.updateInfo()
	hit := pi.player.OfferHit(d, h)
	return hit
}

// Updates a player's points. Returns an error if an invalid final value is the result ( < 0 )
func (pi *PlayerInfo) updatePoints(x int) error {
	if pi.points+x < 0 {
		errString := fmt.Sprintf("Unable to update points to a negative value. Current points: %v  -  Trying to add: %v", pi.points, x)
		return errors.New(errString)
	}

	pi.points = pi.points + x
	pi.updateInfo()
	return nil
}

func (pi *PlayerInfo) dealHand(g *Game) {
	pi.hands = []Hand{{Cards: g.drawCardsFromDeck(2)}}
	pi.updateInfo()
}

// Sends the related player a signal they are done playing, as well as the associated exit message.
func (pi *PlayerInfo) done(reason string) {
	pi.updateInfo()
	pi.player.Done(reason)
}
