package tinyQ

import (
	"net/http"
	"tinyQ/config"
	"tinyQ/route"
)

func Boot() {
	router := route.NewRouter()
	route.Register(router)

	addr := config.Read("ip").(string) + ":" + config.Read("port").(string)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
