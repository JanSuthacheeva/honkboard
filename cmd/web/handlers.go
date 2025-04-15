package main

import (
	"errors"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jansuthacheeva/honkboard/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	listType := app.sessionManager.GetString(r.Context(), "list-type")

	var todos []models.Todo
	var err error
	if listType == "" {
		app.sessionManager.Put(r.Context(), "list-type", "Personal")
		todos, err = app.todos.GetAll("Personal")
	} else {
		todos, err = app.todos.GetAll(listType)
	}
	if err != nil {
		app.serverError(w, r, err)
		return
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
		return
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
		return
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

func (app *application) deleteTodo(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	if id <= 0 {
		app.notFound(w, r)
		return
	}

	err = app.todos.Delete(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) deleteCompletedTodos(w http.ResponseWriter, r *http.Request) {
	listType := app.sessionManager.GetString(r.Context(), "list-type")
	if listType == "" {
		app.serverError(w, r, errors.New("session: list-type not found."))
		return
	}

	err := app.todos.DeleteCompleted(listType)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	todos, err := app.todos.GetAll(listType)
	if err != nil {
		// switch {
		// case errors.Is(err, models.ErrNoRecord) {
		// 422 or something that htmx can use
		// }
		// }
		app.serverError(w, r, err)
		return
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
		ListType: listType,
	}

	err = ts.ExecuteTemplate(w, "todo-list", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
