package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type adventureRaw map[string]json.RawMessage

type Arcs map[string]arc

type arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []options `json:"options"`
}

type options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func LoadAdventure(adventurePath string) Arcs {
	jsonData, err := os.ReadFile(adventurePath)
	if err != nil {
		errMsg := fmt.Sprintf("Error loading file for adventures: %v\n", err.Error())
		panic(errMsg)
	}

	// Need to unmarshal object with rawJson data first to make the raw map
	var a adventureRaw
	err = json.Unmarshal(jsonData, &a)
	if err != nil {
		errMsg := fmt.Sprintf("Error unmarshaling RAW json for adventures: %v\n", err.Error())
		panic(errMsg)
	}

	// Loop through the raw map to unmarshal the raw json data in the true Arcs map
	var advMap = make(Arcs)
	for arcName, value := range a {
		var tempArc arc
		err = json.Unmarshal(value, &tempArc)
		if err != nil {
			errMsg := fmt.Sprintf("Error unmarshalling inner JSON for arc named: %v. Error is: %v\n", arcName, err.Error())
			panic(errMsg)
		}
		if _, ok := advMap[arcName]; ok {
			// Decided that having duplicate arc names warrants a panic, since it should not occur.
			errMsg := fmt.Sprintf("Duplicate entries found in %v for arc with name: %v", adventurePath, arcName)
			panic(errMsg)
		} else {
			advMap[arcName] = tempArc
		}
	}

	return advMap
}
