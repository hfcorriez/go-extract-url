package main

import (
	"github.com/mauidude/go-readability"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
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
			log.Fatal(err)
		} else {
			defer response.Body.Close()
			var by []byte
			by, _ = ioutil.ReadAll(response.Body)
			doc, err := readability.NewDocument(string(by))
			if err != nil {
				// do something ...
			}
			r.JSON(200, map[string]interface{}{"content": doc.Content(), "url": url})
		}
	})

	m.Run()
}
