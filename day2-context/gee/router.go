package gee

import "fmt"

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handlerFunc
}

func (r *router) Handler(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.Writer, "404 NOT FOUND : %s\n", c.Path)
	}
}
