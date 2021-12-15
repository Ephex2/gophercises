package model

import (
	"encoding/json"
	"fmt"
)

type adventureRaw map[string]json.RawMessage

type Arcs map[string]Arc

type Arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var RuntimeArcs Arcs

func NewArc(adventureData []byte) Arcs {
	// Need to unmarshal object with rawJson data first to make the raw map
	var a adventureRaw
	err := json.Unmarshal(adventureData, &a)
	if err != nil {
		errMsg := fmt.Sprintf("Error unmarshaling RAW json for adventures: %v\n", err.Error())
		panic(errMsg)
	}

	// Loop through the raw map to unmarshal the raw json data in the true Arcs map
	var advMap = make(Arcs)
	for arcName, value := range a {
		var tempArc Arc
		err = json.Unmarshal(value, &tempArc)
		if err != nil {
			errMsg := fmt.Sprintf("Error unmarshalling inner JSON for arc named: %v. Error is: %v\n", arcName, err.Error())
			panic(errMsg)
		}
		if _, ok := advMap[arcName]; ok {
			// Decided that having duplicate arc names warrants a panic, since it should not occur.
			errMsg := fmt.Sprintf("Duplicate entries found for arc with name: %v", arcName)
			panic(errMsg)
		} else {
			advMap[arcName] = tempArc
		}
	}

	return advMap
}
