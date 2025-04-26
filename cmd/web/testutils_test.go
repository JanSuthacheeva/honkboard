package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jansuthacheeva/honkboard/internal/models/mocks"
)

func NewTestApplication(t *testing.T) *application {
	templateCache, err := newTemplateCache()
	if err != nil {
		t.Fatal(err)
	}

	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true
	sessionManager.Cookie.Name = "session_id"

	return &application{
		// logger:         slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true,})),
		logger:         slog.New(slog.DiscardHandler),
		todos:          &mocks.TodoModel{},
		sessionManager: sessionManager,
		templateCache:  templateCache,
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	ts.Client().Jar = jar

	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string, c *http.Cookie) (int, http.Header, string) {
	client := ts.Client()
	req, err := http.NewRequest(http.MethodGet, ts.URL+urlPath, nil)
	if err != nil {
		t.Fatal(err)
	}
	if c != nil {
		req.AddCookie(c)
	}
	rs, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}

func (ts *testServer) createSessionWithData(t *testing.T, sessionManager *scs.SessionManager, data map[string]any) *http.Cookie {
	// create a handler that accesses the request context and puts the data
	// into the session. Wrap LoadAndSave around it to load and save (555) the
	// session
	handler := sessionManager.LoadAndSave(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Put data into the session
		for key, val := range data {
			sessionManager.Put(ctx, key, val)
		}
	}))

	// create a new request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	// Actually run the request through the middleware + handler
	handler.ServeHTTP(rr, req)

	// Get session cookie written by LoadAndSave
	for _, c := range rr.Result().Cookies() {
		if c.Name == sessionManager.Cookie.Name {
			return c
		}
	}

	t.Fatal("session cookie not found")
	return nil
}
