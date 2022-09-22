package gee

import "log"

type HandlerFunc func(*Context)

type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}

	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

func New() *Engine {
	engine := &Engine{
		router: NewRouter(),
	}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) addRouter(method string, comp string, handlerFunc HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRouter(method, pattern, handlerFunc)
}

func (group *RouterGroup) GET(pattern string, handlerFunc HandlerFunc) {
	group.addRouter("GET", pattern, handlerFunc)
}
