package blackjack

import (
	"fmt"

	"github.com/Ephex2/gophercises/10.BlackJackGame/deck"
)

type Game struct {
	state        *state
	Options      Options
	roundsPlayed int
}

type state struct {
	Dealer     *Dealer
	playerInfo []PlayerInfo
	deck       *deck.Deck
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

	deck := setupDeck(opts)
	state := &state{deck: &deck, Dealer: &Dealer{}}
	return &Game{Options: opts, state: state}
}

// Resets the deck used by the dealer. Assumes plastic card is always set in the middle of the multideck.
func setupDeck(opts Options) deck.Deck {
	deck := deck.NewMultipleDeck(opts.Decks)
	deck.Shuffle()
	// Set 'plastic card' to reduce impact of card counting - remove part of the deck.
	if deckSize := len(deck.Cards); deckSize%2 == 0 {
		deck.Cards = deck.Cards[0:(deckSize / 2)]
	} else {
		deck.Cards = deck.Cards[0:((deckSize + 1) / 2)]
	}

	return deck
}

func (g *Game) Play(players []Player) {
	//	Assign player IDs and starting points, game's deck should already be shuffled
	for i, player := range players {
		pi := PlayerInfo{id: i + 1, points: g.Options.StartingPoints, player: player}

		if g.state.playerInfo == nil {
			g.state.playerInfo = []PlayerInfo{pi}
		} else {
			g.state.playerInfo = append(g.state.playerInfo, pi)
		}
	}

	// Determine number of rounds to play. Default is until each player busts or leaves the game.
	for !g.isGameOver() {
		g.playRound()
		g.roundsPlayed += 1
	}

	for i := range g.state.playerInfo {
		g.state.playerInfo[i].done("You've made it to the end of the game!")
	}
}

func (g *Game) isGameOver() bool {
	if len(g.state.playerInfo) == 0 {
		return true
	}

	if g.roundsPlayed >= g.Options.Rounds && g.Options.Rounds > 0 {
		return true
	}

	return false
}

/* Turn order:
 - 1. Take bets
 - 2. Dealer shows face-up card.
 - 3. Players in turn-order perform actions.
  - If the face-up dealer card is an Ace, they can sidebet.
  - If a hand has two cards and the values are equal, split is possible.
 - 4. Dealer finishes their turn.
 - 5. Evaluate how to update points.
-- Eliminate players when possible.
*/
func (g *Game) playRound() {
	var doubleDown bool
	g.setupNewRound()

	// Turn order based on order of players passed in, for the moment
	for i, p := range g.state.playerInfo {
		// Offer a side-bet (insurance bet) if dealer is showing an ace
		g.evaluateSidebet(p, p.hands[0])

		// Perform double-down and blackjack cases first, as they close options for all of a player's hands this round.
		// Evaluate if a double-down could occur; ensure player has enough points and is in the right score range.
		// If a player naturals 21, play stops for them. Similarly, if they double-down play is over.
		// Players will perform regular Hit, Split and Stand operations in the last else block
		if p.hands[0].Evaluate() == 21 {
			// BlackJack ! Pay out 1.5 times bet. Rounding down.
			pointIncrease := p.currentBet + (p.currentBet / 2)
			if p.currentBet%2 != 0 {
				pointIncrease = p.currentBet + ((p.currentBet - 1) / 2)
			}

			// Will evaluate payout at end of round. If dealer also hits 21, will not lose any.
			// This allows us to not have to check to make sure we can afford the new bet.
			p.currentBet += pointIncrease

		} else if p.hands[0].Evaluate() >= 9 && p.hands[0].Evaluate() <= 11 {
			if p.currentBet > p.points-p.currentBet {
				if doubleDown = p.player.OfferDoubleDown(*g.state.Dealer, p.hands[0]); doubleDown {
					// Draw a 'hidden' card and close player options for the round; they won't enter into the else{} block.
					p.hands[0].Cards = append(p.hands[0].Cards, g.drawCardsFromDeck(1)...)
					p.currentBet = p.currentBet + p.currentBet
				}
			}
		}

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
				if p.offerSplit(*g.state.Dealer, &p.hands[i]) {
					p.hands[i] = Hand{
						[]deck.Card{
							p.hands[i].Cards[0],
						},
					}

					p.hands = append(p.hands, Hand{
						[]deck.Card{
							p.hands[i].Cards[0],
						},
					},
					)

					p.updateInfo()
				}

				hit = p.offerHit(*g.state.Dealer, p.hands[i])
				if hit {
					p.hands[i].Cards = append(p.hands[i].Cards, g.drawCardsFromDeck(1)...)
				}
			}
		}

		// Ensure playerinfo array has current player data; I find this easier to read than directly using array index, can refactor.
		g.state.playerInfo[i] = p
	}

	g.dealerTurn()
	g.evaluatePayout(doubleDown)
}

func (g *Game) setupNewRound() {
	g.removeDonePlayers()
	for i, p := range g.state.playerInfo {
		// Setup player
		p.newRound(g)
		g.state.playerInfo[i] = p
	}
	g.placeBets()

	dealerCards := g.drawCardsFromDeck(2)
	g.state.Dealer.setup(dealerCards)
}

// Take starting bets, house needs to keep a tally of each player's bet.
// If a player attempts to bet more than they have they are all in (bet all remaining points).
func (g *Game) placeBets() {
	for i, p := range g.state.playerInfo {
		p.offerBet()

		if p.currentBet < 1 {
			p.done("Betting 0 or less indicates that you wish to quit the table.")
			g.removePlayer(p.id)
			return
		}

		g.state.playerInfo[i] = p
	}
}

// Sidebets allow players to have an insurance policy agaist the house's blackjack.
// Only offered when dealer is showing an ace.
func (g *Game) evaluateSidebet(p PlayerInfo, h Hand) {
	if g.state.Dealer.FaceUp() == 11 {
		p.offerSideBet(*g.state.Dealer, h)
	}
}

// Take actions for the dealer. Must hit if not busted or hand value is < 17.
func (g *Game) dealerTurn() {
	hit := true
	for hit {
		handValue := g.state.Dealer.evaluateHand()

		// Dealer must evaluate Ace as eleven, no special handling necessary.
		// Evaluate already returns the highest possible value of a hand.
		if handValue >= 17 || handValue == -1 {
			hit = false
			continue
		}

		// Take the hit.
		g.state.Dealer.faceUp = append(g.state.Dealer.faceUp, g.drawCardsFromDeck(1)...)
	}
}

// Evaluate each player's payout, including sidebets and
func (g *Game) evaluatePayout(doubleDown bool) {
	dealerValue := g.state.Dealer.evaluateHand()

	for i, p := range g.state.playerInfo {
		if p.sideBet != 0 {
			var sideBetValue int
			if g.state.Dealer.evaluateHand() == 21 {
				// Sidebets have a 2-to-1 payout
				sideBetValue = p.sideBet * 2
			} else {
				sideBetValue = p.sideBet * -1
			}

			err := p.updatePoints(sideBetValue)
			if err != nil {
				p.done(err.Error())
				g.removePlayer(p.id)
			}
		}

		// Each split hand is at value p.currentBet. Non-split hands will simply loop once.
		for _, h := range p.hands {
			var betMultiplier = 1
			fmt.Printf("Player hand value: %v -- Dealer hand value: %v\n", h.Evaluate(), dealerValue)
			if doubleDown {
				betMultiplier = 2
			}

			if playerValue := h.Evaluate(); dealerValue > playerValue {
				betMultiplier *= -1
			} else if playerValue == dealerValue {
				betMultiplier *= 0
			}

			p.updatePoints(betMultiplier * p.currentBet)
		}

		g.state.playerInfo[i] = p
	}
}

// Wraps calls from deck, if deck runs out, automatically recreates one with the appropriate number of cards.
func (g *Game) drawCardsFromDeck(count int) []deck.Card {
	var err error
	cards, err := g.state.deck.Draw(count)

	// Keep drawing from deck until we have drawn enough cards, we should only enter loop if we tried to overdraw.
	// If deck is empty, loop should do nothing but reset deck with setupDeck()
	for err != nil && len(cards) != count {
		var cardsToDraw int
		if cardsToDraw = (count - len(cards)); len(g.state.deck.Cards) < cardsToDraw {
			cardsToDraw = len(g.state.deck.Cards)
		}

		// cardsToDraw should always be <= len(deck), shouldn't run into any overdraw errors.
		cardsDrawn, _ := g.state.deck.Draw(cardsToDraw)

		if cardsDrawn != nil {
			cards = append(cards, cardsDrawn...)
		}

		if len(g.state.deck.Cards) == 0 {
			*g.state.deck = setupDeck(g.Options)
		}
	}

	return cards
}

// Removes all players who are done (points <= 0), and sends a message through the player interface.
func (g *Game) removeDonePlayers() {
	for _, p := range g.state.playerInfo {
		if p.points <= 0 {
			p.done("You are out of points.")
			g.removePlayer(p.id)
		}
	}
}

// Removes player from a game, if the game is in progress. Based on index of the palyer in the PlayerInfo slice.
func (g *Game) removePlayer(id int) {
	if len(g.state.playerInfo) > 0 {
		var index int
		for i, p := range g.state.playerInfo {
			if p.id == id {
				index = i
			}
		}

		oldPlayerInfo := g.state.playerInfo
		newPlayerInfo := append(oldPlayerInfo[0:index], oldPlayerInfo[index+1:]...)
		g.state.playerInfo = newPlayerInfo
	}
}
