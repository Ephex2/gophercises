package main

import (
	"flag"
	"fmt"
	"gopheradventures/model"
	"gopheradventures/presentation"
	"os"
	"path/filepath"
)

var adventureData []byte
var arcs model.Arcs

// flag vars
var defaultArc string
var cli bool

func init() {
	// Flag setup
	var jsonFilePath string
	flag.StringVar(&jsonFilePath, "jsonfile", "", "Path to the json file defining model arcs. Default is the json string in gopher.json.go in the model package.")
	flag.StringVar(&defaultArc, "defaultarc", "intro", "Title of the arc that the adventure should start on. In the example given (gopher.json.go), the arc is intro.")
	flag.BoolVar(&cli, "cli", true, "When specified, the choose your own adventure experience will be presented through the terminal rather than a webpage.")

	// Get adventure data depending on flag passed
	if jsonFilePath != "" {
		adventurePath, err := filepath.Abs(jsonFilePath)
		if err != nil {
			errMsg := fmt.Sprintf("Error formatting jsonFilePath %v, the error is: %v\n", jsonFilePath, err.Error())
			panic(errMsg)
		}
		adventureData, err = os.ReadFile(adventurePath)
		if err != nil {
			errMsg := fmt.Sprintf("Error loading file for adventures: %v\n", err.Error())
			panic(errMsg)
		}
	} else {
		adventureData = []byte(model.DefaultAdventureString)
	}
}

func main() {
	arcs = model.LoadAdventure(adventureData)

	if cli {
		presentation.CliFlow(arcs, defaultArc)
	} else {
		presentation.TemplateFlow(arcs, defaultArc)
	}
}
