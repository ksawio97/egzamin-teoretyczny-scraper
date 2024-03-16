package egzaminteoretycznyscraper

type Question struct {
	Title        string    `json:"title"`
	Answers      [4]string `json:"answers"`
	CorrectIndex int       `json:"correctIndex"`
	ImagePath    string    `json:"imagePath,omitempty"`
	VideoPath    string    `json:"videoPath,omitempty"`
}

type Options struct {
	Url    string // url of website to scrape
	Output string // directory for ouput files
	Videos string // directory for videos
	Images string // directory for images
}
