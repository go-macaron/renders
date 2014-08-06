package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/cnphpbb/renders"
)

var templates map[string]*template.Template

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates["index.html"].Execute(w, map[string]interface{}{"Name": "Joakim"}); err != nil {
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
		Funcs: template.FuncMap{
			"greet": func(name string) string {
				return fmt.Sprintf("Hello %s", name)
			},
	}
	templates, tmplErr = renders.LoadWithFuncMap(opt)
	if tmplErr != nil {
		panic(tmplErr)
	}
}
