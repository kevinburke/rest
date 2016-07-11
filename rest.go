package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kevinburke/handlers"
)

// Error implements the HTTP Problem spec laid out here:
// https://tools.ietf.org/html/draft-ietf-appsawg-http-problem-03
type Error struct {
	// The main error message. Should be short enough to fit in a phone's
	// alert box. Do not end this message with a period.
	Title string `json:"title"`

	// Id of this error message ("forbidden", "invalid_parameter", etc)
	ID string `json:"id"`

	// More information about what went wrong.
	Detail string `json:"detail,omitempty"`

	// Path to the object that's in error.
	Instance string `json:"instance,omitempty"`

	// Link to more information (Zendesk, API docs, etc)
	Type       string `json:"type,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
}

func (e *Error) Error() string {
	return e.Title
}

func (e *Error) String() string {
	if e.Detail != "" {
		return fmt.Sprintf("rest: %s. %s", e.Title, e.Detail)
	} else {
		return fmt.Sprintf("rest: %s", e.Title)
	}
}

var serverError = Error{
	StatusCode: http.StatusInternalServerError,
	ID:         "server_error",
	Title:      "Unexpected server error. Please try again",
}

// ServerError logs the error to handlers.Logger, and then responds to the
// request with a generic 500 server error message. ServerError panics if err
// is nil.
func ServerError(w http.ResponseWriter, r *http.Request, err error) error {
	if err == nil {
		panic("rest: no error to log")
	}
	handlers.Logger.Info(fmt.Sprintf("500: %s %s: %s", r.Method, r.URL.Path, err))
	w.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(w).Encode(serverError)
}

func BadRequest(w http.ResponseWriter, r *http.Request, err *Error) error {
	if err == nil {
		panic("rest: no error to write")
	}
	if err.StatusCode == 0 {
		err.StatusCode = http.StatusBadRequest
	}
	log.Printf("400: %s %s: %s", r.Method, r.URL.Path, err.Error())
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(err)
}

var notAllowed = Error{
	Title:      "Method not allowed",
	ID:         "method_not_allowed",
	StatusCode: http.StatusMethodNotAllowed,
}

// NotAllowed returns a generic HTTP 405 Not Allowed status and response body
// to the client.
func NotAllowed(w http.ResponseWriter, r *http.Request) error {
	e := notAllowed
	e.Instance = r.URL.Path
	w.WriteHeader(http.StatusMethodNotAllowed)
	return json.NewEncoder(w).Encode(e)
}

func Forbidden(w http.ResponseWriter, r *http.Request, err *Error) error {
	w.WriteHeader(http.StatusForbidden)
	return json.NewEncoder(w).Encode(err)
}
