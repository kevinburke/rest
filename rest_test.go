package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerError(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	ServerError(w, req, errors.New("foo bar"))
	if w.Code != 500 {
		t.Errorf("expected code to be 500, got %d", w.Code)
	}
	var e Error
	err := json.NewDecoder(w.Body).Decode(&e)
	if err != nil {
		t.Fatal(err)
	}
	if e.Title == "" {
		t.Errorf("expected Title to be %s, got the empty string", serverError.Title)
	}
	if e.Title != serverError.Title {
		t.Errorf("expected Title to be %s, got %s", serverError.Title, e.Title)
	}
	if e.ID != serverError.ID {
		t.Errorf("expected ID to be %s, got %s", serverError.ID, e.ID)
	}
	if e.StatusCode != 500 {
		t.Errorf("expected code to be 500, got %d", e.StatusCode)
	}
}

func TestBadRequest(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	BadRequest(w, req, &Error{
		Title: "Please provide a widget",
		ID:    "missing_widget",
	})
	if w.Code != 400 {
		t.Errorf("expected code to be 400, got %d", w.Code)
	}
	var e Error
	err := json.NewDecoder(w.Body).Decode(&e)
	if err != nil {
		t.Fatal(err)
	}
	if e.Title == "" {
		t.Errorf("expected Title to be %s, got the empty string", "Please provide a widget")
	}
	if e.Title != "Please provide a widget" {
		t.Errorf("expected Title to be %s, got %s", "Please provide a widget", e.Title)
	}
	if e.StatusCode != 400 {
		t.Errorf("expected code to be 400, got %d", e.StatusCode)
	}
	if e.ID != "missing_widget" {
		t.Errorf("expected ID to be %s, got %s", "missing_widget", e.ID)
	}
}

func TestCustomBadRequest(t *testing.T) {
	err := &Error{Title: "bad"}
	RegisterHandler(400, func(w http.ResponseWriter, r *http.Request) {
		err := CtxErr(r)
		if err == nil {
			t.Fatal("expected non-nil error, got nil")
		}
		if rerr, ok := err.(*Error); ok {
			if rerr.Title != "bad" {
				t.Errorf("expected err.Title to be 'bad', got %s", rerr.Title)
			}
		} else {
			t.Fatalf("expected err cast to Error, couldn't, err is %v", err)
		}
		w.Header().Set("Custom-Handler", "true")
		w.WriteHeader(400)
		w.Write([]byte("hello world"))
	})
	defer func() {
		handlerMu.Lock()
		defer handlerMu.Unlock()
		delete(handlerMap, 400)
	}()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	BadRequest(w, req, err)
	if hdr := w.Header().Get("Custom-Handler"); hdr != "true" {
		t.Errorf("expected to get Custom-Handler: true header, got %s", hdr)
	}
	if body := w.Body.String(); body != "hello world" {
		t.Errorf("expected Body to be hello world, got %s", body)
	}
}

func TestNoContent(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	NoContent(w)
	if w.Code != 204 {
		t.Errorf("expected Code to be 204, got %d", w.Code)
	}
	if hdr := w.Header().Get("Content-Type"); hdr != "" {
		t.Errorf("expected Content-Type to be empty, got %s", hdr)
	}
}

func TestUnauthorized(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	req, _ := http.NewRequest("GET", "/", nil)
	Unauthorized(w, req, "foo")
	if w.Code != 401 {
		t.Errorf("expected Code to be 401, got %d", w.Code)
	}
	expected := `Basic realm="foo"`
	if hdr := w.Header().Get("WWW-Authenticate"); hdr != expected {
		t.Errorf("expected WWW-Authenticate header to be %s, got %s", expected, hdr)
	}
}
