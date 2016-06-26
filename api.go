package main

import (
	"github.com/mauidude/go-readability"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"net/http"
)

func main() {
	InitConfig()

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Charset: "UTF-8",
	}))

	// This will set the Content-Type header to "text/html; charset=ISO-8859-1"
	m.Get("/", func(r render.Render) {
		r.Text(200, "hello world")
	})

	// This will set the Content-Type header to "text/plain; charset=ISO-8859-1"
	m.Get("/extract", func(r render.Render, req *http.Request) {
		url := req.URL.Query()["url"][0]
		response, err := http.Get(string(url))
		if err != nil {
			r.JSON(500, map[string]interface{}{
				"code": 500,
				"message": err.Error(),
			})
		} else {
			defer response.Body.Close()
			var by []byte
			var content string

			by, _ = ioutil.ReadAll(response.Body)
			html := string(by)
			doc, err := readability.NewDocument(html)
			if err != nil {
				content = ""
			} else {
				content = doc.Content()
			}

			title := GetTitle(html)

			r.JSON(200, map[string]interface{}{
				"content": content,
				"url": url,
				"title": title,
				"type": GetType(response),
			})
		}
	})

	m.Run()
}
