package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestPost(t *testing.T) {
	t.Parallel()
	var user, pass string
	var ok bool
	var requestUrl *url.URL
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok = r.BasicAuth()
		requestUrl = r.URL
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("{}"))
	}))
	defer s.Close()
	client := NewClient("foo", "bar", s.URL)
	req, err := client.NewRequest("POST", "/", nil)
	assertNotError(t, err, "")
	err = client.Do(req, &struct{}{})
	assertNotError(t, err, "")
	assertEquals(t, user, "foo")
	assertEquals(t, pass, "bar")
	assertEquals(t, requestUrl.Path, "/")
}

func TestPostError(t *testing.T) {
	t.Parallel()
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Error{
			Title: "bad request",
			ID:    "something_bad",
		})
	}))
	defer s.Close()
	client := NewClient("foo", "bar", s.URL)
	req, err := client.NewRequest("POST", "/", nil)
	assertNotError(t, err, "")
	err = client.Do(req, &struct{}{})
	assertError(t, err, "")
	assertEquals(t, err.Error(), "bad request")
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randFilename() string {
	b := make([]byte, 12)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return filepath.Join(os.TempDir(), string(b))
}

func TestSocket(t *testing.T) {
	fname := randFilename()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte("{\"status\": 200, \"message\": \"hello world\"}"))
		assertNotError(t, err, "socket write")
	})
	server := http.Server{
		Handler: mux,
	}
	listener, err := net.Listen("unix", fname)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan bool, 1)
	go func() {
		<-ch
		fmt.Println("closing listener")
		listener.Close()
	}()
	go server.Serve(listener)
	c := &Client{Base: "http://localhost"}
	c.DialSocket(fname)
	req, err := c.NewRequest("GET", "/", nil)
	assertNotError(t, err, "creating http request")
	var b struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	err = c.Do(req, &b)
	assertNotError(t, err, "c.Do")
	assertEquals(t, b.Status, 200)
	assertEquals(t, b.Message, "hello world")
	ch <- true
}

func ExampleClient() {
	client := NewClient("jobs", "secretpassword", "http://ipinfo.io")
	req, _ := client.NewRequest("GET", "/json", nil)
	type resp struct {
		City string `json:"city"`
		Ip   string `json:"ip"`
	}
	var r resp
	client.Do(req, &r)
	fmt.Println(r.Ip)
}