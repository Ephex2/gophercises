package blackjack

import (
	"errors"
	"fmt"
)

// The player interface allows the game type to interact with the different types of players in a game.
type Player interface {
	SendInfo(PlayerInfo)          // Set a player's info.
	OfferBet() int                // Returns a player's starting bet for a round.
	OfferDoubleDown(*Dealer) bool // Allows a player to choose whether or not to double-down
	OfferHit(*Dealer) bool        // Offers a player to hit the deck for a new card in the current hand.
	OfferSideBet(*Dealer) int     // Allows players to determine if they want to take a sideBet, and if they do, for how much (int return). A negative value does nothing.
	OfferSplit(*Dealer) bool      // Offers a split to the player. If return is true, will update the player to have multiple hands, each with the one card in hand.
	UpdatePoints(int) error       // After a round ends, updates a player's Points according to the bets placed on their hands.
	Done(string)                  // After a game ends or a player is eliminated, notifies the player that the game is over. The input string is informational and refers to why the game is over for this player.
}

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
func (pi *PlayerInfo) SendInfo() {
	pi.player.SendInfo(*pi)
}

// Asks player to place bets. Bets larger than current points are set to current points.
// If bet is <= 0, player is Done()
func (pi *PlayerInfo) OfferBet() {
	betValue := pi.player.OfferBet()

	if betValue > pi.GetPoints() {
		betValue = pi.GetPoints()
	}

	if betValue < 1 {
		pi.Done("Betting 0 or less indicates that you wish to quit the table.")
	}

	pi.currentBet = betValue
	pi.SendInfo()
}

// Asks player to place sidebet. Bets larger than current points are set to current points.
// If bet is <= 0, player is Done()
func (pi *PlayerInfo) OfferSideBet(d *Dealer) {
	sidebet := pi.player.OfferSideBet(d)
	totalCurrentBet := pi.currentBet * len(pi.hands)

	if totalCurrentBet <= pi.points {
		pi.sideBet = pi.GetPoints()
	}

	if sidebet < 1 {
		pi.Done("Betting 0 or less indicates that you wish to quit the table.")
	}

	pi.sideBet = sidebet
	pi.SendInfo()
}

// Updates a player's points. Returns an error if an invalid final value is the result ( < 0 )
func (pi *PlayerInfo) UpdatePoints(x int) error {
	if pi.points+x < 0 {
		errString := fmt.Sprintf("Unable to update points to a negative value. Current points: %v  -  Trying to add: %v", pi.points, x)
		return errors.New(errString)
	}

	pi.points = pi.points + x

	pi.player.UpdatePoints(x)
	pi.SendInfo()
	return nil
}

// Sends the related player a signal they are done playing, as well as the associated exit message.
func (pi *PlayerInfo) Done(reason string) {
	pi.points = 0
	pi.player.Done(reason)
}
