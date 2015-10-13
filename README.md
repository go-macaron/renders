renders
=======

New template render of macaron framework

Macaron middleware/handler for rendering serialized JSON, XML, and HTML template responses.

## Installation
```go get github.com/go-macaron/renders```

## Examples
Check out the [examples](https://github.com/go-macaron/renders/tree/master/examples) folder for some examples

## Usage
render uses Go's [html/template](http://golang.org/pkg/html/template/) package to render html templates.

~~~ go
// main.go
package main

import (
	"gopkg.in/macaron.v1"

	"github.com/go-macaron/renders"
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

~~~


### Options
`renders.Renderer` comes with a variety of configuration options:

~~~ go
// ...
m.Use(renders.Renderer(renders.Options{
  Directory: "templates", // Specify what path to load the templates from.
  Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
  //Funcs: template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
  Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
  IndentJSON: true, // Output human readable JSON
  IndentXML: true, // Output human readable XML
  HTMLContentType: "text/html", // Output XHTML content type instead of default "text/html"
}))
// ...
~~~

### Extends
Just use the standard template keyword with a *.html file path.
REMEMBER to pass the context to the parent template with the trailing dot (. }}).

index.html
	
	{{ template "templates/layouts/fullwidth.html" . }}

	{{ define "content" }}
	    content of index to be inserted into the fullwidth template
	{{ end }}

This will also work with multi-level support, e.g. 
```index.html ---extends---> layouts/fullwidth.html ---extends---> base.html```

### Include
Automatically parse the right file just by writing the path to it

    {{ define "content" }}
        content of the fullwidth template
        {{ template "includes/widgets/signup.html" . }}
    {{ end }}

### Overwriting define / default value
Any "define" of the same "template" down the extend chain will overwrite the former content
This can be used to define default values for a {{ template }} like so

base.html

    <!DOCTYPE html>
	<html>
	  <head>
	    <title>{{ template "title" }}</title>
	  </head>
	</html>

	{{ define "title" }}Default Title{{ end }}

profile.html

    {{ template "templates/base.html" . }}
    {{ define "title" }}Hello World{{ end }}

This would produce panic in std lib parsing but now it works by simply renaming the define's further down the chain not to interrupt the most specific one.


## Authors
* [cnphpbb](http://github.com/cnphpbb)
