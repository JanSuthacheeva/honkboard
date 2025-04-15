package main

import (
	"html/template"
	"path/filepath"

	"github.com/jansuthacheeva/honkboard/internal/models"
)

type templateData struct {
	Todo     models.Todo
	Todos    []models.Todo
	ListType string
	Errors   []string
	Form     any
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		partials, err := filepath.Glob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}
		files := []string{
			"./ui/html/base.html",
			page,
		}
		files = append(files, partials...)

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
