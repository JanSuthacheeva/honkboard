package main

import (
	"html/template"
	"path/filepath"

	"github.com/jansuthacheeva/honkboard/internal/models"
)

type templateData struct {
	Todos              []models.Todo
	ID                 int
	Title              string
	Status             string
	ListType           string
	Errors             []string
	Form               any
	User               *models.User
	IsAuthenticated    bool
	CSRFToken          string
	BaseURL            string
	PasswordResetToken models.PasswordResetToken
}

func countDoneTodos(todos []models.Todo) int {
	count := 0
	for _, todo := range todos {
		if todo.Status == "done" {
			count++
		}
	}
	return count
}

var functions = template.FuncMap{
	"countDoneTodos": countDoneTodos,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
