renders
=======

Macaron框架一款新的模板渲染中间件

Macaron 中间件处理器 渲染序列化的XML，JSON和HTML模板

## 安装
```go get github.com/macaron-contrib/renders```

## 例子
浏览 [examples](https://github.com/macaron-contrib/renders/tree/master/examples) 目录中有一些简单的例子

## Usage
使用 Go's [html/template](http://golang.org/pkg/html/template/) 包. 

~~~ go
// main.go
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

~~~


### 配置选项
`renders.Renderer` 包含多种配置选项:

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
使用标准模板与关键字， HTML文件的路径。
记得要通过上下文和尾随"."到模板中 (.}})。

index.html
	
	{{ template "templates/layouts/fullwidth.html" . }}

	{{ define "content" }}
	    content of index to be inserted into the fullwidth template
	{{ end }}

支持多层次的联接，如：
```index.html ---extends---> layouts/fullwidth.html ---extends---> base.html```

### Include
自动解析正确的文件路径 (这个是Macaron中没有的)

    {{ define "content" }}
        content of the fullwidth template
        {{ template "includes/widgets/signup.html" . }}
    {{ end }}

### 覆盖 define / default value

任何"define"在"template"下的扩展将覆盖以前的内容
还可以用来在{{ template }}中，定义默认值 

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



## Authors
* [cnphpbb](http://github.com/cnphpbb)
