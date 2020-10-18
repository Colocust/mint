package route

import (
	"mint/task"
)

func Register(router *Router) {
	router.Post("/task/add", task.Add)
}

func NewRouter() *Router {
	return new(Router)
}
