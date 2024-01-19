package downloader

import (
	"io"
	"net/http"
)

// DownloadPage fetches the content of the given URL and returns it as a string.
func DownloadPage(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	htmlBytes, err := io.ReadAll(response.Body)
	return string(htmlBytes), err
}
