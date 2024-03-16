package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func SaveMedia(client *http.Client, url, htmlSrc, saveFolder string) (string, error) {
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

func saveFileFromUrl(client *http.Client, url, path string) error {
	body, err := fileFromUrl(client, url)
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

func fileFromUrl(client *http.Client, url string) (io.ReadCloser, error) {
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
