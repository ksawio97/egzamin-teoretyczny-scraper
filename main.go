package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// c := colly.NewCollector()
	// c := colly.NewCollector()
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML(".question", func(e *colly.HTMLElement) {
		fmt.Println("Found ", e.ChildText(".title"))
	})
	url := "https://www.praktycznyegzamin.pl/inf04/teoria/wszystko/"
	c.Visit(url)

	fmt.Println("Scraping finished")
}
