package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jansuthacheeva/honkboard/internal/models"
	"github.com/jansuthacheeva/honkboard/internal/validator"
)

type createTodoForm struct {
	Title               string `form:"title"`
	validator.Validator `form:"-"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	listType := app.sessionManager.GetString(r.Context(), "list-type")
	id := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

	var todos []models.Todo
	var err error
	if listType == "" {
		app.sessionManager.Put(r.Context(), "list-type", "Personal")
		todos, err = app.todos.GetAll(id, "Personal")
	} else {
		todos, err = app.todos.GetAll(id, listType)
	}
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)

	data.Todos = todos
	data.ListType = listType
	data.Form = createTodoForm{}

	app.render(w, r, http.StatusOK, "index.html", "base", data)
}

func (app *application) showPersonalTodos(w http.ResponseWriter, r *http.Request) {
	id := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
	todos, err := app.todos.GetAll(id, "Personal")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Todos = todos
	data.ListType = "Personal"

	app.sessionManager.Put(r.Context(), "list-type", "Personal")

	app.render(w, r, http.StatusOK, "index.html", "todo-list", data)
}

func (app *application) showProfessionalTodos(w http.ResponseWriter, r *http.Request) {
	id := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
	todos, err := app.todos.GetAll(id, "Professional")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Todos = todos
	data.ListType = "Professional"

	app.sessionManager.Put(r.Context(), "list-type", "Professional")

	app.render(w, r, http.StatusOK, "index.html", "todo-list", data)
}

func (app *application) createTodo(w http.ResponseWriter, r *http.Request) {

	var form createTodoForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	listType := app.sessionManager.GetString(r.Context(), "list-type")
	id := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 80), "title", "This field cannot be more than 80 characters long")

	data := app.newTemplateData(r)

	if !form.Valid() {
		todos, err := app.todos.GetAll(id, listType)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		data.Todos = todos
		data.Form = form
		data.ListType = listType

		app.render(w, r, http.StatusUnprocessableEntity, "index.html", "main", data)
		return
	}

	_, err = app.todos.Insert(id, form.Title, listType)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrMaxTodos):
			form.AddFieldError("title", "Maximum number of todos reached for this list")
			todos, err := app.todos.GetAll(id, listType)
			if err != nil {
				app.serverError(w, r, err)
				return
			}
			data.Todos = todos
			data.Form = form
			data.ListType = listType
			app.render(w, r, http.StatusUnprocessableEntity, "index.html", "main", data)
			return
		default:
			app.serverError(w, r, err)
			return
		}
	}

	todos, err := app.todos.GetAll(id, listType)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data.Todos = todos
	data.Form = createTodoForm{}
	data.ListType = listType

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
	userId := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
	if userId == 0 {
		app.serverError(w, r, models.ErrUnknownAuth)
		return
	}
	todo, err := app.todos.ToggleStatus(userId, id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	listType := app.sessionManager.GetString(r.Context(), "list-type")
	if listType == "" {
		app.sessionManager.Put(r.Context(), "list-type", "Personal")
	}

	todos, err := app.todos.GetAll(userId, listType)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.ListType = listType
	data.ID = todo.ID
	data.Title = todo.Title
	data.Status = todo.Status.String()
	data.Todos = todos

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
	userId := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
	if userId == 0 {
		app.serverError(w, r, models.ErrUnknownAuth)
		return
	}

	err = app.todos.Delete(userId, id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	listType := app.sessionManager.GetString(r.Context(), "list-type")
	if listType == "" {
		app.sessionManager.Put(r.Context(), "list-type", "Personal")
		todos, err = app.todos.GetAll(userId, "Personal")
	} else {
		todos, err = app.todos.GetAll(userId, listType)
	}
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Todos = todos
	data.ListType = listType

	app.render(w, r, http.StatusOK, "index.html", "todo-list", data)
}

func (app *application) deleteCompletedTodos(w http.ResponseWriter, r *http.Request) {
	listType := app.sessionManager.GetString(r.Context(), "list-type")
	if listType == "" {
		app.serverError(w, r, errors.New("session: list-type not found."))
		return
	}
	userId := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
	if userId == 0 {
		app.serverError(w, r, models.ErrUnknownAuth)
		return
	}

	data := app.newTemplateData(r)
	data.ListType = listType

	err := app.todos.DeleteCompleted(userId, listType)
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
	todos, err := app.todos.GetAll(userId, listType)
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
