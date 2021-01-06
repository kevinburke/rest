package rest

import "testing"

func TestCompat(t *testing.T) {
	_ = &Error{Title: "foo"}
	_ = &Client{}
	_ = NewBearerClient("token", "http://base")
	d := NewClient("http://base", "user", "pass")
	d.Client.Transport = DefaultTransport
}
