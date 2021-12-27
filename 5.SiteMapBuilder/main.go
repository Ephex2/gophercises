package main

import (
	"flag"

	"github.com/Ephex2/gophercises/5/sitemap"
)

var baseUri string
var filePath string
var depth int

func init() {
	flag.StringVar(&baseUri, "uri", "https://courses.calhoun.io/", "Specifies the URI that is the root of the site that you desire to make a sitemap of.")
	flag.StringVar(&filePath, "file", "./sitemap.xml", "Specifies the file path of the output sitemap XML file. Will default to ./sitemap.xml .")
	flag.IntVar(&depth, "depth", -1, "Specifies the maximum depth to look for links on a given site. If not specified, will keep crawling until all internal links have been crawled.")
	flag.Parse()
}

func main() {
	if depth < -1 {
		// Someone is being a weirdo. Set to -1
		depth = -1
	}

	site := sitemap.NewSiteMap()
	site.BuildSitemap(baseUri, depth)
	site.WriteXml(filePath)
}
