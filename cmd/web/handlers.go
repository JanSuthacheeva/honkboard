package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	todos, err := app.todos.GetAll("personal")
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
		Todos: todos,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) showPersonalTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.todos.GetAll("personal")
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
		Todos: todos,
	}

	err = ts.ExecuteTemplate(w, "todo-list", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) showProfessionalTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.todos.GetAll("professional")
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
		Todos: todos,
	}

	err = ts.ExecuteTemplate(w, "todo-list", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
