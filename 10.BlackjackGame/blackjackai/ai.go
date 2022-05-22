package blackjackai

import (
	"fmt"

	"github.com/Ephex2/gophercises/10.BlackJackGame/blackjack"
)

type AiPlayer struct {
	id         int
	points     int
	hands      []blackjack.Hand
	currentBet int
	sideBet    int
}

// Uses the PlayerInfos's Get*() methods to set all the properties of the AiPlayer struct.
// Prints the value of each of property of the AiPlayer object.
func (tp *AiPlayer) UpdateInfo(pi blackjack.PlayerInfo) {
	tp.id = pi.GetId()
	tp.points = pi.GetPoints()
	tp.hands = pi.GetHands()
	tp.currentBet = pi.GetCurrentBet()
	tp.sideBet = pi.GetSideBet()
}

// This ai offers 50 points, or all-in if he has less than 50 points.
// If the ai makes more than 200, they quit
// Optimization might be possible here?
func (tp *AiPlayer) OfferBet() int {
	if tp.points >= 200 {
		return 0
	} else if tp.points < 50 {
		return tp.points
	} else {
		return 50
	}
}

// Following this article for strategy: https://blog.betway.com/casino/blackjack-strategy-101-how-do-you-double-down-in-blackjack/
// Solution is a branching if-tree, seems inelegant. May consider using a decision-matrix instead or something.
func (tp *AiPlayer) OfferDoubleDown(d blackjack.Dealer, h blackjack.Hand) bool {
	var handContainsAce bool
	handValue := h.Evaluate()
	dealerFaceUp := d.FaceUp()

	// IF the dealer is showing an Ace, forget about it
	if d.FaceUp() == 1 {
		return false
	}

	if h.Cards[0].Value == 1 || h.Cards[1].Value == 1 {
		handContainsAce = true
	}

	if handContainsAce {
		// Double-down hands = A2-A7, depending on dealer's face-up card.
		if handValue >= 13 && handValue <= 15 {
			// Double-down on 5 or 6
			return dealerFaceUp >= 5 && dealerFaceUp <= 6
		} else if handValue == 16 {
			// Double-down on 4,5, or 6
			return dealerFaceUp >= 4 && dealerFaceUp <= 6
		} else if handValue == 17 || handValue == 18 {
			// Double-down on value of 3-6
			return dealerFaceUp >= 3 && dealerFaceUp <= 6
		} else {
			// Other values, return false
			return false
		}
	}

	// No ace in hand, only double-down for hand values in 8-11.
	if handValue > 7 {
		// Always double down on value of 11
		if handValue == 11 {
			return true
		} else if handValue == 8 {
			// Double-down on 5 or 6
			return dealerFaceUp >= 5 && dealerFaceUp <= 6
		} else if handValue == 9 {
			// Double-down on value of 2-6
			return dealerFaceUp >= 2 && dealerFaceUp <= 6
		} else if handValue == 10 {
			// Double-down on value of 2-9
			return dealerFaceUp >= 2 && dealerFaceUp <= 9
		} else {
			return false
		}
	} else {
		return false
	}
}

/*
Soft totals: A soft total is any hand that has an Ace as one of the first two cards, the ace counts as 11 to start.
Soft 20 (A,9) always stands.
Soft 19 (A,8) doubles against dealer 6, otherwise stand.
Soft 18 (A,7) doubles against dealer 2 through 6, and hits against 9 through Ace, otherwise stand.
Soft 17 (A,6) doubles against dealer 3 through 6, otherwise hit.
Soft 16 (A,5) doubles against dealer 4 through 6, otherwise hit.
Soft 15 (A,4) doubles against dealer 4 through 6, otherwise hit.
Soft 14 (A,3) doubles against dealer 5 through 6, otherwise hit.
Soft 13 (A,2) doubles against dealer 5 through 6, otherwise hit.

Hard totals: A hard total is any hand that does not start with an ace in it, or it has been dealt an ace that can only be counted as 1 instead of 11.
17 and up always stands.
16 stands against dealer 2 through 6, otherwise hit.
15 stands against dealer 2 through 6, otherwise hit.
14 stands against dealer 2 through 6, otherwise hit.
13 stands against dealer 2 through 6, otherwise hit.
12 stands against dealer 4 through 6, otherwise hit.
11 always doubles.
10 doubles against dealer 2 through 9 otherwise hit.
9 doubles against dealer 3 through 6 otherwise hit.
8 always hits.

from: https://www.blackjackapprenticeship.com/blackjack-strategy-charts/
Assumption: double-down has already occured, skip those evaluations.
*/
func (tp *AiPlayer) OfferHit(d blackjack.Dealer, h blackjack.Hand) bool {
	dealerFaceUp := d.FaceUp()
	handContainsAce := handHasAces(h)

	if handContainsAce {
		switch h.Evaluate() {
		case 19, 20:
			return false
		case 18:
			return dealerFaceUp >= 9 && dealerFaceUp <= 11
		default:
			return false
		}
	} else {
		switch h.Evaluate() {
		case 4, 5, 6, 7, 8, 9, 10, 11:
			// Some of these values should always double, oh well...should hopefully be covered there.
			return true
		case 12:
			return dealerFaceUp <= 4 && dealerFaceUp >= 6
		case 13, 14, 15, 16:
			return dealerFaceUp > 7
		default:
			return false
		}
	}
}

// Side bets are not worth it to this player.
func (tp *AiPlayer) OfferSideBet(d blackjack.Dealer, h blackjack.Hand) int {
	return 0
}

/*
Splits:
Always split aces.
Never split tens.
A pair of 9’s splits against dealer 2 through 9, except for 7, otherwise stand.
Always split 8’s
A pair of 7’s splits against dealer 2 through 7, otherwise hit.
A pair of 6’s splits against dealer 2 through 6, otherwise hit.
A pair of 5’s doubles against dealer 2 through 9, otherwise hit.
A pair of 4’s splits against dealer 5 and 6, otherwise hit.
A pair of 3’s splits against dealer 2 through 7, otherwise hit.
A pair of 2’s splits against dealer 2 through 7, otherwise hit.
*/
func (tp *AiPlayer) OfferSplit(d blackjack.Dealer, h blackjack.Hand) bool {
	handHasAces := handHasAces(h)
	dealerFaceUp := d.FaceUp()

	if handHasAces {
		return true
	} else {
		switch h.Cards[0].Value {
		case 9:
			switch dealerFaceUp {
			// Omit 7 -- we do not split on dealer faceup value of 7
			case 2, 3, 4, 5, 6, 8, 9:
				return true
			default:
				return false
			}
		case 8:
			return true
		case 7:
			return dealerFaceUp >= 2 && dealerFaceUp <= 7
		case 6:
			return dealerFaceUp >= 2 && dealerFaceUp <= 6
		case 4:
			return dealerFaceUp >= 5 && dealerFaceUp <= 6
		case 2, 3:
			return dealerFaceUp >= 2 && dealerFaceUp <= 7
		default:
			// Never split 5, 10
			return false
		}
	}
}

// Let our test player know that the game is over and also why.
func (tp *AiPlayer) Done(doneMsg string) {
	fmt.Printf("Game is over for ai with id %v -- message from game: %v\n", tp.id, doneMsg)
	fmt.Printf("Ended with points: %v\n", tp.points)
}

// Returns true if any card in the Hand contains an ace.
func handHasAces(h blackjack.Hand) bool {
	for _, c := range h.Cards {
		if c.Value == 1 {
			return true
		}
	}

	return false
}
