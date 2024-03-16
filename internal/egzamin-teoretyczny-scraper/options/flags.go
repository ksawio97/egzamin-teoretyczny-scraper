package options

import (
	"flag"
)

type Flags struct {
	Url    string // url to scrape
	Output string // output directory
}

func ReadFlags() Flags {
	url := flag.String("url", "https://www.praktycznyegzamin.pl/inf04/teoria/wszystko/", "Url from witch program will scrape data")
	output := flag.String("out", "_out", "Output folder witch will contain scraped data")
	flag.Parse()

	return Flags{
		Url:    *url,
		Output: *output,
	}
}
