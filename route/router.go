package route

import (
	"mint/api"
	"sync"
)

var (
	instance *Router
	once     sync.Once
)

type (
	HandlerFunc func([]string) *api.Response
	Router      struct {
		route map[string]HandlerFunc
	}
)

func (r *Router) Handle(request []string) *api.Response {
	if len(request) <= 1 {
		resp := api.NewResponse(api.StatusArgsError, "Wrong Args")
		return resp
	}

	if f, ok := r.route[request[0]]; ok {
		resp := f(request[1:])
		return resp
	}
	resp := api.NewResponse(api.StatusNotFound, "Wrong API")
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
		Register(instance)
	})
	return instance
}
