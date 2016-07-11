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
