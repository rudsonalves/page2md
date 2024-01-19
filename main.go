package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rudsonalves/page2md/converter"
	"github.com/rudsonalves/page2md/downloader"
	"github.com/rudsonalves/page2md/markdown"
	"github.com/rudsonalves/page2md/utils"
)

func main() {
	log.SetPrefix("INFO: ")
	log.SetFlags(log.Ldate | log.Ltime)

	log.Println("Starting application")

	// Check for the correct number of arguments
	if len(os.Args) < 2 {
		log.Println("No URL provided, exiting")
		fmt.Println("Usage: page2md <url> [outputFileName.md]")
		os.Exit(1)
	}
	urlInput := os.Args[1]

	if !utils.IsValidURL(urlInput) {
		log.Fatalf("Invalid URL provided: %s\n", urlInput)
	}

	// Determine the output file name
	var outputFile string
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	} else {
		var err error
		outputFile, err = utils.GetLastPartOfURL(urlInput)
		outputFile += ".md"
		if err != nil {
			outputFile = "output.md"
		}
	}

	// Download the HTML content of the page
	log.Printf("Downloading HTML content from %s\n", urlInput)
	html, err := downloader.DownloadPage(urlInput)
	if err != nil {
		log.Fatalf("An error occurred while downloading the page: %v\n", err)
		os.Exit(1)
	}

	// Convert the HTML content to Markdown
	log.Println("Converting HTML content to Markdown format")
	cmdOutput, err := converter.HtmlToMarkdown(html)
	if err != nil {
		log.Fatalf("An error occurred while converting HTML to Markdown: %v\n", err)
		os.Exit(1)
	}

	// Create a temporary file to store the intermediate markdown content
	log.Printf("Creating temporary file for intermediate markdown content")
	tmpFile, err := os.CreateTemp("", "temp-*.md")
	if err != nil {
		log.Fatalf("An error occurred while creating a temporary file: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = tmpFile.Write(cmdOutput)
	if err != nil {
		log.Fatalf("An error occurred while writing to the temporary file: %v\n", err)
		os.Exit(1)
	}

	// Apply filters to the markdown content and write the result to the output file
	log.Printf("Applying filters and writing the result to %s\n", outputFile)
	err = markdown.ApplyFilters(tmpFile.Name(), outputFile)
	if err != nil {
		log.Fatalf("An error occurred while applying filters: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Page converted to %s\n", outputFile)
}
