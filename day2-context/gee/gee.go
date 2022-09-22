package gee

import (
	"log"
	"net/http"
)

type H map[string]any

type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRouter(method, pattern, handlerFunc)
}

func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("GET", pattern, handlerFunc)
}

func (engine *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("POST", pattern, handlerFunc)
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	engine.router.Handler(c)
}
