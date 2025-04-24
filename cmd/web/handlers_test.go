package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/jansuthacheeva/honkboard/internal/assert"
	"github.com/jansuthacheeva/honkboard/internal/enums"
)

func TestHome(t *testing.T) {
	os.Chdir("../../")
	app := NewTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		listType string
		wantCode int
		wantBody string
	}{
		{
			name:     "personal todos",
			listType: enums.TodoTypePersonal.String(),
			wantCode: http.StatusOK,
			wantBody: "Mock Todo",
		},
		{
			name:     "professional todos",
			listType: enums.TodoTypeProfessional.String(),
			wantCode: http.StatusOK,
			wantBody: "Quiet day today... Let's add some tasks!",
		},
		{
			name:     "invalid list type",
			listType: "invalid",
			wantCode: http.StatusInternalServerError,
		},
		{
			name:     "no list type",
			wantCode: http.StatusOK,
			wantBody: "Mock Todo",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			cookie := ts.createSessionWithData(t, app.sessionManager, map[string]any{
				"list-type": tt.listType,
			})

			code, _, body := ts.get(t, "/", cookie)

			assert.Equal(t, code, tt.wantCode)
			if tt.wantBody != "" {
				assert.StringContains(t, body, tt.wantBody)
			}
		})
	}
}

func TestShowPersonalTodos(t *testing.T) {
	app := NewTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name         string
		wantCode     int
		wantBody     string
		wantListType string
	}{
		{
			name:         "shows todo",
			wantCode:     http.StatusOK,
			wantBody:     "Mock Todo",
			wantListType: "Personal",
		},
	}

	for _, tt := range tests {
		code, _, body := ts.get(t, "/personal", nil)

		assert.Equal(t, code, tt.wantCode)
		if tt.wantBody != "" {
			assert.StringContains(t, body, tt.wantBody)
		}
	}
}

func TestShowProfessionalTodos(t *testing.T) {
	app := NewTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name         string
		wantCode     int
		wantBody     string
		wantListType string
	}{
		{
			name:         "shows nothing",
			wantCode:     http.StatusOK,
			wantBody:     "Quiet day today",
			wantListType: "Professional",
		},
	}

	for _, tt := range tests {
		code, _, body := ts.get(t, "/professional", nil)

		assert.Equal(t, code, tt.wantCode)
		if tt.wantBody != "" {
			assert.StringContains(t, body, tt.wantBody)
		}
	}
}
