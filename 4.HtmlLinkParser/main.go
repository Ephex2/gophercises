package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Ephex2/gophercises/4/linkparser"
)

var htmlPath string

func init() {
	flag.StringVar(&htmlPath, "path", "./sampleHtml/ex2.html", "This flag is used to provide the path towards an html file that must be parsed.")
	flag.Parse()
}

func main() {
	reader, err := os.Open(htmlPath)
	if err != nil {
		str := fmt.Sprintf("Unable to open file: %v. Error given by GO is: %v", htmlPath, err.Error())
		log.Fatal(str)
	}

	r := bufio.NewReader(reader)
	output := linkparser.GetLinks(r)
	linkparser.PrintOutput(output)
}
