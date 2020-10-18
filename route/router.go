package route

import (
	"encoding/json"
	"mint/task"
	"net/http"
	"strings"
)

type (
	HandlerFunc func(*task.AddTaskRequest) *task.AddTaskResponse

	Router struct {
		route map[string]map[string]HandlerFunc
	}
)

func (r *Router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if f, ok := r.route[req.Method][req.URL.String()]; ok {
		decoder := json.NewDecoder(req.Body)
		params := new(task.AddTaskRequest)
		decoder.Decode(params)

		resp := f(params)

		data, _ := json.Marshal(resp)
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
