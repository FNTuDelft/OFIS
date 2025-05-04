package middleware

import (
	"errors"
	"log"
	"net/http"

	perrors "ofis/internal/errors"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

// Error is middleware which can be used on a http handler function to properly log errors.
func Error(next HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err != nil {
			log.Printf("handler error: %s", err.Error())

			var httpErr *perrors.HTTPError
			if errors.As(err, &httpErr) {
				writeHTTPError(w, httpErr.Code)

				return
			}

			writeHTTPError(w, http.StatusInternalServerError)
		}
	}
}

func writeHTTPError(w http.ResponseWriter, code int) {
	status := http.StatusText(code)
	if status == "" {
		status = http.StatusText(http.StatusInternalServerError)
		code = http.StatusInternalServerError
	}

	http.Error(w, status, code)
}
