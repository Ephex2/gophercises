package presentation

import (
	"bufio"
	"fmt"
	"gopheradventures/model"
	"os"
	"strings"
)

func CliFlow(arcs model.Arcs, defaultArc string) {
	var nextArc string
	currentArc := arcs[defaultArc]
	fmt.Printf("In this choose ")

	// Loop through different arcs, stopping when we get to the last arc
	for len(currentArc.Options) != 0 {
		printStoryArc(currentArc)

		fmt.Printf("Please enter the story option you would like to follow from the choices above: \n\n")
		nextArc = readArcStringInput(arcs)
		currentArc = arcs[nextArc]
	}

	// Print story details again for last arc
	printStoryArc(currentArc)
}

func printStoryArc(currentArc model.Arc) {
	fmt.Printf("\n")
	fmt.Printf("Title: %v\n\n", currentArc.Title)

	for _, story := range currentArc.Story {
		fmt.Printf("%v\n", story)
	}

	fmt.Printf("\n\n")
	for _, option := range currentArc.Options {
		fmt.Printf("Option title: %v\n", option.Arc)
		fmt.Printf(option.Text + "\n")
	}
	fmt.Printf("\n")
}

func readArcStringInput(arcs model.Arcs) string {
	var arcChoice string
	var err error
	var reader = bufio.NewReader(os.Stdin)
	var i int = 0

	for i == 0 {
		arcChoice, err = reader.ReadString('\n')
		arcChoice = strings.TrimSuffix(arcChoice, "\n")
		arcChoice = strings.TrimSuffix(arcChoice, "\r") // was necessary for windows, doesn't errpr out when \r is absent
		_, ok := arcs[arcChoice]

		if err != nil {
			fmt.Printf("***Unable to read your input, please try again. Error: %v\n", err.Error())
		} else if !ok {
			fmt.Printf("***The arc you entered does not match the arcs listed. You entered: %v\n", arcChoice)
		} else {
			i++
		}
	}

	return arcChoice
}
