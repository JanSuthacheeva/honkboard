package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	fileServer := http.FileServer(http.Dir("./static/"))

	router := http.NewServeMux()

	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf)
	// Users
	router.Handle(http.MethodGet+" /login", dynamic.ThenFunc(app.showLoginForm))
	router.Handle(http.MethodPost+" /sessions", dynamic.ThenFunc(app.createSession))
	router.Handle(http.MethodGet+" /register", dynamic.ThenFunc(app.showRegisterForm))
	router.Handle(http.MethodPost+" /users", dynamic.ThenFunc(app.createUser))

	protected := dynamic.Append(app.requireAuthentication)
	// Todos
	router.Handle(http.MethodDelete+" /sessions", dynamic.ThenFunc(app.deleteSession))
	router.Handle("GET /", protected.ThenFunc(app.home))
	router.Handle("GET /professional", protected.ThenFunc(app.showProfessionalTodos))
	router.Handle("GET /personal", protected.ThenFunc(app.showPersonalTodos))
	router.Handle("POST /todos", protected.ThenFunc(app.createTodo))
	router.Handle("DELETE /todos/{id}", protected.ThenFunc(app.deleteTodo))
	router.Handle("PATCH /todos/{id}/status", protected.ThenFunc(app.toggleTodoStatus))
	router.Handle("DELETE /todos", protected.ThenFunc(app.deleteCompletedTodos))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(router)
}
