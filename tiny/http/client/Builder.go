package client

type Builder interface {
	SetUrl(url string) Builder
	SetHeader(key string, value string) Builder
	SetContent(content string) Builder
	DelHeader(key string) Builder
	NewNetSender() *NetSender
}
