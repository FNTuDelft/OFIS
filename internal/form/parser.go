package form

import "io"

// Parser parses an input file, defining a form in a specific file type, into a generic Spec.
type Parser interface {
	Parse(io.Reader) (*Spec, error)
}
