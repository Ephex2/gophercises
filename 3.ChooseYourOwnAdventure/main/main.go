package main

import (
	"flag"
	"fmt"
	"gopheradventures/model"
	"gopheradventures/presentation"
	"net/http"
	"os"
	"path/filepath"
)

var adventureData []byte

// flag vars
var DefaultArc string
var cli bool
var jsonFilePath string

func init() {
	// Flag setup
	flag.StringVar(&jsonFilePath, "filepath", "", "Path to the json file defining model arcs. Default is the json string in gopher.json.go in the model package.")
	flag.StringVar(&DefaultArc, "defaultarc", "intro", "Title of the arc that the adventure should start on. In the example given (gopher.json.go), the arc is intro.")
	flag.BoolVar(&cli, "cli", false, "When specified, the choose your own adventure experience will be presented through the terminal rather than a webpage.")
	flag.Parse()

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
	model.RuntimeArcs = model.NewArc(adventureData)

	fmt.Printf("Value of cli flag in main block: %v", cli)

	if cli {
		presentation.CliFlow(model.RuntimeArcs, DefaultArc)
	} else {
		presentation.SetDefaultArc(DefaultArc)
		http.HandleFunc("/", presentation.TemplateFlow)

		for _, arc := range model.RuntimeArcs {
			arcPath := "/" + arc.Title
			http.HandleFunc(arcPath, presentation.TemplateFlow)
		}

		http.ListenAndServe(":8080", http.DefaultServeMux)
	}
}
