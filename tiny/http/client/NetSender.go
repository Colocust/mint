package client

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

type NetSender struct {
	builder NetBuilder
}

func (n *NetSender) Post() (response string, err error) {
	requestPost, err := http.NewRequest(http.MethodPost, n.builder.Url,
		bytes.NewReader([]byte(n.builder.Content)))
	if err != nil {
		return
	}

	requestPost.Header = n.builder.Header

	resp, err := http.DefaultClient.Do(requestPost)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	response = string(bodyContent)

	return
}

