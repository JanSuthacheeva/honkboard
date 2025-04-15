package main

import (
	"errors"
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

	data := templateData{
		Todos:    todos,
		ListType: listType,
	}

	app.render(w, r, http.StatusOK, "index.html", "base", data)
}

func (app *application) showPersonalTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.todos.GetAll("Personal")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Todos:    todos,
		ListType: "Personal",
	}

	app.sessionManager.Put(r.Context(), "list-type", "Personal")

	app.render(w, r, http.StatusOK, "index.html", "todo-list", data)
}

func (app *application) showProfessionalTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.todos.GetAll("Professional")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Todos:    todos,
		ListType: "Professional",
	}

	app.sessionManager.Put(r.Context(), "list-type", "Professional")

	app.render(w, r, http.StatusOK, "index.html", "todo-list", data)
}

func (app *application) createTodo(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteTodo(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	var err error

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
	listType := app.sessionManager.GetString(r.Context(), "list-type")
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

	data := templateData{
		Todos:    todos,
		ListType: listType,
	}

	app.render(w, r, http.StatusOK, "index.html", "todo-list", data)

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

	data := templateData{
		Todos:    todos,
		ListType: listType,
	}

	app.render(w, r, http.StatusOK, "index.html", "todo-list", data)
}
