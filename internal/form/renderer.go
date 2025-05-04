package form

import "io"

// Renderer renders a Spec into a form to be rendered.
type Renderer interface {
	Render(*Spec, io.Writer) error
}
