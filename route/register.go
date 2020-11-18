package route

import "mint/api"

func Register(r *Router) {
	r.Add("add", api.Add)
}
