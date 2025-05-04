package submission

import (
	"fmt"
	"ofis/internal/errors"
	"ofis/internal/form"
)

// Validator validates the form submission.
type Validator interface {
	Validate(*Submission) error
}

// DefaultValidator implements Validator.
type DefaultValidator struct {
	spec *form.Spec
}

// Validate validates if all the required elements are in the submission.
func (v *DefaultValidator) Validate(submission *Submission) error {
	for _, elem := range v.spec.Elements {
		switch e := elem.(type) {
		case form.FieldSpec:
			err := validateField(submission, e)
			if err != nil {
				return fmt.Errorf("invalid field in submission: %w", err)
			}
		case form.SectionSpec:
			err := validateSection(submission, e)
			if err != nil {
				return fmt.Errorf("invalid section in submission: %w", err)
			}
		default:
		}
	}

	return nil
}

func validateSection(submission *Submission, section form.SectionSpec) error {
	for _, field := range section.Fields {
		err := validateField(submission, field)
		if err != nil {
			return fmt.Errorf("invalid field in submission: %w", err)
		}
	}

	return nil
}

func validateField(submission *Submission, f form.FieldSpec) error {
	// TODO also add validation of input types!
	if !f.Optional && submission.Values.Get(f.Name) == "" {
		return &errors.ValidationError{
			Message: fmt.Sprintf("missing required field %s", f.Name),
		}
	}

	return nil
}

func NewDefaultValidator(spec *form.Spec) Validator {
	return &DefaultValidator{
		spec: spec,
	}
}
