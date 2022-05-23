package terminalplayer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Ephex2/gophercises/10.BlackJackGame/blackjack"
)

type Player struct {
	id         int
	points     int
	hands      []blackjack.Hand
	currentBet int
	sideBet    int
}

// Uses the PlayerInfos's Get*() methods to set all the properties of the Player struct.
func (tp *Player) UpdateInfo(pi blackjack.PlayerInfo) {
	// Check if info values have changed, print them if True
	// Exclude hands; these will be printed when needed.
	printInfo := false

	if tp.id != pi.GetId() ||
		tp.points != pi.GetPoints() ||
		tp.currentBet != pi.GetCurrentBet() ||
		tp.sideBet != pi.GetSideBet() {
		printInfo = true
	}

	// Update current values
	tp.id = pi.GetId()
	tp.points = pi.GetPoints()
	tp.hands = pi.GetHands()
	tp.currentBet = pi.GetCurrentBet()
	tp.sideBet = pi.GetSideBet()

	// print values if changes occur, preferably print info before hand(s)
	if printInfo {
		tp.printInfo()
	}
}

func (tp *Player) printInfo() {
	fmt.Println("Current player information:")
	fmt.Printf("\tID: %v\n"+
		"\tPoints: %v\n"+
		"\tCurrent Bet: %v\n"+
		"\tSide bet: %v\n", tp.id, tp.points, tp.currentBet, tp.sideBet)

	// try to call this in any function that is defined on player interface other than placing bet.
	// may revise, I'm not sure what's more user-friendly yet.
}

func (tp *Player) printHands() {
	for i, h := range tp.hands {
		fmt.Printf("Cards in hand %v:\n", i+1)
		h.Print(true)
	}

	// mainly use this whenever the hand needs to be known
}

func getPlayerResponseInt(message string) (out int) {
	valueReceived := false
	reader := bufio.NewReader(os.Stdin)

	for !valueReceived {
		fmt.Printf("%v\n", message)
		in, _ := reader.ReadString('\n')
		in = strings.Replace(in, "\n", "", 1)
		in = strings.Replace(in, "\r", "", 1)

		temp, err := strconv.Atoi(in)

		if err == nil && temp >= 0 {
			out = temp
			valueReceived = true
		} else {
			fmt.Printf("Error received after attempting to read your input: %v\n", err.Error())
		}
	}

	return out
}

func getPlayerResponseBool(message string) (out bool) {
	reader := bufio.NewReader(os.Stdin)
	receivedInput := false

	for !receivedInput {
		fmt.Printf("%v\n", message)
		fmt.Printf("Values accepted are\n 1, t, T, TRUE, true, True for a 'yes'\n0, f, F, FALSE, false, False for a 'no'\n")
		in, _ := reader.ReadString('\n')
		in = strings.Replace(in, "\n", "", 1)
		in = strings.Replace(in, "\r", "", 1)
		temp, err := strconv.ParseBool(in)

		if err == nil {
			out = temp
			receivedInput = true
		} else {
			fmt.Printf("Error received after attempting to read your input: %v\n", err.Error())
		}
	}

	return out
}

func (tp *Player) OfferBet() int {
	msg := "Place your bet for the next round!\nPlease note, a bet of 0 will end the game for you.\nBetting more points than you currently have will have you go all-in with the points you currently have.\n"
	out := getPlayerResponseInt(msg)
	return out
}

func (tp *Player) OfferDoubleDown(d blackjack.Dealer, h blackjack.Hand) bool {
	d.PrintFaceUp()
	msg := "Due to your starting hand, you've been offered a double-down!\nIf you accept, you will double your bet, then draw one more card and your hand will be closed.\nThe card will remain hidden until all players have finished their turns."
	tp.printHands()
	out := getPlayerResponseBool(msg)
	return out
}

func (tp *Player) OfferHit(d blackjack.Dealer, h blackjack.Hand) bool {
	d.PrintFaceUp()
	tp.printHands()
	msg := "You've been offered a hit off the top of the deck!"
	out := getPlayerResponseBool(msg)
	return out
}

func (tp *Player) OfferSideBet(d blackjack.Dealer, h blackjack.Hand) int {
	d.PrintFaceUp() // may be redudant; must be ace if side bet is offered
	msg := "You've been offered a sidebet!\nSidebets allow players to have an insurance policy agaist the house's blackjack.\nOnly offered when dealer is showing an ace."
	out := getPlayerResponseInt(msg)
	return out
}

func (tp *Player) OfferSplit(d blackjack.Dealer, h blackjack.Hand) bool {
	d.PrintFaceUp()
	tp.printHands()
	msg := "You've been offered a split!\nNote that you need to pay your original bet to split your hand.\n"
	out := getPlayerResponseBool(msg)
	return out
}

func (tp *Player) Done(doneMsg string) {
	fmt.Printf("Game over! Message from the game: %v\n", doneMsg)
	fmt.Printf("Ended with %v points\n", tp.points)
}
