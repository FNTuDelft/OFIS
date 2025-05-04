package submission

import "io"

// HTMLRenderer implements Renderer.
type HTMLRenderer struct{}

func (r *HTMLRenderer) Render(*Submission, io.Writer) error {
	panic("To be implemented")
}

func NewHTMLRenderer() Renderer {
	return &HTMLRenderer{}
}
