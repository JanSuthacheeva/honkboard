package main

import (
	"errors"
	"net/http"
	"strconv"

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

type resetPasswordCodeForm struct {
	Code                string `form:"code"`
	validator.Validator `form:"-"`
}

type newPasswordForm struct {
	Password        string `form:"password"`
	PasswordConfirm string `form:"password_confirm"`
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
		app.render(w, r, http.StatusUnprocessableEntity, "request-password-reset.html", "base", data)
		return
	}

	userId, err := app.users.GetByEmail(form.Email)
	if err != nil {
		data := app.newTemplateData(r)
		form.AddFieldError("email", "This email is not known in our system")
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "request-password-reset.html", "base", data)
		return
	}

	code, err := app.validationCodes.Insert(userId, "reset-password")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.background(func() {
		err = app.mailer.Send(form.Email, "reset_password.html", code)
		if err != nil {
			app.logger.Error(err.Error())
		}
	})

	http.Redirect(w, r, "/reset-password-code", http.StatusSeeOther)
}

func (app *application) showResetPasswordCode(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = resetPasswordCodeForm{}
	app.render(w, r, http.StatusOK, "request-password-validation.html", "base", data)
}

func (app *application) postResetPasswordCode(w http.ResponseWriter, r *http.Request) {
	var form resetPasswordCodeForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Code), "code", "This field must be exactly six digits long")
	form.CheckField(validator.MinChars(form.Code, 6), "code", "This field must be exactly six digits long")
	form.CheckField(validator.MaxChars(form.Code, 6), "code", "This field must be exactly six digits long")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "request-password-validation.html", "base", data)
		return
	}

	codeAsInt, err := strconv.Atoi(form.Code)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	_, err = app.validationCodes.GetByCode(codeAsInt)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, r, http.StatusUnprocessableEntity, "request-password-validation.html", "base", data)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	// authenticate user
	// redirect to reset-password
}

func (app *application) showNewPassword(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = newPasswordForm{}

	app.render(w, r, http.StatusOK, "reset-password.html", "base", data)
}
