package gee

import "net/http"

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Method     string
	Path       string
	Params     map[string]string
	StatusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Method: r.Method,
		Path:   r.URL.Path,
	}
}
