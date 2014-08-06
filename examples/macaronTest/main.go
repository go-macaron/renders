package main

import (
	"github.com/Unknwon/macaron"

	"github.com/macaron-contrib/renders"
)

func main() {
	m := macaron.Classic()
	m.Use(renders.Renderer(
		renders.Options{
			Directory:  "templates",                // Specify what path to load the templates from.
			Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
			//Funcs:           FuncMap,                    // Specify helper function maps for templates to access.
			Charset:         "UTF-8",     // Sets encoding for json and html content-types. Default is "UTF-8".
			IndentJSON:      true,        // Output human readable JSON
			IndentXML:       true,        // Output human readable XML
			HTMLContentType: "text/html", // Output XHTML content type instead of default "text/html"
		}))
	m.Get("/", func(r renders.Render) {
		r.HTML(200, "pages/index.html", map[string]interface{}{"Title": "Home"})
	})
	m.Get("/profile", func(r renders.Render) {
		r.HTML(200, "pages/profile.tmpl", map[string]interface{}{"Title": "Profile"})
	})

	m.Get("/map", func(r renders.Render) {
		r.HTML(200, "pages/map.html", map[string]interface{}{"Title": "Map"})
	})
	m.Run()
}
