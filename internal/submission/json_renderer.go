package submission

import "io"

// JSONRenderer implements Renderer.
type JSONRenderer struct{}

func (r *JSONRenderer) Render(*Submission, io.Writer) error {
	panic("To be implemented")
}

func NewJSONRenderer() Renderer {
	return &JSONRenderer{}
}
