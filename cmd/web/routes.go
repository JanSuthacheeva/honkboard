package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {

	fileServer := http.FileServer(http.Dir("./static/"))

	router := http.NewServeMux()

	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("/", app.home)

	return router
}
