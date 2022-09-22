package gee

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handlerFunc
}

func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	engine.addRoute("GET", pattern, handlerFunc)
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := request.Method + "-" + request.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(writer, request)
	} else {
		fmt.Fprintf(writer, "404 NOT FOUND : %s\n", request.URL.Path)
	}
}
