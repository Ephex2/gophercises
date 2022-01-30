package linkparser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	xHtml "golang.org/x/net/html"
)

type LinkOutput struct {
	Href string
	Text string
}

func GetLinks(r *bufio.Reader) (output []LinkOutput) {
	doc, err := xHtml.Parse(r)
	if err != nil {
		log.Fatal(err.Error())
	}

	// based on sample recursive parsing of html tree from xHtml package documentation
	var parseLinks func(*xHtml.Node)
	var workingOutput []LinkOutput

	// function defined here so it can use the workingOutput variable without accepting it as a parameter and returning it (is this better though?)
	parseLinks = func(n *xHtml.Node) {
		var output LinkOutput // reset output for each function call

		if n.Type == xHtml.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					output.Href = attr.Val
					text := parseText(n, nil)
					output.Text = text
				}
			}

			if output.Href != "" {
				workingOutput = append(workingOutput, output)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseLinks(c)
		}
	}

	parseLinks(doc)
	return workingOutput
}

func parseText(n *xHtml.Node, parent *xHtml.Node) (text string) {
	// This is to avoid exiting the scope of the current element while parsing the text within it
	if parent == nil {
		parent = n.Parent
	} else if parent == n.Parent {
		return text
	}

	// Data we are interested in for this function is only in text nodes - this allows us to avoid comments, etc.
	if n.Type == xHtml.NodeType(xHtml.TextNode) {
		text += n.Data
	}

	if n.FirstChild != nil {
		text += parseText(n.FirstChild, parent)
	}

	if n.NextSibling != nil {
		text += parseText(n.NextSibling, parent)
	}

	text = sanitizeText(text)
	return text
}

func sanitizeText(text string) (sanitizedText string) {
	// Remove all leading and trailing new-line characters and spaces
	sanitizedText = text
	var oldText string

	for oldText != sanitizedText {
		// This might not be super efficient, but it seems simple enough and works
		oldText = sanitizedText
		sanitizedText = strings.TrimPrefix(sanitizedText, "\n")
		sanitizedText = strings.TrimPrefix(sanitizedText, " ")
		sanitizedText = strings.TrimSuffix(sanitizedText, " ")
		sanitizedText = strings.TrimSuffix(sanitizedText, "\n")
	}

	return sanitizedText
}

func PrintOutput(output []LinkOutput) {
	// Decided to output marshalled json instead of hacking together the same output as in the exercise page.
	// Got introduced to the json.MarshalIndent() function to pretty print json.
	jsonByte, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(jsonByte))
}
