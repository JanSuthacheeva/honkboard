package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	fileServer := http.FileServer(http.Dir("./static/"))

	router := http.NewServeMux()

	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("GET /", app.home)
	router.HandleFunc("GET /professional", app.showProfessionalTodos)
	router.HandleFunc("GET /personal", app.showPersonalTodos)

	router.HandleFunc("POST /todos", app.createTodo)
	router.HandleFunc("DELETE /todos/{id}", app.deleteTodo)
	router.HandleFunc("DELETE /todos", app.deleteCompletedTodos)
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders, app.sessionManager.LoadAndSave)

	return standard.Then(router)
}
