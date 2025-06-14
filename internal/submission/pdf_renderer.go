package submission

import (
	"fmt"
	"io"
	"net/url"
	"ofis/internal/errors"

	"github.com/jung-kurt/gofpdf"
)

// PDFRenderer implements Renderer.
type PDFRenderer struct{}

// Render writes the submission into PDF to the writer.
func (r *PDFRenderer) Render(submission *Submission, w io.Writer) error {
	// Initialize A4 portrait PDF
	pdf := gofpdf.New("P", "mm", "A4", "Arial")
	defer pdf.Close()

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// Title
	pdf.CellFormat(0, 10, "Submission Confirmation", "", 1, "C", false, 0, "")

	// List each submitted field name and value.
	renderValues(pdf, submission.Values)

	// Add all filenames to the PDF and add the actual files as attachments.
	renderFiles(pdf, submission.Files)

	// Stream PDF to writer
	err := pdf.Output(w)
	if err != nil {
		return &errors.RenderError{
			Message: "could not output PDF with submission results",
			Err:     err,
		}
	}

	return nil
}

// renderValues adds each submitted field name and value.
func renderValues(pdf *gofpdf.Fpdf, values url.Values) {
	for name, vals := range values {
		pdf.CellFormat(40, 8, fmt.Sprintf("%s:", name), "", 0, "L", false, 0, "")

		// Field values (may be multi-valued)
		for _, v := range vals {
			pdf.MultiCell(0, 8, v, "", "L", false)
		}
	}
}

// renderFiles adds all filenames to the PDF and add the actual files as attachments.
func renderFiles(pdf *gofpdf.Fpdf, fileSubmission FileSubmissions) {
	attachments := make([]gofpdf.Attachment, 0)
	for fieldname, files := range fileSubmission {
		pdf.CellFormat(40, 8, fieldname+":", "", 0, "L", false, 0, "")

		// Each file submission can have multiple files.
		for filename, file := range files {
			attachments = append(attachments, gofpdf.Attachment{
				Content:  file,
				Filename: filename,
			})

			pdf.MultiCell(0, 8, filename+" (See attachments)", "", "L", false)
		}
	}

	pdf.SetAttachments(attachments)
}

func NewPDFRenderer() Renderer {
	return &PDFRenderer{}
}
