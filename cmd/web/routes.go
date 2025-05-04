package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	fileServer := http.FileServer(http.Dir("./static/"))

	router := http.NewServeMux()

	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)
	router.Handle(http.MethodGet+" /", dynamic.ThenFunc(app.landingPage))

	public := dynamic.Append(app.noAuth)
	// Users
	router.Handle(http.MethodGet+" /login", public.ThenFunc(app.showLoginForm))
	router.Handle(http.MethodPost+" /sessions", public.ThenFunc(app.createSession))
	router.Handle(http.MethodGet+" /register", public.ThenFunc(app.showRegisterForm))
	router.Handle(http.MethodPost+" /users", public.ThenFunc(app.createUser))
	router.Handle(http.MethodGet+" /request-password-link", public.ThenFunc(app.showPasswordRequest))
	router.Handle(http.MethodPost+" /request-password-link", public.ThenFunc(app.postPasswordRequest))
	router.Handle(http.MethodGet+" /reset-password-code", public.ThenFunc(app.showResetPasswordCode))
	router.Handle(http.MethodPost+" /reset-password-code", public.ThenFunc(app.postResetPasswordCode))

	protected := dynamic.Append(app.requireAuthentication)
	// Todos
	router.Handle(http.MethodDelete+" /sessions", dynamic.ThenFunc(app.deleteSession))
	router.Handle("GET /todos", protected.ThenFunc(app.home))
	router.Handle("GET /professional", protected.ThenFunc(app.showProfessionalTodos))
	router.Handle("GET /personal", protected.ThenFunc(app.showPersonalTodos))
	router.Handle("POST /todos", protected.ThenFunc(app.createTodo))
	router.Handle("DELETE /todos/{id}", protected.ThenFunc(app.deleteTodo))
	router.Handle("PATCH /todos/{id}/status", protected.ThenFunc(app.toggleTodoStatus))
	router.Handle("DELETE /todos", protected.ThenFunc(app.deleteCompletedTodos))

	router.Handle(http.MethodGet+" /reset-password", protected.ThenFunc(app.showResetPassword))
	router.Handle(http.MethodPost+" /reset-password", protected.ThenFunc(app.postResetPassword))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(router)
}
