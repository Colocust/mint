package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, content map[string]interface{}, contentType string) {
	body, _ := json.Marshal(content)
	resp, err := http.Post(url, contentType, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func Get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
