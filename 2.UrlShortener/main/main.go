package main

import (
	"flag"
	"fmt"
	"gophercises/urlshortener/urlshort"
	"net/http"
	"os"
	"path/filepath"
)

const (
	yamlDefault string = `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
)

var yamlFilePath string
var jsonFilePath string
var yaml []byte
var json []byte
var designatedHandler http.HandlerFunc
var sql bool

func init() {
	flag.StringVar(&yamlFilePath, "yamlfile", "none", "Path to the yaml file defining path + URL pairs for redirection. If left blank, will use default string yamlDefault in main.go.")
	flag.StringVar(&jsonFilePath, "jsonfile", "none", "Path to the json file defining path + URL pairs for redirection. Overrules the use of yaml files, or the saved yaml variable in main.go.")
	flag.BoolVar(&sql, "sql", false, "Specify this flag to use a PGSQL database connected to with env variable DATABASE_URL which has a table called pairs.")
	flag.Parse()
}

func main() {
	// Build the YAMLHandler using the mapHandler as the
	// fallback
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// presence of jsonfile parameter overrules use of yaml, sql parameter overrides all.
	// logically, only one flag should be presented.
	switch {
	case sql:
		sqlHandler := urlshort.PgSqlHandler(mapHandler)
		designatedHandler = sqlHandler
	case jsonFilePath != "none":
		json = readConfigFile(jsonFilePath)
		jsonHandler, err := urlshort.JSONHandler(json, mapHandler)
		if err != nil {
			panic(err)
		}

		designatedHandler = jsonHandler
	default:
		if yamlFilePath == "none" {
			yaml = []byte(yamlDefault)
		} else {
			yaml = readConfigFile(yamlFilePath)
		}

		yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
		if err != nil {
			panic(err)
		}

		designatedHandler = yamlHandler
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", designatedHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readConfigFile(path string) []byte {
	// Ensure file is present, load into a byte buffer and then convert to string.
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	f, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	return f
}
