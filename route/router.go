package route

import (
	"encoding/json"
	"net/http"
	"strings"
	"tinyQ/http/server"
)

type (
	HandlerFunc func(req *http.Request) server.Response

	Router struct {
		route map[string]map[string]HandlerFunc
	}
)

func NewRouter() *Router {
	return new(Router)
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	f, ok := r.route[req.Method][req.URL.String()]
	if ok {
		response := f(req)

		data, _ := json.Marshal(response)
		writer.Write(data)
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func (r *Router) handleFunc(method string, path string, handlerFunc HandlerFunc) {
	method = strings.ToUpper(method)
	if r.route == nil {
		r.route = make(map[string]map[string]HandlerFunc)
	}
	if r.route[method] == nil {
		r.route[method] = make(map[string]HandlerFunc)
	}
	r.route[method][path] = handlerFunc
}

func (r *Router) Get(path string, handlerFunc HandlerFunc) {
	r.handleFunc(http.MethodGet, path, handlerFunc)
}

func (r *Router) Post(path string, handlerFunc HandlerFunc) {
	r.handleFunc(http.MethodPost, path, handlerFunc)
}
