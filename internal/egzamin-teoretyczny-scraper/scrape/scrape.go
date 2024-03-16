package scrape

import (
	egzaminteoretycznyscraper "egzamin-teoretyczny-scraper/internal/egzamin-teoretyczny-scraper"
	"egzamin-teoretyczny-scraper/internal/egzamin-teoretyczny-scraper/files"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/gocolly/colly"
)

var titlePrefixPattern = regexp.MustCompile(`^(\d+)\. `)

// it scrapes data from website and saves everything in defined directory
func Scrape(questionsFile *os.File, options egzaminteoretycznyscraper.Options) {
	c := colly.NewCollector()
	client := &http.Client{}

	questions := []egzaminteoretycznyscraper.Question{}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML(".question", func(e *colly.HTMLElement) {
		title := e.ChildText(".title")
		if options.RmTitlePrefix {
			title = removeTitlePrefix(title)
		}

		// expected only 4 answers
		answers := [4]string{}
		correctIndex := -1

		e.ForEach("div.answer", func(i int, h *colly.HTMLElement) {
			if i < len(answers) {
				if h.DOM.HasClass("correct") {
					correctIndex = i
				}
				answers[i] = h.Text
				if options.RmAnswerPrefix {
					answers[i] = removeAnswerPrefix(answers[i])
				}
			}
		})

		if correctIndex == -1 {
			return
		}

		imgName, err := files.SaveMedia(client, options.Url, e.ChildAttr("img", "src"), options.Images)
		if err != nil {
			fmt.Printf("Error saving image from url: %v", err)
			return
		}

		videoPath, err := files.SaveMedia(client, options.Url, e.ChildAttr("source", "src"), options.Videos)
		if err != nil {
			fmt.Printf("Error saving video from url: %v", err)
			return
		}

		questions = append(questions, egzaminteoretycznyscraper.Question{
			Title:        title,
			Answers:      answers,
			CorrectIndex: correctIndex,
			ImagePath:    imgName,
			VideoPath:    videoPath,
		})
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("Scrape completed, found %v questions.\n", len(questions))

		data, err := json.Marshal(questions)
		if err != nil {
			fmt.Println(err)
			return
		}
		questionsFile.Write(data)
	})

	c.Visit(options.Url)
}

// removes question number from title
func removeTitlePrefix(title string) string {
	return titlePrefixPattern.ReplaceAllLiteralString(title, "")
}

// removes answer prefix ([ABCD]. )
func removeAnswerPrefix(answer string) string {
	return answer[3:]
}
