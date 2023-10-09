package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/amrikmalhans/go-htmx-portfolio.git/utils"
	"github.com/go-chi/chi/v5"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

type FrontMatter struct {
	Title       string `yaml:"title"`
	Date        string `yaml:"date"`
	Description string `yaml:"description"`
}

func GetJournals(w http.ResponseWriter, _ *http.Request) {

	files, err := os.ReadDir("./journals")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var journals []map[string]interface{}

	for _, file := range files {
		filePath := "./journals/" + file.Name()

		// Read and separate Front Matter and Markdown content
		frontMatterContent, _, err := readFileWithFrontMatter(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse Front Matter
		frontMatter, err := parseFrontMatter(frontMatterContent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		journal := map[string]interface{}{
			"Slug":        strings.TrimSuffix(file.Name(), ".md"),
			"FrontMatter": frontMatter,
		}

		journals = append(journals, journal)
	}

	if err := utils.Temp.ExecuteTemplate(w, "base.html", journals); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func GetJournal(w http.ResponseWriter, r *http.Request) {

	journalSlug := chi.URLParam(r, "journal")
	filePath := "./journals/" + journalSlug + ".md"

	frontMatterContent, markdownContent, err := readFileWithFrontMatter(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	frontMatter, err := parseFrontMatter(frontMatterContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert Markdown to HTML using Goldmark
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM), // GitHub-flavored markdown
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // Enable rendering of raw HTML
		),
	)

	if err := md.Convert([]byte(markdownContent), &buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	htmlContent := buf.String()

	journal := map[string]interface{}{
		"Slug":        journalSlug,
		"FrontMatter": frontMatter,
		"Content":     template.HTML(htmlContent),
	}

	if err := utils.Temp.ExecuteTemplate(w, "base.html", journal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func readFileWithFrontMatter(path string) (string, string, error) {
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}

	content := string(contentBytes)
	parts := strings.SplitN(content, "---\n", 3)
	if len(parts) < 3 {
		return "", "", fmt.Errorf("front Matter not found")
	}

	return parts[1], parts[2], nil
}

func parseFrontMatter(frontMatterContent string) (FrontMatter, error) {
	var fm FrontMatter
	err := yaml.Unmarshal([]byte(frontMatterContent), &fm)
	if err != nil {
		return FrontMatter{}, err
	}

	return fm, nil
}
