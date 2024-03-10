package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

type Question struct {
	Title        string    `json:"title"`
	Answers      [4]string `json:"answers"`
	CorrectIndex int       `json:"correctIndex"`
}

func main() {
	c := colly.NewCollector()
	questions := []Question{}
	// Where questions will be saved
	f, err := os.Create("_out/questions.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

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
			Title:        title,
			Answers:      answers,
			CorrectIndex: correctIndex,
		})
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("Scrape completed, found %v questions.\n", len(questions))

		data, err := json.Marshal(questions)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Write(data)
	})

	url := "https://www.praktycznyegzamin.pl/inf04/teoria/wszystko/"
	c.Visit(url)
}
