package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"net/http"
)

func GetTitle(r string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(r))
	if err != nil {
		log.Fatal(err)
	}

	return doc.Find("title").Text()
}

func GetType(r *http.Response) string {
	mimeType := strings.Split(r.Header.Get("Content-Type"), ";")[0]
	contentType := ""
	for k, v := range Config("contentTypes").Obj {
		if (mimeType == v.Str) {
			contentType = strings.Split(k, ".")[0]
		}
	}

	return contentType
}