package submission

import (
	"fmt"
	"io"
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
	for name, vals := range submission.Values {
		pdf.CellFormat(40, 8, fmt.Sprintf("%s:", name), "", 0, "L", false, 0, "")

		// Field values (may be multi-valued)
		for _, v := range vals {
			pdf.MultiCell(0, 8, v, "", "L", false)
		}
	}

	// Add all filenames to the PDF and add the actual files as attachments.
	attachments := make([]gofpdf.Attachment, 0)
	for fieldname, files := range submission.Files {
		pdf.CellFormat(40, 8, fieldname+":", "", 0, "L", false, 0, "")

		for filename, file := range files {
			attachments = append(attachments, gofpdf.Attachment{
				Content:  file,
				Filename: filename,
			})

			pdf.MultiCell(0, 8, filename+" (See attachments)", "", "L", false)
		}
	}

	pdf.SetAttachments(attachments)

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

func NewPDFRenderer() Renderer {
	return &PDFRenderer{}
}
