package submission

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"ofis/internal/errors"
)

type Parser struct{}

// Parse parses the input from a POST request into a Submission.
func (p *Parser) Parse(r *http.Request) (*Submission, error) {
	// Parse POST results into r.Form field
	err := r.ParseMultipartForm(1 << 20) // 1 MB
	if err != nil {
		return nil, &errors.ParseError{
			Message: "form submission could not be parsed",
			Err:     err,
		}
	}

	// Parse POSTed files into FileSubmissions
	files := make(FileSubmissions, len(r.MultipartForm.File))
	for inputName, fileHeaders := range r.MultipartForm.File {
		filesBytes, err := parseFiles(fileHeaders)
		if err != nil {
			return nil, &errors.ParseError{
				Message: "could not parse file(s)",
				Err:     err,
			}
		}

		files[inputName] = filesBytes
	}

	return NewSubmission(r.Form, files), nil
}

// parseFiles takes file headers and reads them into bytes.
// It returns a map where each key is the file name and the value is the file byte slice.
func parseFiles(fileHeaders []*multipart.FileHeader) (map[string][]byte, error) {
	files := make(map[string][]byte, len(fileHeaders))
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("could not open file: %w", err)
		}
		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("could not read file %s: %w", fileHeader.Filename, err)
		}

		files[fileHeader.Filename] = fileBytes
	}

	return files, nil
}

func NewParser() *Parser {
	return &Parser{}
}
