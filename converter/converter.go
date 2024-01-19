package converter

import (
	"os/exec"
	"strings"
)

// HtmlToMarkdown converts HTML content to Markdown format using the Pandoc tool.
func HtmlToMarkdown(html string) ([]byte, error) {
	markdown := strings.NewReader(html)
	cmd := exec.Command("pandoc", "-f", "html", "-t", "markdown")
	cmd.Stdin = markdown
	return cmd.Output()
}
