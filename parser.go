package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func GetTitle(r string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(r))
	if err != nil {
		log.Fatal(err)
	}

	return doc.Find("title").Text()
}