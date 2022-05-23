package cmd

import (
	"fmt"

	"bufio"
	"os"
	"strconv"

	"strings"

	aiplayer "github.com/Ephex2/gophercises/10.BlackJackGame/aiplayer"
	"github.com/Ephex2/gophercises/10.BlackJackGame/blackjack"
	terminalplayer "github.com/Ephex2/gophercises/10.BlackJackGame/terminalplayer"
	"github.com/spf13/cobra"
)

var numPoints int
var numRounds int
var numAi int

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "plays a game of BlackJack.",
	Long:  `plays a game of BlackJack. The game will prompt for some options, then will prompt you with different game decisions. Number of AI, number of rounds against the house, and number of starting points can be modified.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			getUserInput("starting points", &numPoints)
			getUserInput("ai players", &numAi)
			getUserInput("rounds", &numRounds) //TODO: make an unlimited number of rounds possible.
		*/
		numAi := 1
		numRounds := 3
		numPoints := 10

		opts := blackjack.Options{
			Decks:          10,
			Rounds:         numRounds,
			StartingPoints: numPoints,
		}
		game := blackjack.New(opts)

		aiPlayers := []blackjack.Player{}
		for i := 0; i < numAi; i++ {
			ai := &aiplayer.Player{}
			aiPlayers = append(aiPlayers, ai)
		}

		players := []blackjack.Player{&terminalplayer.Player{}}
		players = append(players, aiPlayers...)

		game.Play(players)
	},
}

func getUserInput(paramName string, output *int) {
	*output = -1
	reader := bufio.NewReader(os.Stdin)

	for *output == -1 {
		fmt.Printf("Enter how many %v you would like to play with, must be >= 0.\n", paramName)
		in, _ := reader.ReadString('\n')
		in = strings.Replace(in, "\n", "", 1)
		in = strings.Replace(in, "\r", "", 1)

		temp, err := strconv.Atoi(in)

		if err == nil && temp >= 0 {
			*output = temp
		} else {
			fmt.Printf("Error received after attempting to read your input: %v\n", err.Error())
		}
	}
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
