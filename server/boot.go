package server

import (
	"encoding/json"
	"net/http"
	"tiny-tq/api"
	http_ "tiny-tq/http"
	"tiny-tq/route"
)

func Runner(writer http.ResponseWriter, request *http.Request) {
	var ret http_.Response
	req := new(api.Request)

	body := make([]byte, request.ContentLength)
	_, _ = request.Body.Read(body)
	_ = json.Unmarshal(body, req)

	isValidReq := validateReq(*req)
	if !isValidReq {
		ret = http_.Response{Code: http_.ARGS, Message: "Wrong Args"}
		goto End
	}

	switch request.URL.String() {

	case route.Ticker:
		go api.Ticker(*req)
		ret = http_.Response{Code: http_.SUCCESS, Message: "OK"}

	default:
		ret = http_.Response{Code: http_.ERROR, Message: "Wrong Route"}
		break
	}

End:
	result, _ := json.Marshal(ret)
	_, _ = writer.Write(result)
}

func validateReq(request api.Request) bool {
	if request.Url == "" {
		return false
	}
	if request.Method == "" {
		return false
	}
	if request.Duration == 0 {
		return false
	}
	return true
}
