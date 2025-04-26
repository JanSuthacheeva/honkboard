package main

import (
	"net/http"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {

	// validate input

	// create user

	// set session

	// redirect

}

func (app *application) showLoginForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "login.html", "base", templateData{})
}

func (app *application) showRegisterForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "register.html", "base", templateData{})
}
