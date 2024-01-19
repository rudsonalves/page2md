# page2md

This is a terminal application to download html and convert it to md. It is focused on performing these HTML conversions on my page (jrblog.com.br).

I did this just to use goLang for development and training, replacing my shell scripts.

## Package converter

`import "page2md/converter"`

### Overview

### Index

func HtmlToMarkdown(html string) ([]byte, error)

### Package files

[converter.go](http://localhost:6060/src/page2md/converter/converter.go)

#### func HtmlToMarkdown

`func HtmlToMarkdown(html string)`

HtmlToMarkdown converts HTML content to Markdown format using the Pandoc tool.

## Package downloader

import "page2md/downloader"

### Overview

### Index

func DownloadPage(url string) (string, error)

### Package files

[downloader.go](http://localhost:6060/src/page2md/downloader/downloader.go)

#### func DownloadPage ¶

`func DownloadPage(url string) (string, error)`

DownloadPage fetches the content of the given URL and returns it as a string.

## Package markdown

`import "page2md/markdown"`

### Overview

### Index

`func ApplyFilters(inputFile string, outputFile string) error`
`func ApplyMarkdownFilters(inputString string, re1, re2, re3, re4, re5 *regexp.Regexp) string`

### Package files

[markdown.go](http://localhost:6060/src/page2md/markdown/markdown.go)

#### func ApplyFilters

`func ApplyFilters(inputFile string, outputFile string) error`

ApplyFilters reads the input file, applies markdown filters, and writes the result to the output file.

#### func ApplyMarkdownFilters

`func ApplyMarkdownFilters(inputString string, re1, re2, re3, re4, re5 *regexp.Regexp) string`

ApplyMarkdownFilters applies several filters to clean up and format the markdown text. It uses regular expressions to remove or replace certain patterns in the input string.

## Package utils

`import "page2md/utils"`

### Overview

### Index

`func GetLastPartOfURL(rawURL string) (string, error)`
`func IsValidURL(toTest string) bool`
`func Trim(s string) string`

### Package files

[utils.go](http://localhost:6060/src/page2md/utils/utils.go)

#### func GetLastPartOfURL

`func GetLastPartOfURL(rawURL string) (string, error)`

GetLastPartOfURL extrai a última parte da URL. Se a última parte terminar com '.html', essa extensão é removida.

#### func IsValidURL ¶

`func IsValidURL(toTest string) bool`

IsValidURL verifica se a string fornecida é uma URL válida.

#### func Trim

`func Trim(s string) string`

Trim removes all leading and trailing white spaces from a string.

This commit marks a significant advancement in the project's development, introducing new functionalities and enhancing the code structure. It encapsulates the creation of critical components, the structuring of utility functions, and the refinement of the application's control flow. Below are the detailed additions and improvements made:

- README.md:
  - Created the README.md file with basic instructions about the project.
- converter/converter.go:
  - Implemented the 'converter' package with the HtmlToMarkdown function, responsible for converting HTML to Markdown.
- downloader/downloader.go:
  - Introduced the 'downloader' package containing the DownloadPage function, dedicated to downloading web pages.
- main.go:
  - Refactored the main function, improving the control and structure of the application.
- markdown/markdown.go:
  - Developed the 'markdown' package which includes the functions compileRegex, ApplyFilters, and ApplyMarkdownFilters, aimed at processing Markdown text.
- utils/utils.go:
  - Created the 'utils' package with auxiliary functions, including Trim (for string handling), IsValidURL (for URL validation), and GetLastPartOfURL (for extracting specific segments from URLs).