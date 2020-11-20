package http

type Builder struct {
	Url     string
	Content string
}

func (b *Builder) setUrl(url string) *Builder {
	b.Url = url
	return b
}

func (b *Builder) setContent(content string) *Builder {
	b.Content = content
	return b
}

func NewBuilder() *Builder {
	return new(Builder)
}
