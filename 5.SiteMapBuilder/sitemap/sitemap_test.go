package sitemap_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/Ephex2/gophercises/5/sitemap"
)

var tables = []struct {
	url        string
	outputpath string
	testpath   string
}{
	{
		"https://courses.calhoun.io",
		"./testMaps/test1.xml",
		"./testMaps/courses.sitemap.xml",
	},
	{
		"https://www.calhoun.io/",
		"./testMaps/test2.xml",
		"./testMaps/calhoun.sitemap.xml",
	},
}

func TestLoadXml(t *testing.T) {
	// Setup
	for _, table := range tables {
		testSiteMap := sitemap.NewSiteMap()
		err := testSiteMap.LoadXml(table.testpath)
		if err != nil {
			t.Fatalf(".LoadXml methof of Sitemap object failed. Error was: %v, Path was: %v", err.Error(), table.testpath)
		}
	}
}

func TestBuildSiteMap(t *testing.T) {
	for _, table := range tables {
		// Setup
		testSiteMap := sitemap.NewSiteMap()
		testSiteMap.LoadXml(table.testpath) // if this errors out, previous test should take care of it, don't handle error.

		outputSiteMap := sitemap.NewSiteMap()
		outputSiteMap.BuildSitemap(table.url, -1)

		// Evaluate
		err := compareSitemapUrls(testSiteMap, outputSiteMap, "TestBuildSiteMap")
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}

func TestWriteXml(t *testing.T) {

	for _, table := range tables {
		// Setup
		testSiteMap := sitemap.NewSiteMap()
		testSiteMap.LoadXml(table.testpath)
		testSiteMap.WriteXml(table.outputpath)

		outputSiteMap := sitemap.NewSiteMap()
		outputSiteMap.LoadXml(table.outputpath)

		// Evaluate
		err := compareSitemapUrls(testSiteMap, outputSiteMap, "TestWriteXml")
		if err != nil {
			t.Errorf(err.Error())
		}

		// Cleanup
		err = os.Remove(table.outputpath)
		if err != nil {
			// Logging error since this may indicate a failure to clean up file handles.
			t.Errorf("Error removing test output file: %v. Error is: %v", table.outputpath, err.Error())
		}
	}
}

func TestMaxDepth(t *testing.T) {
	// Setup
	testMap := sitemap.NewSiteMap()
	testMap.LoadXml("./testMaps/courses.sitemap.0depth.xml")

	outputMap := sitemap.NewSiteMap()

	// Evaluate
	outputMap.BuildSitemap("https://courses.calhoun.io/", 0)
	err := compareSitemapUrls(testMap, outputMap, "TestMaxDepth")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func compareSitemapUrls(should sitemap.Sitemap, actual sitemap.Sitemap, testName string) error {
	// should.Url U actual.Url == should == actual. We hope. Here we evaluate if this is the case. Order of Urls in slice is irrelevant
	// See if there are any elements in should.Url which are not in actual.Url
	for _, url := range should.Url {
		found := false
		for _, url2 := range actual.Url {
			if url.Loc == url2.Loc {
				found = true
			}
		}

		if !found {
			errMsg := fmt.Sprintf("Unable to find url: %v from 'should' within 'actual' for test: %v\n", url.Loc, testName)
			return errors.New(errMsg)
		}
	}

	// Reverse order to see if there are any urls not present in should that are present in actual
	for _, url := range actual.Url {
		found := false
		for _, url2 := range should.Url {
			if url.Loc == url2.Loc {
				found = true
			}
		}

		if !found {
			errMsg := fmt.Sprintf("Unable to find url: %v from 'actual' within 'should' for test: %v\n", url.Loc, testName)
			return errors.New(errMsg)
		}
	}

	return nil
}
