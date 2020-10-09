package router

import (
	"net/http"
	"strings"
)

type Router struct {
	route map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return new(Router)
}

func (r *Router) handleFunc(method string, path string, handlerFunc http.HandlerFunc) {
	method = strings.ToUpper(method)
	if r.route == nil {
		r.route = make(map[string]map[string]http.HandlerFunc)
	}
	if r.route[method] == nil {
		r.route[method] = make(map[string]http.HandlerFunc)
	}
	r.route[method][path] = handlerFunc
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handlerFunc, ok := r.route[request.Method][request.URL.String()]
	if ok {
		handlerFunc(writer, request)
	} else {
		writer.WriteHeader(404)
	}
}

func (r *Router) Get(path string, handlerFunc http.HandlerFunc) {
	r.handleFunc(http.MethodGet, path, handlerFunc)
}

func (r *Router) Post(path string, handlerFunc http.HandlerFunc) {
	r.handleFunc(http.MethodPost, path, handlerFunc)
}
