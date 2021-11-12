# URL Shortener; Gophercise number 2
This project is meant to be a URL shortener, or more accurately, a URL redirector. It can take config files that are json,
yaml, or hard-coded variables that will build a mapping between pathes on the web server and URLs. 

Note: the yaml files must be in the same format as the testyaml.yml file (only path entries with url subentries, all below a pairs header).
Ditto for json; look at the test.json file for the pattern to follow.

TODO: make it work with a database

TBH this could be extended fairly easily into a url shortener web-page by adding functionality to generate URLs and then hosting it in a container on some cloud service. 