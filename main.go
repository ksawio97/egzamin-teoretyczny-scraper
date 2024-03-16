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

type Question struct {
	Title        string    `json:"title"`
	Answers      [4]string `json:"answers"`
	CorrectIndex int       `json:"correctIndex"`
	ImagePath    string    `json:"imagePath,omitempty"`
	VideoPath    string    `json:"videoPath,omitempty"`
}

func main() {
	url := "https://www.praktycznyegzamin.pl/inf04/teoria/wszystko/"

	// create output files
	outputDir := "_out"
	imagesPath := outputDir + "/images"
	videosPath := outputDir + "/videos"

	err := createFoldersStructure(outputDir, imagesPath, videosPath)
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

		if correctIndex == -1 {
			return
		}

		imgName, err := saveMedia(client, url, e.ChildAttr("img", "src"), imagesPath)
		if err != nil {
			fmt.Printf("Error saving image from url: %v", err)
			return
		}

		videoPath, err := saveMedia(client, url, e.ChildAttr("source", "src"), videosPath)
		if err != nil {
			fmt.Printf("Error saving video from url: %v", err)
			return
		}

		questions = append(questions, Question{
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
		f.Write(data)
	})

	c.Visit(url)
}

func saveMedia(client *http.Client, url, htmlSrc, saveFolder string) (string, error) {
	name := ""
	if htmlSrc != "" {
		// parse file name from imgsrc
		pattern, _ := regexp.Compile(`[^/]+$`)
		name = string(pattern.Find([]byte(htmlSrc)))
		if name == "" {
			return "", fmt.Errorf("failed parsing file name from %v", htmlSrc)
		}

		mediaPath := saveFolder + "/" + name
		err := saveFileFromUrl(client, fmt.Sprintf("%v/%v", url, htmlSrc), mediaPath)
		if err != nil {
			return "", fmt.Errorf("failed saving file to %v from url %v: %v", mediaPath, url, err)
		}
	}

	return name, nil
}

func createFoldersStructure(dirs ...string) error {
	for i := range dirs {
		err := os.MkdirAll(dirs[i], 0777)
		if err != nil {
			return fmt.Errorf("failed creating %v: %v", dirs[i], err)
		}
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
