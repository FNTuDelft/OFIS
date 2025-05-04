package form

// Spec is the generic, language-agnostic description of a form.
type Spec struct {
	Elements []interface{}
}

// FieldSpec describes a single form input.
type FieldSpec struct {
	Name       string            // input name attribute
	Type       string            // e.g., "Text", "Enumeration"
	FieldType  string            // e.g., "Select", "TextBox", "File"
	Optional   bool              // true if field is optional
	Caption    string            // label or caption text
	Attributes map[string]string // any extra attributes from JSON/XML
	Options    []OptionSpec      // for select/enumeration types
}

// OptionSpec represents a single option in a select input.
type OptionSpec struct {
	Name  string
	Value string
}

// SectionSpec groups multiple fields under a common title.
type SectionSpec struct {
	Title  string
	Fields []FieldSpec
}
