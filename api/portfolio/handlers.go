package portfolio

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"gopkg.in/yaml.v3"
)

type FrontMatter struct {
	Title       string   `yaml:"title"`
	Date        string   `yaml:"date"`
	Description string   `yaml:"description"`
}

func getPortfolio(w http.ResponseWriter, _ *http.Request) {
	temp := template.Must(template.ParseFiles("web/templates/index.html"))
	temp.Execute(w, nil)
}

func getHobbies(w http.ResponseWriter, _ *http.Request) {
	temp := template.Must(template.ParseFiles("web/templates/hobbies.html"))
	temp.Execute(w, nil)
}

func getWork(w http.ResponseWriter, _ *http.Request) {
	temp := template.Must(template.ParseFiles("web/templates/work.html"))
	temp.Execute(w, nil)
}

func getSchedule(w http.ResponseWriter, _ *http.Request) {
	temp := template.Must(template.ParseFiles("web/templates/schedule.html"))
	temp.Execute(w, nil)
}


func getJournals(w http.ResponseWriter, _ *http.Request) {

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

	temp := template.Must(template.ParseFiles("web/templates/journals.html"))
	temp.Execute(w, journals)

}

func getJournal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Journal: %s", chi.URLParam(r, "journal"))
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