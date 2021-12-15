package presentation

import (
	"fmt"
	"gopheradventures/model"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	defaultArc  string
	defaultHtml template.HTML = `
<html>
<head>{{.Title}}</head>
<br></br>
<body>{{block "list" .Story}}{{"\n"}}{{range .}}{{.}}{{end}}{{end}}</body>
<br></br>
{{block "list2" .Options}}{{"\n"}}{{range .}}<a href="/{{.Arc}}">{{.Text}}<a/><br></br>{{end}}{{end}}
</html>
`
)

func SetDefaultArc(arcName string) {
	defaultArc = arcName
}

func TemplateFlow(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/" {
		defaultArcPath := "/" + defaultArc
		http.Redirect(w, r, defaultArcPath, http.StatusPermanentRedirect)
		return
	}

	arcName := strings.TrimPrefix(r.RequestURI, "/")

	t, err := template.New("test").Parse(string(defaultHtml))
	if err != nil {
		errMsg := fmt.Sprintf("Error creating template 'test': %v", err.Error())
		w.Write([]byte(errMsg))
		log.Fatal(errMsg)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = t.ExecuteTemplate(w, "test", model.RuntimeArcs[arcName])
	if err != nil {
		errMsg := fmt.Sprintf("Error executing template 'test': %v", err.Error())
		w.Write([]byte(errMsg))
		log.Fatal(errMsg)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
