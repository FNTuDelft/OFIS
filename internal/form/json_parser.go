package form

import "io"

// JSONParser implements Parser.
type JSONParser struct{}

func (p *JSONParser) Parse(io.Reader) (*Spec, error) {
	panic("To be implemented")
}

func NewJSONParser() Parser {
	return &JSONParser{}
}
