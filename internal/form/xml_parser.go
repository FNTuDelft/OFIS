package form

import (
	"encoding/xml"
	"fmt"
	"io"

	"ofis/internal/errors"
)

// XMLParser implements Parser.
type XMLParser struct{}

// Helper types matching the XML structure
type xmlLabel struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:",chardata"`
}
type xmlField struct {
	Name      string     `xml:"Name,attr"`
	Type      string     `xml:"Type,attr"`
	Optional  string     `xml:"Optional,attr"`
	FieldType string     `xml:"FieldType,attr"`
	Caption   string     `xml:"Caption"`
	Labels    []xmlLabel `xml:"Labels>Label"`
}
type xmlSection struct {
	Title    string     `xml:"Title"`
	Contents []xmlField `xml:"Contents>Field"`
}
type xmlForm struct {
	Elements []interface{}
}

// UnmarshalXML decodes XML into xmlForm. This custom method is needed to retain ordering of Elements
// while elements are decoded into their right types.
func (f *xmlForm) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return fmt.Errorf("error reading token: %w", err)
		}
		if token == nil {
			break
		}

		switch t := token.(type) {
		case xml.StartElement:
			err := f.decodeElement(t.Name.Local, d, &t)
			if err != nil {
				return fmt.Errorf("error decoding an element: %w", err)
			}
		case xml.EndElement:
			if t.Name.Local == start.Name.Local {
				return nil
			}
		}
	}

	return nil
}

// decodeElement Decodes Elements into their specific types while retaining their order (independant of type).
func (f *xmlForm) decodeElement(b string, d *xml.Decoder, t *xml.StartElement) error {
	switch b {
	case "Field":
		var field xmlField
		if err := d.DecodeElement(&field, t); err != nil {
			return fmt.Errorf("error decoding <Field>: %w", err)
		}

		f.Elements = append(f.Elements, field)
	case "Section":
		var section xmlSection
		if err := d.DecodeElement(&section, t); err != nil {
			return fmt.Errorf("error decoding <Section>: %w", err)
		}

		f.Elements = append(f.Elements, section)
	}

	return nil
}

// Parse decodes the XML form template into custom models, which are then parsed into the
// generic form model Spec.
func (p *XMLParser) Parse(formTemplate io.Reader) (*Spec, error) {
	// Decode into the xmlForm
	var xf xmlForm
	err := xml.NewDecoder(formTemplate).Decode(&xf)
	if err != nil {
		return nil, &errors.ParseError{
			Message: "could not decode XML form template into XML form model",
			Err:     err,
		}
	}

	// Map to generic Spec
	spec := &Spec{}
	for _, elem := range xf.Elements {
		switch e := elem.(type) {
		case xmlField:
			mapField(spec, &e)
		case xmlSection:
			mapSection(spec, &e)
		default:
		}
	}

	return spec, nil
}

func mapField(spec *Spec, f *xmlField) {
	field := field(f)
	spec.Elements = append(spec.Elements, field)
}

func field(f *xmlField) FieldSpec {
	field := FieldSpec{
		Name:      f.Name,
		Type:      f.Type,
		FieldType: f.FieldType,
		Optional:  (f.Optional == "True" || f.Optional == "true"),
		Caption:   f.Caption,
	}

	for _, lbl := range f.Labels {
		field.Options = append(field.Options, OptionSpec(lbl))
	}

	return field
}

func mapSection(spec *Spec, s *xmlSection) {
	sec := SectionSpec{Title: s.Title}

	for _, f := range s.Contents {
		field := field(&f)
		sec.Fields = append(sec.Fields, field)
	}

	spec.Elements = append(spec.Elements, sec)
}

func NewXMLParser() Parser {
	return &XMLParser{}
}
