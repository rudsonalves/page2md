package markdown

import (
	"os"
	"page2md/utils"
	"regexp"
	"strings"
	"unicode/utf8"
)

// compileRegex compiles all regular expressions used in applyMarkdownFilters.
// Returns pointers to the compiled regular expressions.
func compileRegex() (*regexp.Regexp, *regexp.Regexp, *regexp.Regexp, *regexp.Regexp, *regexp.Regexp) {
	re1 := regexp.MustCompile(`.*:::.*\n`)
	re2 := regexp.MustCompile(`{#[^}]*}`)
	re3 := regexp.MustCompile(` {.wp-block-code}`)
	re4 := regexp.MustCompile("```\\s*\\{.*\\}")
	re5 := regexp.MustCompile(`^\-+$`)
	return re1, re2, re3, re4, re5
}

// ApplyFilters reads the input file, applies markdown filters, and writes the result to the output file.
func ApplyFilters(inputFile string, outputFile string) error {
	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	re1, re2, re3, re4, re5 := compileRegex()
	inputString := ApplyMarkdownFilters(string(inputBytes), re1, re2, re3, re4, re5)

	err = os.WriteFile(outputFile, []byte(inputString), 0644)
	if err != nil {
		return err
	}

	// Explicitly set file permissions for readability by all users
	return os.Chmod(outputFile, 0644)
}

// ApplyMarkdownFilters applies several filters to clean up and format the markdown text.
// It uses regular expressions to remove or replace certain patterns in the input string.
func ApplyMarkdownFilters(inputString string, re1, re2, re3, re4, re5 *regexp.Regexp) string {
	// Apply regex-based replacements
	inputString = re1.ReplaceAllString(inputString, "")
	inputString = re2.ReplaceAllString(inputString, "")
	inputString = re3.ReplaceAllString(inputString, "")
	inputString = re4.ReplaceAllString(inputString, "```")

	// Process the lines of the markdown content
	lines := strings.Split(inputString, "\n")
	filteredLines := []string{}
	previousLine := ""

	for _, line := range lines {
		trimmedLine := utils.Trim(line)
		// Check if the current line and the previous line form a specific pattern
		if len(previousLine) > 0 && re5.MatchString(trimmedLine) && utf8.RuneCountInString(trimmedLine) == utf8.RuneCountInString(previousLine) {
			// Modify the previous line in the filtered output
			filteredLines[len(filteredLines)-1] = "## " + previousLine
		} else {
			// Add the current line to the filtered output
			filteredLines = append(filteredLines, line)
		}
		previousLine = trimmedLine
	}
	return strings.Join(filteredLines, "\n")
}
