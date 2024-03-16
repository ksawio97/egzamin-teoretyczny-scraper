package main

import (
	"egzamin-teoretyczny-scraper/internal/egzamin-teoretyczny-scraper/files"
	"egzamin-teoretyczny-scraper/internal/egzamin-teoretyczny-scraper/options"
	"egzamin-teoretyczny-scraper/internal/egzamin-teoretyczny-scraper/scrape"
	"fmt"
)

func main() {
	options := options.Options()
	questionsFile, err := files.SetupStructure(options)
	if err != nil {
		fmt.Printf("Error creating files structure: %v", err)
		return
	}
	scrape.Scrape(questionsFile, options)
}
