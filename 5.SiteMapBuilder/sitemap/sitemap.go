package sitemap

import (
	"bufio"
	"encoding/xml"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/Ephex2/gophercises/5/linkparser"
)

const SitemapSchemaVersion = "http://www.sitemaps.org/schemas/sitemap/0.9"

type Sitemap struct {
	Xmnls   string `xml:"xmnls,attr"` // field is public to permit marshalling, should not be modified
	Url     []url  `xml:"url"`
	baseUri string
	urls    []string // easily accessible string slice of all Loc values of Url structs. Not to be output as part of sitemap, to be used for comparison only.
}

type url struct {
	Loc string `xml:"loc"`
}

// This struct allows us to unmarshall the xml with the root node having name urlset rather than Sitemap, the name of the type
// Based on solution from: https://stackoverflow.com/questions/12398925/go-xml-marshalling-and-the-root-element
// Even if we named the struct Urlset we would need a similar struc to make the 'U' lowercase when marshalling.
type marshallableSitemap struct {
	XMLName string `xml:"urlset"`
	Sitemap
}

func NewSiteMap() (s Sitemap) {
	s.Xmnls = SitemapSchemaVersion
	return s
}

func (s *Sitemap) BuildSitemap(uri string, maxDepth int) {
	if !strings.HasPrefix(uri, "https://") && !strings.HasPrefix(uri, "http://") {
		log.Fatal("FATAL - Uri provided does not start with http:// or https://, aborting. Uri provided: " + uri)
	}

	// Setup MaxDepth in case where it should be 'unlimited'
	if maxDepth == -1 {
		maxDepth = math.MaxInt
	}

	// Establish baseUri for function call
	s.baseUri = strings.TrimSuffix(uri, "/")
	s.AddUrl(s.baseUri)

	// perform recurstive crawling operation
	s.crawlSite(s.baseUri, maxDepth, 0)
}

func (s *Sitemap) crawlSite(uri string, maxDepth int, depth int) {

	if depth > maxDepth {
		// return from this recursive call, max depth exceepded
		log.Printf("DEBUG - Max link crawl depth exceeded. Actual depth: %v. Maximum depth: %v\n", depth, maxDepth)
		return
	}

	// Get web page contents
	log.Println("INFO - Getting: " + uri)
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err.Error())
	}

	bufReader := bufio.NewReader(res.Body)

	// Go through the links and add them to the site variable. Call function recursively if we have not yet added the uri to the sitemap.
	links := linkparser.GetLinks(bufReader)
	for _, link := range links {
		// Escape Xml characters in link
		link.Href = escapeXml(link.Href)

		// Ensure that relative links are fully qualified
		if string(link.Href[0]) == "/" && link.Href != "/" {
			link.Href = s.baseUri + link.Href
		}

		// Only crawl internal links
		if strings.HasPrefix(link.Href, s.baseUri) {
			// Call links recursively; ensure we do not revisit URIs we've already hit.
			visitedURIs := s.GetUrls()
			_, found := find(visitedURIs, link.Href)
			_, found2 := find(visitedURIs, link.Href+"/")
			if !found && !found2 {
				s.AddUrl(link.Href)
				s.crawlSite(link.Href, maxDepth, depth+1)
			}
		}
	}
}

func escapeXml(uri string) (escapedUri string) {
	// Implemented a simple xml escape function from string to string. Not sure on performance but it seems straightforward enough
	// xml package functions to unescape needed conversions to byte and use of readers.... on top of this, creating a byte buffer seemed to double the string for some reason.
	escapedUri = strings.Replace(uri, "&", "&amp;", -1)
	escapedUri = strings.Replace(escapedUri, "<", "&lt;", -1)
	escapedUri = strings.Replace(escapedUri, ">", "&gt;", -1)
	escapedUri = strings.Replace(escapedUri, "\"", "&quot;", -1)
	escapedUri = strings.Replace(escapedUri, "'", "&apos;", -1)
	return escapedUri
}

func (s *Sitemap) GetUrls() []string {
	// Return slice of all loc values within sitemap object
	return s.urls
}

func find(slice []string, val string) (int, bool) {
	// Determines whether a given string value (val) is found within a slice of strings (slice)
	// Taken from: https://golangcode.com/check-if-element-exists-in-slice/
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func (s *Sitemap) AddUrl(uri string) {
	// Add url to sitemap object
	if !strings.HasSuffix(uri, "/") {
		uri = uri + "/"
	}

	var urlVar url
	urlVar.Loc = uri
	s.Url = append(s.Url, urlVar)
	s.urls = append(s.urls, uri)
}

func (s *Sitemap) WriteXml(path string) {
	// This temporary struct allows us to unmarshall the xml with the root node having name urlset rather than Sitemap, the name of the type
	// Based on solution from: https://stackoverflow.com/questions/12398925/go-xml-marshalling-and-the-root-element
	tmp := marshallableSitemap{Sitemap: *s}

	// Marshall sitemap into xml bytes
	b, err := xml.MarshalIndent(tmp, "", "    ")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create file of deletecontents if already present
	siteMapFile, err := os.Create(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer siteMapFile.Close()

	// Add XML header string
	_, err = siteMapFile.WriteString(xml.Header)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Write marshalled xml bytes to file
	_, err = siteMapFile.Write(b)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("INFO - Sitemap written to path: " + path)
}

func (s *Sitemap) LoadXml(path string) (err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	*s = NewSiteMap()
	err = xml.Unmarshal(file, &s)
	if err != nil {
		return err
	}

	return nil
}
