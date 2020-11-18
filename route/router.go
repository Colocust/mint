package route

import (
	"mint/api"
)

type (
	HandlerFunc func([]string) *api.Response
	Router      struct {
		route map[string]HandlerFunc
	}
)

func (r *Router) Handle(request []string) *api.Response {
	if len(request) <= 1 {
		res := api.NewResponse(api.StatusArgsError, "Wrong Args")
		return res
	}

	if f, ok := r.route[request[0]]; ok {
		res := f(request[1:])
		return res
	}
	response := api.NewResponse(api.StatusNotFound, "Wrong API")
	return response
}

func (r *Router) Add(path string, handler HandlerFunc) {
	if r.route == nil {
		r.route = make(map[string]HandlerFunc)
	}
	r.route[path] = handler
}

func NewRouter() *Router {
	return new(Router)
}
