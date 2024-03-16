package files

import (
	egzaminteoretycznyscraper "egzamin-teoretyczny-scraper/internal/egzamin-teoretyczny-scraper"
	"fmt"
	"os"
)

// creates folders structure witch will be used for saving questions data
func SetupStructure(options egzaminteoretycznyscraper.Options) (*os.File, error) {
	err := createDirs(options.Output, options.Images, options.Videos)
	if err != nil {
		return nil, fmt.Errorf("failed to create folders structure: %v", err)
	}

	f, err := os.Create(options.Output + "/questions.json")
	if err != nil {
		return nil, fmt.Errorf("failed to create questions output file: %v", err)
	}

	return f, nil
}

// creates multiple empty directories
func createDirs(dirs ...string) error {
	for i := range dirs {
		err := os.MkdirAll(dirs[i], 0777)
		if err != nil {
			return fmt.Errorf("failed creating %v: %v", dirs[i], err)
		}
	}

	return nil
}
