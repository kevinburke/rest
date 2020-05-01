package rest

import "testing"

func TestCompat(t *testing.T) {
	_ = &Error{Title: "foo"}
	_ = &Client{}
	d := NewClient("http://base", "user", "pass")
	d.Client.Transport = DefaultTransport
}
