package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	fileServer := http.FileServer(http.Dir("./static/"))

	router := http.NewServeMux()

	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("/", app.home)
	router.HandleFunc("/professional", app.showProfessionalTodos)
	router.HandleFunc("/personal", app.showPersonalTodos)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders, app.sessionManager.LoadAndSave)

	return standard.Then(router)
}
