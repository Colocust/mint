package route

import "mint/api"

func Register(router *Router) {
	router.Add("add", api.Add)
}
