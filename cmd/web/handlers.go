package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jansuthacheeva/honkboard/internal/models"
	"github.com/jansuthacheeva/honkboard/internal/validator"
)

type createTodoForm struct {
	Title string
	validator.Validator
}

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
		Form:     createTodoForm{},
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
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	listType := app.sessionManager.GetString(r.Context(), "list-type")

	form := createTodoForm{
		Title: r.PostForm.Get("title"),
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 80), "title", "This field cannot be more than 80 characters long")

	if !form.Valid() {
		todos, err := app.todos.GetAll(listType)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		data := templateData{
			Todos:    todos,
			Form:     form,
			ListType: listType,
		}
		app.render(w, r, http.StatusUnprocessableEntity, "index.html", "main", data)
		return
	}

	_, err = app.todos.Insert(form.Title, listType)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	todos, err := app.todos.GetAll(listType)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Todos:    todos,
		ListType: listType,
		Form:     createTodoForm{},
	}

	app.render(w, r, http.StatusCreated, "index.html", "main", data)
}

func (app *application) toggleTodoStatus(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	if id <= 0 {
		app.notFound(w, r)
		return
	}
	todo, err := app.todos.ToggleStatus(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	listType := app.sessionManager.GetString(r.Context(), "list-type")
	if listType == "" {
		app.sessionManager.Put(r.Context(), "list-type", "Personal")
	}

	todos, err := app.todos.GetAll(listType)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		ListType: listType,
		ID:       todo.ID,
		Title:    todo.Title,
		Status:   todo.Status.String(),
		Todos:    todos,
	}

	app.render(w, r, http.StatusOK, "index.html", "todo-row-swap", data)
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

	data := templateData{
		ListType: listType,
	}

	err := app.todos.DeleteCompleted(listType)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			data.Errors = []string{
				"no completed tasks to delete",
			}
		default:
			app.serverError(w, r, err)
			return
		}
	}
	todos, err := app.todos.GetAll(listType)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data.Todos = todos

	if data.Errors != nil {
		app.render(w, r, http.StatusUnprocessableEntity, "index.html", "todo-list", data)
	} else {

		app.render(w, r, http.StatusOK, "index.html", "todo-list", data)
	}
}
