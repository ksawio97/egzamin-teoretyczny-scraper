package options

import (
	"flag"
)

type Flags struct {
	Url            string // url to scrape
	Output         string // output directory
	RmTitlePrefix  bool   // title prefix should be removed
	RmAnswerPrefix bool   // answer prefix should be removed
}

func ReadFlags() Flags {
	url := flag.String("url", "https://www.praktycznyegzamin.pl/inf04/teoria/wszystko/", "Url from witch program will scrape data")
	output := flag.String("out", "_out", "Output folder witch will contain scraped data")
	rmTitlePrefix := flag.Bool("rmtitlep", false, "Whether title prefix '{number}. ' should be removed")
	rmAnswerPrefix := flag.Bool("rmanswerp", false, "Whether answer prefix '{[ABCD]. } should be removed'")
	flag.Parse()

	return Flags{
		Url:            *url,
		Output:         *output,
		RmTitlePrefix:  *rmTitlePrefix,
		RmAnswerPrefix: *rmAnswerPrefix,
	}
}
