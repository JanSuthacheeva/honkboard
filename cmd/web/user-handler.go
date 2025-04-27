package main

import (
	"errors"
	"net/http"

	"github.com/jansuthacheeva/honkboard/internal/models"
	"github.com/jansuthacheeva/honkboard/internal/validator"
)

type registerForm struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	PasswordConfirm     string `form:"password_confirm"`
	validator.Validator `form:"-"`
}

type loginForm struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {

	var form registerForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Name, 2), "name", "This field must have at least 2 chars")
	form.CheckField(validator.MaxChars(form.Name, 90), "name", "This field must not have more than 90 chars")
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must have at least 8 chars")
	form.CheckField(validator.ValidPassword(form.Password), "password", "This field must have at least 1 of each: upper case, lower case, number, special char")
	form.CheckField(validator.NotBlank(form.PasswordConfirm), "password_confirm", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.PasswordConfirm, 8), "password_confirm", "This field must have at least 8 chars")
	form.CheckField(validator.ValidPassword(form.PasswordConfirm), "password_confirm", "This field must have at least 1 of each: upper case, lower case, number, special char")
	form.CheckField(validator.Equal(form.Password, form.PasswordConfirm), "password_confirm", "This field must equal the password field")

	if !form.Valid() {
		data := templateData{
			Form: form,
		}
		app.render(w, r, http.StatusUnprocessableEntity, "register.html", "main", data)
		return
	}

	err = app.users.Insert(form.Name, form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.AddFieldError("email", "Email address is already in use")
			data := templateData{
				Form: form,
			}
			app.render(w, r, http.StatusUnprocessableEntity, "register.html", "base", data)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) createSession(w http.ResponseWriter, r *http.Request) {
	var form loginForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must have at least 8 chars")
	form.CheckField(validator.ValidPassword(form.Password), "password", "This field must have at least 1 of each: upper case, lower case, number, special char")

	if !form.Valid() {
		data := templateData{
			Form: form,
		}
		app.render(w, r, http.StatusUnprocessableEntity, "login.html", "base", data)
		return
	}

	id, err := app.users.Authenticate(form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.AddNonFieldError("email or password incorrect")
			data := templateData{
				Form: form,
			}

			app.render(w, r, http.StatusUnprocessableEntity, "login.html", "base", data)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (app *application) showLoginForm(w http.ResponseWriter, r *http.Request) {
	data := templateData{
		Form: loginForm{},
	}
	app.render(w, r, http.StatusOK, "login.html", "base", data)
}

func (app *application) showRegisterForm(w http.ResponseWriter, r *http.Request) {
	data := templateData{
		Form: registerForm{},
	}
	app.render(w, r, http.StatusOK, "register.html", "base", data)
}
