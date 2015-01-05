package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	kyodoNewsURL := "http://prw.kyodonews.jp/opn/category/release/kokunai/kokunai020/kokunai020-0150/"

	doc, err := goquery.NewDocument(kyodoNewsURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Println(url)
	})
}
