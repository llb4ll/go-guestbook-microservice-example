package main

import (
	"net/http"
	"net/url"
	"testing"
)

// NopResponseWriter does nothing at all
type NopResponseWriter struct {}

func (w NopResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w NopResponseWriter) Write(d []byte) (int, error) {
	return len(d), nil
}

func (w NopResponseWriter) WriteHeader(int) {
}

// TestGuestBookHandlerRace invokes guestBookHandler from multiple go routines
// in order to detect possible data races.
func TestGuestBookHandlerRace(t *testing.T) {
	f := func() {
		req := http.Request{
			Form: url.Values{},
		}
		req.Form.Set("author", "Heinz")
		req.Form.Set("message", "Hello")
		guestBookHandler(NopResponseWriter{}, &req)
	}

	// Invoke from two different go routines
	go f()
	go f()
}
