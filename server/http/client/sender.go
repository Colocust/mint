package client

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

type Sender struct {
	builder *Builder
}

func (s *Sender) Send(method string) (result string, err error) {
	request, err := http.NewRequest(method, s.builder.GetUrl(),
		bytes.NewReader([]byte(s.builder.GetContent())))
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(request)
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
	result = string(bodyContent)

	return
}

func NewSender(b *Builder) *Sender {
	return &Sender{builder: b}
}
