package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	res, err := http.Get("https://en.wikipedia.org/wiki/List_of_cities_in_Morocco")
	check(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	check(err)

	doc.Find(".wikitable tbody tr td:nth-child(2) a").Not("sup a").Each(func(i int, s *goquery.Selection) {
		fmt.Printf(s.Text() + "\n")
	})
}
