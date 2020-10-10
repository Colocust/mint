package route

import (
	"tinyQ/http/server"
)

func Register(router *Router) {
	router.Post("/api/Ticker", server.Runner)
}
