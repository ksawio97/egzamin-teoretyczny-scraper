package options

import (
	egzaminteoretycznyscraper "egzamin-teoretyczny-scraper/internal/egzamin-teoretyczny-scraper"
)

func Options() egzaminteoretycznyscraper.Options {
	flags := ReadFlags()

	return egzaminteoretycznyscraper.Options{
		Url:            flags.Url,
		Output:         flags.Output,
		Videos:         flags.Output + "/videos",
		Images:         flags.Output + "/images",
		RmTitlePrefix:  flags.RmTitlePrefix,
		RmAnswerPrefix: flags.RmAnswerPrefix,
	}
}
