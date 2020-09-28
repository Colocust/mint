package main

import (
	"net/http"
	"tiny-task/config"
)

func main() {
	config.Load()
}

func run(writer http.ResponseWriter, request *http.Request) {
}
