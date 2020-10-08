package server

import (
	"encoding/json"
	"net/http"
	"tiny-tq/src/api"
	"tiny-tq/src/parameter"
)

func Runner(writer http.ResponseWriter, request *http.Request) {
	var ret parameter.Response
	req := new(parameter.Request)

	body := make([]byte, request.ContentLength)
	_, _ = request.Body.Read(body)
	_ = json.Unmarshal(body, req)

	isValidReq := validateReq(*req)
	if !isValidReq {
		ret = parameter.Response{Code: http.StatusUnsupportedMediaType, Message: "Wrong Args"}
		goto End
	}

	switch request.URL.String() {

	case Ticker:
		go api.Ticker(*req)
		ret = parameter.Response{Code: http.StatusOK, Message: "OK"}

	default:
		ret = parameter.Response{Code: http.StatusInternalServerError, Message: "Wrong Route"}
		break
	}

End:
	result, _ := json.Marshal(ret)
	_, _ = writer.Write(result)
}

func validateReq(request parameter.Request) bool {
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
