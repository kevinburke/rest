// +build go1.7

package rest_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/kevinburke/rest"
)

func ExampleRegisterHandler() {
	rest.RegisterHandler(500, func(w http.ResponseWriter, r *http.Request) {
		err := rest.CtxErr(r)
		fmt.Println("Server error:", err)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(500)
		w.Write([]byte("<html><body>Server Error</body></html>"))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	rest.ServerError(w, req, errors.New("Something bad happened"))
	// Output: Server error: Something bad happened
}

func ExampleClient() {
	client := rest.NewClient("jobs", "secretpassword", "http://ipinfo.io")
	req, _ := client.NewRequest("GET", "/json", nil)
	type resp struct {
		City string `json:"city"`
		Ip   string `json:"ip"`
	}
	var r resp
	client.Do(req, &r)
	fmt.Println(r.Ip)
}
