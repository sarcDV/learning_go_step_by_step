package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func main() {
	// Define the URL to scrape
	url := "https://golang.org/doc/install"

	// Create a new Colly collector
	c := colly.NewCollector()

	// On request response, extract title from the page
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		// Parse the HTML document
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		title := doc.Find("title").Text()
		fmt.Println("Title:", title)
	})

	// Visit the URL
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
