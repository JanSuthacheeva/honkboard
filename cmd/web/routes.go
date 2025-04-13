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

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(router)
}
