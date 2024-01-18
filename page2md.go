package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"unicode/utf8"
)

func trim(s string) string {
	return strings.TrimSpace(s)
}

func compileRegex() (*regexp.Regexp, *regexp.Regexp, *regexp.Regexp, *regexp.Regexp, *regexp.Regexp) {
	re1 := regexp.MustCompile(`.*:::.*\n`)
	re2 := regexp.MustCompile(`{#[^}]*}`)
	re3 := regexp.MustCompile(` {.wp-block-code}`)
	re4 := regexp.MustCompile("```\\s*\\{.*\\}")
	re5 := regexp.MustCompile(`^\-+$`)
	return re1, re2, re3, re4, re5
}

func applyMarkdownFilters(inputString string, re1, re2, re3, re4, re5 *regexp.Regexp) string {
	inputString = re1.ReplaceAllString(inputString, "")
	inputString = re2.ReplaceAllString(inputString, "")
	inputString = re3.ReplaceAllString(inputString, "")
	inputString = re4.ReplaceAllString(inputString, "```")

	lines := strings.Split(inputString, "\n")
	filteredLines := []string{}
	previousLine := ""

	for _, line := range lines {
		trimmedLine := trim(line)
		if len(previousLine) > 0 && re5.MatchString(trimmedLine) && utf8.RuneCountInString(trimmedLine) == utf8.RuneCountInString(previousLine) {
			filteredLines[len(filteredLines)-1] = "## " + previousLine
		} else {
			filteredLines = append(filteredLines, line)
		}
		previousLine = trimmedLine
	}
	return strings.Join(filteredLines, "\n")
}

func applyFilters(inputFile string, outputFile string) error {
	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	re1, re2, re3, re4, re5 := compileRegex()
	inputString := applyMarkdownFilters(string(inputBytes), re1, re2, re3, re4, re5)

	err = os.WriteFile(outputFile, []byte(inputString), 0644)
	if err != nil {
		return err
	}

	return os.Chmod(outputFile, 0644)
}

func downloadPage(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	htmlBytes, err := io.ReadAll(response.Body)
	return string(htmlBytes), err
}

func htmlToMarkdown(html string) ([]byte, error) {
	markdown := strings.NewReader(html)
	cmd := exec.Command("pandoc", "-f", "html", "-t", "markdown")
	cmd.Stdin = markdown
	return cmd.Output()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <url> [output_file.md]")
		os.Exit(1)
	}
	url := os.Args[1]

	outputFile := "page.md"
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	html, err := downloadPage(url)
	if err != nil {
		fmt.Printf("An error occurred while downloading the page: %v\n", err)
		os.Exit(1)
	}

	cmdOutput, err := htmlToMarkdown(html)
	if err != nil {
		fmt.Printf("An error occurred while converting HTML to Markdown: %v\n", err)
		os.Exit(1)
	}

	tmpFile, err := os.CreateTemp("", "temp-*.md")
	if err != nil {
		fmt.Printf("An error occurred while creating a temporary file: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = tmpFile.Write(cmdOutput)
	if err != nil {
		fmt.Printf("An error occurred while writing to the temporary file: %v\n", err)
		os.Exit(1)
	}

	err = applyFilters(tmpFile.Name(), outputFile)
	if err != nil {
		fmt.Printf("An error occurred while applying filters: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Page converted to %s\n", outputFile)
}
