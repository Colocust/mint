package route

import (
	"mint/http/server"
	"sync"
)

var (
	instance *Router
	once     sync.Once
)

type (
	HandlerFunc func([]string, *server.Response)
	Router      struct {
		route map[string]HandlerFunc
	}
)

func (r *Router) Handle(request []string) *server.Response {
	if len(request) <= 1 {
		resp := server.NewResponse(server.StatusArgsError, "Wrong Args")
		return resp
	}

	if f, ok := r.route[request[0]]; ok {
		resp := &server.Response{}
		f(request[1:], resp)
		return resp
	}
	resp := server.NewResponse(server.StatusNotFound, "Wrong API")
	return resp
}

func (r *Router) Add(path string, handler HandlerFunc) {
	if r.route == nil {
		r.route = make(map[string]HandlerFunc)
	}
	r.route[path] = handler
}

func GetInstance() *Router {
	once.Do(func() {
		instance = new(Router)
	})
	return instance
}
