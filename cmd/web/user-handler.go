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

type requestPasswordResetForm struct {
	Email               string `form:"email"`
	validator.Validator `form:"-"`
}

type loginForm struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

type passwordResetForm struct {
	Password            string `form:"password"`
	PasswordConfirm     string `form:"password_confirm"`
	Token               string `form:"token"`
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
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "register.html", "main", data)
		return
	}

	err = app.users.Insert(form.Name, form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.AddFieldError("email", "Email address is already in use")
			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, r, http.StatusUnprocessableEntity, "register.html", "base", data)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	http.Redirect(w, r, "/todos", http.StatusSeeOther)
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
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "login.html", "base", data)
		return
	}

	id, err := app.users.Authenticate(form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.AddNonFieldError("email or password incorrect")
			data := app.newTemplateData(r)
			data.Form = form
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

	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

func (app *application) deleteSession(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Remove(r.Context(), "authenticatedUserID")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) showLoginForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = loginForm{}
	app.render(w, r, http.StatusOK, "login.html", "base", data)
}

func (app *application) showRegisterForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = registerForm{}
	app.render(w, r, http.StatusOK, "register.html", "base", data)
}

func (app *application) showPasswordRequest(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = requestPasswordResetForm{}
	app.render(w, r, http.StatusOK, "request-password-reset.html", "base", data)
}

func (app *application) postPasswordRequest(w http.ResponseWriter, r *http.Request) {
	var form requestPasswordResetForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "request-password-reset.html", "request-password-reset-form", data)
		return
	}

	_, err = app.users.GetByEmail(form.Email)

	if err == nil {
		token, err := app.passwordResetTokens.Insert(form.Email)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		app.background(func() {
			data := app.newTemplateData(r)
			data.PasswordResetToken = token
			err = app.mailer.Send(form.Email, "reset_password.html", data)
			if err != nil {
				app.logger.Error(err.Error())
			}
		})
	}

	data := app.newTemplateData(r)
	app.render(w, r, http.StatusOK, "request-password-reset.html", "password-reset-link-send", data)

}

func (app *application) showPasswordReset(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	queryToken := r.URL.Query().Get("token")
	if queryToken == "" {
		form := requestPasswordResetForm{}
		form.AddFieldError("email", "No token found in request")
		data.Form = form
		app.render(w, r, http.StatusSeeOther, "request-password-reset.html", "base", data)
		return
	}

	token, err := app.passwordResetTokens.Get(queryToken)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			form := requestPasswordResetForm{}
			form.AddFieldError("email", "Invalid or expired token")
			data.Form = form
			app.render(w, r, http.StatusSeeOther, "request-password-reset.html", "base", data)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	data.PasswordResetToken = token
	data.Form = passwordResetForm{}

	app.render(w, r, http.StatusOK, "reset-password.html", "base", data)
}

func (app *application) postPasswordReset(w http.ResponseWriter, r *http.Request) {
	var form passwordResetForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Token), "token", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Token, 64), "token", "This field must be exactly 64 characters long")
	form.CheckField(validator.MaxChars(form.Token, 64), "token", "This field must be exactly 64 characters long")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must have at least 8 chars")
	form.CheckField(validator.ValidPassword(form.Password), "password", "This field must have at least 1 of each: upper case, lower case, number, special char")
	form.CheckField(validator.NotBlank(form.PasswordConfirm), "password_confirm", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.PasswordConfirm, 8), "password_confirm", "This field must have at least 8 chars")
	form.CheckField(validator.ValidPassword(form.PasswordConfirm), "password_confirm", "This field must have at least 1 of each: upper case, lower case, number, special char")
	form.CheckField(validator.Equal(form.Password, form.PasswordConfirm), "password_confirm", "This field must equal the password field")

	data := app.newTemplateData(r)
	if !form.Valid() {
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "reset-password.html", "reset-password-form", data)
		return
	}
	token, err := app.passwordResetTokens.Get(form.Token)
	if err != nil {
		app.serverError(w, r, err)
	}

	err = app.passwordResetTokens.Delete(token.ID)
	if err != nil {
		app.serverError(w, r, err)
	}

	err = app.users.UpdatePassword(token.Email, form.Password)
	if err != nil {
		app.serverError(w, r, err)
	}

	app.render(w, r, http.StatusOK, "reset-password.html", "reset-password-success", data)
}
