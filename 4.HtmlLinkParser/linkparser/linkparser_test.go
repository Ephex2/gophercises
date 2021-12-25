package linkparser_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Ephex2/gophercises/4/linkparser"
)

// doing this without mocking or anything fancy. Main function of interest is getLinks() since it calls parseHref and sanitizeText
// using  something similar to table-based tests since they seem to be relatively common.
// TODO: write tests for other functions in linkparser.go.
func TestGetLinks(t *testing.T) {
	pathes := []string{"../sampleHtml/ex1.html", "../sampleHtml/ex2.html", "../sampleHtml/ex3.html", "../sampleHtml/ex4.html"}
	tables := []struct {
		output []linkparser.LinkOutput
		reader *bufio.Reader
	}{
		{[]linkparser.LinkOutput{
			{
				Href: "/other-page",
				Text: "A link to another page",
			},
		}, &bufio.Reader{}},
		{[]linkparser.LinkOutput{
			{
				Href: "https://www.twitter.com/joncalhoun",
				Text: "Check me out on twitter",
			},
			{
				Href: "https://github.com/gophercises",
				Text: "Gophercises is on Github!",
			},
		}, &bufio.Reader{}},
		{[]linkparser.LinkOutput{
			{Href: "#",
				Text: "Login",
			},
			{
				Href: "/lost",
				Text: "Lost? Need help?",
			},
			{
				Href: "https://twitter.com/marcusolsson",
				Text: "@marcusolsson",
			},
		}, &bufio.Reader{}},
		{[]linkparser.LinkOutput{
			{
				Href: "/dog-cat",
				Text: "dog cat",
			},
		}, &bufio.Reader{}},
	}

	// Overwrite default readers in test table with specific readers for test files
	for i, path := range pathes {
		reader, err := os.Open(path)
		if err != nil {
			str := fmt.Sprintf("Unable to open file: %v. Error given by GO is: %v", path, err.Error())
			log.Fatal(str)
		}

		r := bufio.NewReader(reader)
		tables[i].reader = r
	}

	// Look through each test and table output slice and check if values of href and text do have matches.
	// Potential for error if values are meant to repeat a certain number of times (will not validate number of times a value appears).
	// Advantage of doing it this way: return order of GetLinks() vis à vis order in test tables is irrelevant
	for i, table := range tables {
		testOutput := linkparser.GetLinks(table.reader)

		var HrefFound bool = false
		var TextFound bool = false

		for _, tableValue := range table.output {
			for _, outputValue := range testOutput {
				if outputValue.Href == tableValue.Href {
					HrefFound = true
				}
				if outputValue.Text == tableValue.Text {
					TextFound = true
				}
			}

			if !HrefFound {
				t.Errorf("Value of Href: %v of table with index: %v was not found in the test output.\nTest output:%v\n", tableValue.Href, i, testOutput)
			}
			if !TextFound {
				t.Errorf("Value of Text: %v of table with index: %v was not found in the test output.\nTest output:%v\n", tableValue.Text, i, testOutput)
			}

			HrefFound = false
			TextFound = false
		}
	}
}