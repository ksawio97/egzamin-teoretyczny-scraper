package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/gocolly/colly/v2"
)

// TODO add saving videos
type Question struct {
	Title        string    `json:"title"`
	Answers      [4]string `json:"answers"`
	CorrectIndex int       `json:"correctIndex"`
	ImagePath    string    `json:"imagePath,omitempty"`
}

func main() {
	url := "https://www.praktycznyegzamin.pl/inf04/teoria/wszystko/"

	// create output files
	outputDir := "_out"
	imagesPath := outputDir + "/images"
	err := createFoldersStructure(outputDir, imagesPath)
	if err != nil {
		fmt.Printf("failed to create folders structure: %v", err)
		return
	}

	f, err := os.Create(outputDir + "/questions.json")
	if err != nil {
		fmt.Printf("failed to create questions output file: %v", err)
		return
	}

	c := colly.NewCollector()
	client := &http.Client{}

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

		imgSrc := e.ChildAttr("img", "src")
		imgName := ""
		if imgSrc != "" {
			// parse file name from imgsrc
			pattern, _ := regexp.Compile(`[^/]+$`)
			imgName = string(pattern.Find([]byte(imgSrc)))
			if imgName == "" {
				fmt.Printf("failed parsing file name from %v", imgSrc)
				return
			}

			imgPath := imagesPath + "/" + string(imgName)
			err := saveFileFromUrl(client, fmt.Sprintf("%v/%v", url, imgSrc), imgPath)
			if err != nil {
				fmt.Println(err)
				imgName = ""
			}
		}

		if correctIndex == -1 {
			return
		}

		questions = append(questions, Question{
			Title:        title,
			Answers:      answers,
			CorrectIndex: correctIndex,
			ImagePath:    imgName,
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

	c.Visit(url)
}

func createFoldersStructure(outputDir, imagesPath string) error {
	// Create file where questions will be saved
	err := os.MkdirAll(outputDir, 0777)
	if err != nil {
		return fmt.Errorf("failed creating %v file: %v", outputDir, err)
	}

	// Create dir where images will be saved
	err = os.MkdirAll(imagesPath, 0777)
	if err != nil {
		return fmt.Errorf("failed creating %v directory: %v", imagesPath, err)
	}
	return nil
}

func saveFileFromUrl(client *http.Client, url, path string) error {
	body, err := getFileFromUrl(client, url)
	if err != nil {
		return fmt.Errorf("failed to get file from url %v: %v", url, err)
	}
	defer body.Close()

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	if err != nil {
		return fmt.Errorf("failed to copy body to new file %v: %v", path, err)
	}
	return nil
}

func getFileFromUrl(client *http.Client, url string) (io.ReadCloser, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusOK {
		return response.Body, nil
	}
	response.Body.Close()
	return nil, fmt.Errorf("failed to get file: http status code %v", response.StatusCode)
}
