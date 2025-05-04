package form

import (
	"log"
	"path/filepath"
)

type ParserBuilder struct{}

// Build builds the appropriate parser based on the file type of the file to be parsed.
func (b *ParserBuilder) Build(fileToBeParsed string) Parser {
	fileExtension := filepath.Ext(fileToBeParsed)

	switch fileExtension {
	case ".xml":
		return NewXMLParser()
	case ".json":
		return NewJSONParser()
	default:
		log.Panicf("file extension '%s' does not have matching parser", fileExtension)

		return nil
	}
}

func NewParserBuilder() *ParserBuilder {
	return &ParserBuilder{}
}
