package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/cnphpbb/render"
)

var templates map[string]*template.Template

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates["pages/index.html"].Execute(w, map[string]interface{}{"Title": "Home"}); err != nil {
			log.Println(err)
		}
	})
	log.Println("web server listening at :8008")
	log.Fatal(http.ListenAndServe(":8008", nil))
}

func init() {
	var tmplErr error
	opt := &renders.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
	}
	if templates, tmplErr = render.Load("templates"); tmplErr != nil {
		panic(tmplErr)
	}
}
