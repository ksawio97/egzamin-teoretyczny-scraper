package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Question struct {
	title        string
	answers      [4]string
	correctIndex int
}

func main() {
	c := colly.NewCollector()
	questions := []Question{}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML(".question", func(e *colly.HTMLElement) {
		title := e.ChildText(".title")
		// expected only 4 answers
		answers := [4]string{}
		correctIndex := -1

		e.ForEach("div.answer", func(i int, h *colly.HTMLElement) {
			if i < len(answers) {
				if h.DOM.HasClass("correct") {
					correctIndex = i
				}
				answers[i] = h.Text
			}
		})

		if correctIndex == -1 {
			return
		}

		questions = append(questions, Question{
			title:        title,
			answers:      answers,
			correctIndex: correctIndex,
		})
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("Scrape completed, found %v questions.\n", len(questions))
	})

	url := "https://www.praktycznyegzamin.pl/inf04/teoria/wszystko/"
	c.Visit(url)
}
