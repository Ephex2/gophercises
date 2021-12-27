# Site map builder - Gophercise 5

This project will take in a url and a path and generate a sitemap of the url, using it as the root site. The site map will be output into the specified path.
Here is a sample (simple) sitemap. Note that this probably wouldn't be a valid root sitemap.xml file to use for google.
```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmnls="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>https://courses.calhoun.io/</loc>
    </url>
    <url>
        <loc>https://courses.calhoun.io/signup?/</loc>
    </url>
    <url>
        <loc>https://courses.calhoun.io/signin?/</loc>
    </url>
    <url>
        <loc>https://courses.calhoun.io/reset-pw?/</loc>
    </url>
</urlset>
```

Note that this tool hasn't been definitively tested, but instead was made to the specs of the fifth gophercises exercise on [https://courses.calhoon.io](https://courses.calhoon.io).
Executable name: ```siteMapBuilder.exe```.
<br></br>
## Optional Flags
Here's an overview of the different flags available for main.go:
- -uri: Specifies the root site to crawl to build the site map. Default value is https://courses.calhoun.io/.
- -path: This flag is used to provide the path where the sitemap xml will be output. You don't have to provide the .xml extension but the format of the file's contents will be xml. Default value is ./sitemap.xml.
- -depth: This flag specifies the maximum depth that will be crawled on the site specified in -url. This can stop the crawler from running for too long. The default value is -1, which specifies an unlimited crawl depth.

## Note
The fifth exercise makes use of the linkparser package developped in this exercise.
	