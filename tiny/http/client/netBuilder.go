package client

import (
	"net/http"
)

type NetBuilder struct {
	Url     string
	Content string
	Header  http.Header
}

func (n *NetBuilder) SetHeader(key string, value string) Builder {
	n.Header.Add(key, value)
	return n
}

func (n *NetBuilder) SetContent(content string) Builder {
	n.Content = content
	return n
}

func (n *NetBuilder) DelHeader(key string) Builder {
	n.Header.Del(key)
	return n
}

func (n *NetBuilder) SetUrl(url string) Builder {
	n.Url = url
	return n
}

func (n *NetBuilder) NewNetSender() *NetSender {
	netSender := new(NetSender)
	netSender.builder = *n
	return netSender
}
