package main

import (
	"flag"
	"fmt"
	"gopheradventures/model"
	"path/filepath"
)

var jsonFilePath string
var arcs model.Arcs

func init() {
	flag.StringVar(&jsonFilePath, "jsonfile", "", "Path to the json file defining model arcs. Default path should be sufficient, but if other adventures are written it can be overwritten with this flag.")

	if jsonFilePath == "" {
		// WARNING: Changing relative location of the main.go and gopher.json files will break the default load.
		jsonFilePath = "../gopher.json"
	}

	adventurePath, err := filepath.Abs(jsonFilePath)
	if err != nil {
		errMsg := fmt.Sprintf("Error formatting jsonFilePath %v, the error is: %v\n", jsonFilePath, err.Error())
		panic(errMsg)
	}

	arcs = model.LoadAdventure(adventurePath)

	for name, arc := range arcs {
		fmt.Println(name)
		fmt.Println(arc)
	}

}

func main() {
	// do nothing lol
}
