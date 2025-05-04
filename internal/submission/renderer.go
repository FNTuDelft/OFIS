package submission

import (
	"io"
)

// Renderer the form submission into a specific file type (PDF, JSON, HTML).
type Renderer interface {
	Render(*Submission, io.Writer) error
}
