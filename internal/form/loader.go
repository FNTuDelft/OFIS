package form

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

// InputSpec loads the input form template and parses it into a Spec.
func InputSpec(templatePath string) *Spec {
	fileBytes, err := os.ReadFile(templatePath)
	if err != nil {
		log.Panicf("can't read form template: %v", err)
	}

	if len(fileBytes) == 0 {
		log.Panic("form template is empty")
	}

	parser := NewParserBuilder().Build(templatePath)
	spec, err := parser.Parse(bytes.NewReader(fileBytes))
	if err != nil || spec == nil {
		log.Panicf("Could not parse XML form template %v", err)
	}

	return spec
}

// HTMLTemplate loads the HTML form template that is used to translate the input form template into.
// It can take multiple file paths as the HTML template might be build up of multiple parts that life in
// separate files.
func HTMLTemplate() *template.Template {
	form, err := template.New("form").
		Funcs(template.FuncMap{
			"getElementType": getElementType,
		}).
		ParseFiles(
			"./templates/form_template.html",
			"./templates/field_template.html",
			"./templates/section_template.html",
		)
	t := template.Must(form, err)
	if t == nil {
		log.Panic("HTML form template could not be found")
	}

	return t
}

// Helper function for the form templates to help determine what type an element is.
func getElementType(e interface{}) string {
	switch e.(type) {
	case FieldSpec:
		return "Field"
	case SectionSpec:
		return "Section"
	default:
		return "Unknown"
	}
}
