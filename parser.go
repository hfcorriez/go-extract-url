package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"net/http"
)

var doc *goquery.Document

func InitDocument(r string) *goquery.Document {
	var err error
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(r))
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func GetTitle() string {
	return doc.Find("title").Text()
}

func GetDescription() string {
	metaSection := doc.Find("meta[name=description]")
	if (metaSection.Length() == 0) {
		return ""
	}

	value, exists := metaSection.Attr("content")
	if (!exists) {
		return ""
	}

	return value
}

func GetKeywords() []string {
	metaSection := doc.Find("meta[name=keywords]")
	if (metaSection.Length() == 0) {
		return make([]string, 0)
	}

	value, exists := metaSection.Attr("content")
	if (!exists) {
		return make([]string, 0)
	}

	return strings.Split(value, ",")
}

func GetImages() []string {
	return doc.Find("img[src]").Map(func(_ int, section *goquery.Selection) string {
		value, exists := section.Attr("src")
		if (!exists) {
			return ""
		} else {
			return value
		}
	})
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