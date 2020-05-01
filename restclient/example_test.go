package restclient_test

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/kevinburke/rest/restclient"
)

func ExampleClient() {
	client := restclient.New("jobs", "secretpassword", "http://ipinfo.io")
	req, _ := client.NewRequest("GET", "/json", nil)
	type resp struct {
		City string `json:"city"`
		Ip   string `json:"ip"`
	}
	var r resp
	client.Do(req, &r)
	fmt.Println(r.Ip)
}

func ExampleNew() {
	client := restclient.New("jobs", "secretpassword", "http://ipinfo.io")
	req, _ := client.NewRequest("GET", "/json", nil)
	type resp struct {
		City string `json:"city"`
		Ip   string `json:"ip"`
	}
	var r resp
	client.Do(req, &r)
	fmt.Println(r.Ip)
}

func ExampleClient_NewRequest() {
	client := restclient.New("jobs", "secretpassword", "http://ipinfo.io")
	req, _ := client.NewRequest("GET", "/json", nil)
	type resp struct {
		City string `json:"city"`
		Ip   string `json:"ip"`
	}
	var r resp
	client.Do(req, &r)
	fmt.Println(r.Ip)
}

func ExampleClient_Do() {
	client := restclient.New("jobs", "secretpassword", "http://ipinfo.io")
	req, _ := client.NewRequest("GET", "/json", nil)
	type resp struct {
		City string `json:"city"`
		Ip   string `json:"ip"`
	}
	var r resp
	client.Do(req, &r)
	fmt.Println(r.Ip)
}

func ExampleTransport() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}))
	defer server.Close()
	b := new(bytes.Buffer)
	client := http.Client{
		Transport: &restclient.Transport{Debug: true, Output: b},
	}
	req, _ := http.NewRequest("GET", server.URL+"/bar", nil)
	client.Do(req)

	// Dump the HTTP request from the buffer, but skip the lines that change.
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "Host:") || strings.HasPrefix(text, "Date:") {
			continue
		}
		fmt.Println(text)
	}
	// Output:
	// GET /bar HTTP/1.1
	// User-Agent: Go-http-client/1.1
	// Accept-Encoding: gzip
	//
	// HTTP/1.1 200 OK
	// Content-Length: 11
	// Content-Type: text/plain; charset=utf-8
	//
	// Hello World
}
