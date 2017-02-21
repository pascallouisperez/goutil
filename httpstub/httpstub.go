package httpstub

import (
	"gopkg.in/check.v1"
	"net/http"
	"net/url"
)

// MustParseURL parses a URL by delegating to net/url#Parse and panics in the
// case of an error.
func MustParseURL(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}
	return u
}

type StubResponseWriter struct {
	c              *check.C
	RecordedBody   string
	RecordedHeader map[string][]string
	RecordedCode   int

	headerCodeRecorded bool
}

// Assert that StubResponseWriter implements the http.ResponseWriter interface.
var _ http.ResponseWriter = &StubResponseWriter{}

// New creates a stub response writer, and a request to help test HTTP handlers.
// Intended us is as follows:
//
//     w, r := httpstub.New()
//     ... configure w and r
//     myHandler(w, r)
//     ... assert properties on w and r
func New(c *check.C) (*StubResponseWriter, *http.Request) {
	w := &StubResponseWriter{
		c:              c,
		RecordedHeader: make(map[string][]string),
		RecordedCode:   http.StatusOK,
	}
	r := &http.Request{}
	return w, r
}

func (w *StubResponseWriter) Header() http.Header {
	return w.RecordedHeader
}

func (w *StubResponseWriter) Write(data []byte) (int, error) {
	if !w.headerCodeRecorded {
		w.WriteHeader(http.StatusOK)
	}
	w.RecordedBody = string(data)
	return len(data), nil
}

func (w *StubResponseWriter) WriteHeader(code int) {
	if !w.headerCodeRecorded {
		w.RecordedCode = code
		w.headerCodeRecorded = true
	} else {
		w.c.Errorf("header already set to %d", w.RecordedCode)
	}
}
