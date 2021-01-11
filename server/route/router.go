package route

import (
	"mint/server/http/server"
	"sync"
)

var (
	instance *Router
	once     sync.Once
)

type (
	HandlerFunc func(string, *server.Response)
	Router      struct {
		route map[string]HandlerFunc
	}
)

func (r *Router) Handle(request []string, resp *server.Response) {
	if len(request) <= 1 {
		resp.Code = server.StatusArgsError
		return
	}
	if f, ok := r.route[request[0]]; ok {
		f(request[1], resp)
		resp.Code = server.StatusSuccess
		return
	}
	resp.Code = server.StatusNotFound
	return
}

func (r *Router) Add(uri string, handler HandlerFunc) {
	if r.route == nil {
		r.route = make(map[string]HandlerFunc)
	}
	r.route[uri] = handler
}

func GetInstance() *Router {
	once.Do(func() {
		instance = new(Router)
	})
	return instance
}

func Register(router *Router) {
	router.Add("delay", server.Delay)
}
