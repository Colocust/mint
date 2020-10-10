package route

import (
	"tinyQ/api"
)

func Register(router *Router) {
	router.Post("/api/Ticker", api.Ticker)
}
