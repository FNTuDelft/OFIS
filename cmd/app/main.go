package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"time"

	perrors "ofis/internal/errors"
	"ofis/internal/form"
	"ofis/internal/middleware"
	"ofis/internal/submission"
)

// htmlTemplate is the HTML template which the input form template is supposed to be build into.
var htmlTemplate *template.Template = form.HTMLTemplate()

func main() {
	mux := http.NewServeMux()

	handler := middleware.Logging(
		middleware.Error(Handle),
	)

	mux.Handle("/", handler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on %s", server.Addr)

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("error while listening: %s", err)
		}
	}()

	<-quit
	log.Println("Server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Panicf("server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

// Handle handles the incoming form request.
func Handle(w http.ResponseWriter, r *http.Request) error {
	// inputSpec is the input form template parsed into a generic form template model form.Spec.
	inputSpec := form.InputSpec("./templates/form_template.xml")

	switch r.Method {
	case http.MethodGet:
		// The GET request is handled by rendering the HTML form template which we build from
		// the input form template.
		err := form.NewHTMLRenderer(htmlTemplate).Render(inputSpec, w)
		if err != nil {
			return &perrors.HTTPError{
				Code:     http.StatusInternalServerError,
				Message:  "could not render HTML form template",
				Internal: err,
			}
		}

		return nil
	case http.MethodPost:
		// The POST request is handled by parsing and validating the form submission and
		// rendering the results.
		err := submission.NewService(
			submission.NewParser(),
			submission.NewDefaultValidator(inputSpec),
			submission.NewPDFRenderer(),
		).HandleFormSubmission(r, w)
		if err != nil {
			return &perrors.HTTPError{
				Code:     http.StatusInternalServerError,
				Message:  "could not handle form submission",
				Internal: err,
			}
		}

		return nil
	default:
		return &perrors.HTTPError{
			Code:     http.StatusMethodNotAllowed,
			Message:  r.Method + " does not have a matching handler",
			Internal: nil,
		}
	}
}
