package form

import (
	"io"
	"ofis/internal/errors"
	"text/template"
)

// HTMLRenderer implements Renderer.
type HTMLRenderer struct {
	tmpl *template.Template
}

// Render renders the form spec into a HTML form and writes it to the Writer.
func (r *HTMLRenderer) Render(form *Spec, w io.Writer) error {
	err := r.tmpl.ExecuteTemplate(w, "form", form)
	if err != nil {
		return &errors.RenderError{
			Message: "could not execute form template with given form spec",
			Err:     err,
		}
	}

	return nil
}

func NewHTMLRenderer(tmpl *template.Template) Renderer {
	return &HTMLRenderer{
		tmpl: tmpl,
	}
}
