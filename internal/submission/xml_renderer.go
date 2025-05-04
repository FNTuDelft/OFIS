package submission

import "io"

// XMLRenderer implements Renderer.
type XMLRenderer struct{}

func (r *XMLRenderer) Render(*Submission, io.Writer) error {
	panic("To be implemented")
}

func NewXMLRenderer() Renderer {
	return &XMLRenderer{}
}
