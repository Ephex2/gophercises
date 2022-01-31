package blackjack

import (
	"github.com/Ephex2/gophercises/10.DeckOfCards/deck"
)

type Game struct {
	state        *state
	Options      Options
	roundsPlayed int
}

type state struct {
	Dealer     *Dealer
	playerInfo *[]PlayerInfo
	deck       *deck.Deck
}

type Dealer struct {
	faceUp   []deck.Card
	faceDown deck.Card
}

type Options struct {
	Decks          int
	StartingPoints int
	Rounds         int
}

// Returns the int value of the card that the dealer is currently showing
// Allows us to make the faceup []card property unexported, allowing us to keep players from editing it.
func (d *Dealer) FaceUpValue() int {
	publicHand := Hand{d.faceUp}
	return publicHand.Evaluate()
}

func (d *Dealer) setup(deckPointer *deck.Deck) {
	d.faceDown = deckPointer.Draw(1)[0]
	d.faceUp = deckPointer.Draw(1)
}

// Get the current value of the dealer's hand, from the game's perspective.
// Method cannot be exported otherwise Players could evaluate the hidden value of a dealer's hand.
func (d *Dealer) evaluateHand() int {
	cards := append(d.faceUp, d.faceDown)
	h := Hand{Cards: cards}

	return h.Evaluate()
}

// Sets-up a game for play. Will need to call .Play() on the returned Game pointer, passing in the Players to start the game.
func New(opts Options) *Game {
	if opts.Decks == 0 {
		// Default case -- use 6 decks according to da rulez
		opts.Decks = 6
	}

	if opts.StartingPoints == 0 {
		opts.StartingPoints = 100
	}

	deck := deck.NewMultipleDeck(opts.Decks)
	deck.Shuffle()
	// Set plastic card to reduce impact of card counting - remove part of the deck.
	if deckSize := len(deck); deckSize%2 == 0 {
		deck = deck[0:(deckSize / 2)]
	} else {
		deck = deck[0:((deckSize + 1) / 2)]
	}

	state := &state{deck: &deck}
	return &Game{Options: opts, state: state}
}

func (g *Game) Play(players []Player) {
	//	Assign player IDs and starting points, game's deck should already be shuffled
	for i, player := range players {
		pi := PlayerInfo{id: i + 1, points: g.Options.StartingPoints, player: player}
		*g.state.playerInfo = append(*g.state.playerInfo, pi)
	}

	// Determine number of rounds to play. Default is until each player busts or leaves the game.
	for !g.gameOver() {
		g.playRound()
		g.roundsPlayed += 1
	}
}

func (g *Game) gameOver() bool {
	if len(*g.state.playerInfo) == 0 {
		return true
	}

	if g.roundsPlayed >= g.Options.Rounds && g.Options.Rounds > 0 {
		return true
	}

	return false
}

func (g *Game) playRound() {
	/* Turn order:
	1. Take bets
	2. Dealer shows face-up card.
	3. Players in turn-order perform actions.
		- If dealer's face-up card is an Ace, they can sidebet.
		- If a hand has two cards and the cards' values are equal, split is possible.
	4. Dealer finishes their turn.
	5. Evaluate how to update points.
	-- Throughout; eliminate players when possible.
	*/
	var doubleDown bool
	g.placeBets()
	g.removeDonePlayers()
	g.state.Dealer.setup(g.state.deck)

	// Turn order based on order of players passed in, for the moment
	for _, p := range *g.state.playerInfo {
		p.hands[0] = Hand{Cards: g.state.deck.Draw(2)}
		p.sendInfo()

		// Offer a side-bet (insurance bet) if dealer is showing an ace
		if g.state.Dealer.FaceUpValue() == 11 {
			sideBet := p.player.OfferSideBet(g.state.Dealer)
			if sideBet > 0 {
				// Ensure sidebet is not too large, if so, set it to all current points.
				//disposablePoints = p.points - p.currentBet
				p.sideBet = sideBet
			}
		}

		// Perform double-down and blackjack cases first, as they close options for all hands this round.
		// Evaluate if a double-down could occur; ensure player has enough points and is in the right score range.
		if p.hands[0].Evaluate() >= 9 && p.hands[0].Evaluate() <= 11 {
			if p.currentBet > p.points-p.currentBet {
				doubleDown = p.player.OfferDoubleDown(g.state.Dealer)
			}
		}

		// If a player naturals 21, play stops for them. Similarly, if they double-down play is over.
		// Players will perform regular Hit, Split and Stand operations in the last else block
		if p.hands[0].Evaluate() == 21 {
			// BlackJack ! Pay out 1.5 times bet. Rounding down.
			// TODO: use decimals instead of rounding down integer values for Points.
			pointIncrease := p.currentBet + (p.currentBet / 2)
			if p.currentBet%2 != 0 {
				pointIncrease = p.currentBet + ((p.currentBet - 1) / 2)
			}

			// Will evaluate payout at end of round. If dealer also hits 21, will not lose any.
			// This allows us to not have to check to make sure we can afford the new bet.
			p.currentBet += pointIncrease

		} else if doubleDown {
			// Draw a 'hidden' card and close player options for the round; they won't enter into the else{} block.
			p.hands[0].Cards = append(p.hands[0].Cards, g.state.deck.Draw(1)...)
			p.sendInfo()
			p.currentBet = p.currentBet + p.currentBet
		} else {
			// Increase in the len of p.hands will make this loop again for each split.
			for i := 0; i < len(p.hands); i++ {
				var hit = true

				for hit {
					if p.hands[i].Evaluate() == -1 {
						// Player is bust
						break
					}

					// Check if split is possible, offer split to player. If accepted, loop may run multiple times.
					// Ensure player has enough points for split ; each split costs the current bet.
					// 3 splits costs 4 * the current bet, for example.
					if len(p.hands[i].Cards) == 2 && p.hands[i].Cards[0] == p.hands[i].Cards[1] {
						if p.currentBet*(len(p.hands)+1) <= p.points {
							if p.player.OfferSplit(g.state.Dealer) {
								p.hands[i].Cards = []deck.Card{
									p.hands[i].Cards[0],
								}

								p.hands[i+1] = Hand{
									[]deck.Card{
										p.hands[i].Cards[0],
									},
								}
							}
						}
					}

					hit = p.player.OfferHit(g.state.Dealer)
				}

			}
		}
	}

	g.dealerTurn(doubleDown)
	g.evaluatePayout()
}

// Take starting bets, house needs to keep a tally of each player's bet.
// If a player attempts to bet more than they have they are all in (bet all remaining points).
func (g *Game) placeBets() {
	for _, playerInfo := range *g.state.playerInfo {
		playerInfo.PlaceBet()
	}
}

func (g *Game) dealerTurn(doubleDown bool) {
	hit := true
	for hit {
		handValue := g.state.Dealer.evaluateHand()
		if handValue >= 17 || handValue == -1 {
			hit = false
		}

		// Take the hit.
		g.state.Dealer.faceUp = append(g.state.Dealer.faceUp, g.state.deck.Draw(1)...)
	}
}

func (g *Game) evaluatePayout() {
	dealerValue := g.state.Dealer.evaluateHand()

	for _, p := range *g.state.playerInfo {
		// Start by evaluating if a player warrants a sidebet payout.
		// Otherwise, collect the sidebet.
		// This is 0 most of the time strategically and is so by default.
		if p.sideBet != 0 {
			if g.state.Dealer.evaluateHand() == 21 {
				// Sidebets have a 2-to-1 payout

				p.player.UpdatePoints(p.sideBet * 2)
			} else {
				p.player.UpdatePoints(p.sideBet * -1)
			}

			p.sendInfo()
		}

		for _, h := range p.hands {
			// Each split hand is at value p.currentBet
			playerValue := h.Evaluate()

		}
	}
}

// Removes all players who are done (points == 0)
func (g *Game) removeDonePlayers() {
	for i, p := range *g.state.playerInfo {
		if p.points == 0 {
			g.removePlayer(p.id, i)
		}
	}
}

// Removes player from a game, if the game is in progress.
func (g *Game) removePlayer(id int, index int) {
	if len(*g.state.playerInfo) > 0 {
		oldPlayerInfo := *g.state.playerInfo
		newPlayerInfo := append(oldPlayerInfo[0:index], oldPlayerInfo[index+1:]...)
		*g.state.playerInfo = newPlayerInfo
	}
}
