package errors

type HTTPError struct {
	Code     int
	Message  string
	Internal error
}

func (e *HTTPError) Error() string {
	if e.Internal != nil {
		return e.Message + ": " + e.Internal.Error()
	}

	return e.Message
}
