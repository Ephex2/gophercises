package blackjack_test

import (
	"testing"
	"time"

	"github.com/Ephex2/gophercises/10.DeckOfCards/blackjack"
)

// This is a sort of integration test between a test ai and its use by the blackjack game.
// Very basic, just makes sure the game completes before an arbitrary 'time is up'.
func TestBlackJackGame(t *testing.T) {
	var players []blackjack.Player
	playerCount := 10
	for i := 0; i < playerCount; i++ {
		players = append(players, &TestPlayer{})
	}

	maxDuration := time.Duration(5)
	opt := blackjack.Options{Decks: 10, StartingPoints: 100}
	g := blackjack.New(opt)

	c := make(chan bool)
	timeIsUp := make(chan bool)

	go func() {
		g.Play(players)
		c <- true
	}()

	go func() {
		time.Sleep(time.Second * maxDuration)
		timeIsUp <- true
	}()

	select {
	case <-c:
		t.Logf("Game with %v ai players completed before timeout of %v seconds.", playerCount, int(maxDuration))
	case <-timeIsUp:
		t.Errorf("Game with %v ai players was not able to complete before the timeout of %v seconds", playerCount, int(maxDuration))
	}
}
