package main

import "github.com/jansuthacheeva/honkboard/internal/models"

type templateData struct {
	Todo     models.Todo
	Todos    []models.Todo
	ListType string
}
