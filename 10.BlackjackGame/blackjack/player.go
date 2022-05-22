package blackjack

// The player interface allows the game type to interact with the different types of players in a game.
// When being offered a choice in the game, player's will be sent a pointer to the dealer and the current hand.
// Dealer type only allows the dealer's face-up card(s) to be perused by Players.
// Type implementing Player interface should not be defined in the blackjack package.
type Player interface {
	UpdateInfo(PlayerInfo)             // Set a player's info. Players should use the Get*() methods on the PlayerInfo type to learn about their current game state.
	OfferBet() int                     // Returns a player's starting bet for a round.
	OfferDoubleDown(Dealer, Hand) bool // Allows a player to choose whether or not to double-down
	OfferHit(Dealer, Hand) bool        // Offers a player to hit the deck for a new card in the current hand.
	OfferSideBet(Dealer, Hand) int     // Allows players to determine if they want to take a sideBet, and if they do, for how much (int return). A negative value does nothing.
	OfferSplit(Dealer, Hand) bool      // Offers a split to the player. If return is true, game will update the player to have multiple hands, each with the one card in hand.
	Done(string)                       // After a game ends or a player is eliminated, notifies the player that the game is over. The input string is informational and refers to why the game is over for this player.
}
