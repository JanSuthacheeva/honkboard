package main

import "net/http"

func (app *application) landingPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, http.StatusOK, "landing.html", "base", data)
}
