package http

type Builder struct {
	url     string
	content string
}

func (b *Builder) SetUrl(url string) *Builder {
	b.url = url
	return b
}

func (b *Builder) SetContent(content string) *Builder {
	b.content = content
	return b
}

func (b *Builder) GetContent() string {
	return b.content
}

func (b *Builder) GetUrl() string {
	return b.url
}

func NewBuilder() *Builder {
	return new(Builder)
}
