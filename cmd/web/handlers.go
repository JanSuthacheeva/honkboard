package main

import (
	"html/template"
	"net/http"

	"github.com/jansuthacheeva/honkboard/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	listType := app.sessionManager.GetString(r.Context(), "list-type")

	var todos []models.Todo
	var err error
	if listType == "" {
		todos, err = app.todos.GetAll("Personal")
	} else {
		todos, err = app.todos.GetAll(listType)
	}
	if err != nil {
		app.serverError(w, r, err)
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/index.html",
		"./ui/html/partials/todo-list.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Todos:    todos,
		ListType: listType,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) showPersonalTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.todos.GetAll("Personal")
	if err != nil {
		app.serverError(w, r, err)
	}

	files := []string{
		"./ui/html/pages/index.html",
		"./ui/html/partials/todo-list.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Todos:    todos,
		ListType: "Personal",
	}

	app.sessionManager.Put(r.Context(), "list-type", "Personal")

	err = ts.ExecuteTemplate(w, "todo-list", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) showProfessionalTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.todos.GetAll("Professional")
	if err != nil {
		app.serverError(w, r, err)
	}

	files := []string{
		"./ui/html/pages/index.html",
		"./ui/html/partials/todo-list.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Todos:    todos,
		ListType: "Professional",
	}

	app.sessionManager.Put(r.Context(), "list-type", "Professional")

	err = ts.ExecuteTemplate(w, "todo-list", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
