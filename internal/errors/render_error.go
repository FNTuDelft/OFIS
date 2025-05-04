package errors

type RenderError struct {
	Message string
	Err     error
}

func (e *RenderError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}

	return e.Message
}
