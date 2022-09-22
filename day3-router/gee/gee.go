package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

func (e *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	e.addRouter("GET", pattern, handlerFunc)
}

func (e *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	e.addRouter("POST", pattern, handlerFunc)
}

func (e *Engine) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	e.router.addRouter(method, pattern, handlerFunc)
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := NewContext(writer, request)
	e.router.hanler(c)
}
