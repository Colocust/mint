package main

import (
	"net/http"
	"tiny-tq/config"
	"tiny-tq/server"
)

func init() {
	//加载配置文件
	config.Load()
}

func main() {
	http.HandleFunc("/", server.Runner)

	addr := config.Read("ip").(string) + ":" + config.Read("port").(string)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
