package utils

import (
	"net/url"
	"path"
	"strings"
)

// Trim removes all leading and trailing white spaces from a string.
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// IsValidURL verifica se a string fornecida é uma URL válida.
func IsValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

// GetLastPartOfURL extrai a última parte da URL.
// Se a última parte terminar com '.html', essa extensão é removida.
func GetLastPartOfURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// gate the last part from path
	lastPart := path.Base(parsedURL.Path)

	// remove '.html'
	lastPart = strings.TrimSuffix(lastPart, ".html")

	return lastPart, nil
}
