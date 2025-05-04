package submission

import (
	"fmt"
	"io"
	"net/http"
)

type Service struct {
	parser    *Parser
	validator Validator
	renderer  Renderer
}

// HandleFormSubmission handles the form submition.
// It validates the submission, and if valid, renders the results.
func (s *Service) HandleFormSubmission(r *http.Request, w io.Writer) error {
	submission, err := s.parser.Parse(r)
	if err != nil {
		return fmt.Errorf("could not parse form submission: %w", err)
	}

	err = s.validator.Validate(submission)
	if err != nil {
		return fmt.Errorf("invalid form submission: %w", err)
	}

	err = s.renderer.Render(submission, w)
	if err != nil {
		return fmt.Errorf("could not render form submission: %w", err)
	}

	return nil
}

func NewService(parser *Parser, validator Validator, renderer Renderer) *Service {
	return &Service{
		parser:    parser,
		validator: validator,
		renderer:  renderer,
	}
}
