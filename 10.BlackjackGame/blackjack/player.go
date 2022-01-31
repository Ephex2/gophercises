package blackjack

import "github.com/Ephex2/gophercises/10.DeckOfCards/deck"

// PlayerInfo represents all interesting information about a player during a game of blackjack.
// Derived Player types implementing the Player interface should be of type PlayerInfo.
// Fields are not exported but have getter methods defined. Game orchestrates values.
type PlayerInfo struct {
	id         int    // ID used to distinguish between players. Set by game engine.
	points     int    // Proxy for money in our game. All players start wil a pool of points and can make bets from this pool. Gets set to 0 when Done() is called or bet <= 0 is made.
	hands      []Hand // Represents the hands of cards a player can have, and the bet on each hand. Abstrated due to splitting of the hands into subhands.
	player     Player // Reference to the player, in order to call methods.
	currentBet int    // Value of bet for current round.
	sideBet    int    // Value of side bet, will be 0 by default, may not be set.
}

// Sends current version of struct to players.
func (pi *PlayerInfo) sendInfo() {
	pi.player.sendInfo(*pi)
}

// Asks player to place bets. Bets larger than current points are set to current points.
// If bet is <= 0, player is Done()
func (pi *PlayerInfo) PlaceBet() {
	betValue := pi.player.PlaceBet()

	if betValue > pi.GetPoints() {
		betValue = pi.GetPoints()
	}

	if betValue <= 0 {
		// player is quitting the game.
		pi.points = 0
		pi.Done()
	}

	pi.currentBet = betValue
}

func (pi *PlayerInfo) UpdatePoints(x int) {
	pi.points = pi.points + x
	pi.player.UpdatePoints(x)
}

func (pi *PlayerInfo) Done() {
	pi.points = 0
	pi.player.Done()
}

// A hand held by a player. Each hand has an individual bet associated with it to support splitting.
type Hand struct {
	Cards []deck.Card // The set of cards in a given hand, at a given moment in the game.
}

// The player interface allows the game orchestrator to interact with the different types of players in a game.
type Player interface {
	sendInfo(PlayerInfo)          // Set a player's info.
	PlaceBet() int                // Returns a player's starting bet for a round.
	OfferDoubleDown(*Dealer) bool // Allows a player to choose whether or not to double-down
	OfferHit(*Dealer) bool        // Offers a player to hit the deck for a new card in the current hand.
	OfferSideBet(*Dealer) int     // Allows players to determine if they want to take a sideBet, and if they do, for how much (int return). A negative value does nothing.
	OfferSplit(*Dealer) bool      // Offers a split to the player. If return is true, will update the player to have multiple hands, each with the one card in hand.
	UpdatePoints(int) error       // After a round ends, updates a player's Points according to the bets placed on their hands.
	Done()                        // After a game ends or a player is eliminated, notifies the player that the game is over. The input string is informational and refers to why the game is over for this player.
}

func (pi *PlayerInfo) GetId() int {
	return pi.id
}

func (pi *PlayerInfo) GetPoints() int {
	return pi.points
}

func (pi *PlayerInfo) GetHands() []Hand {
	return pi.hands
}

func (pi *PlayerInfo) GetCurrentBet() int {
	return pi.currentBet
}

func (pi *PlayerInfo) GetSideBet() int {
	return pi.sideBet
}

// Busting has a negative value.
// Return value of hand, evaluating aces as either 1 or 11.
// Always returns highest possible legal value.
// Game will take care of special blackjack opening rules.
func (h *Hand) Evaluate() int {
	var sum int
	var sum2 int

	for i, card := range h.Cards {
		switch card.Value {
		case 1:
			newHand := []Hand{{h.Cards}}
			newHand[0].Cards[i].Value = 0
			sum += 1
			sum2 += newHand[0].Evaluate()
		case 0:
			sum += 11
			sum2 += 11
		case 11:
			sum += 10
			sum2 += 10
		case 12:
			sum += 10
			sum2 += 10
		case 13:
			sum += 10
			sum2 += 10
		default:
			sum += card.Value
			sum2 += card.Value
		}
	}

	if sum > 21 && sum2 > 21 {
		return -1
	} else if sum >= sum2 {
		return sum
	} else {
		return sum2
	}
}
