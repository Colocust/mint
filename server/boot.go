package server

import (
	"fmt"
	"net/http"
)

func Runner(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.String() {
	case "/send":
		fmt.Println("s")
		break

	default:
		fmt.Println("err")
		break
	}
}
