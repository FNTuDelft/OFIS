package submission

import "net/url"

type Submission struct {
	Values url.Values
	Files  FileSubmissions
}

type FileSubmissions map[string]map[string][]byte

func NewSubmission(values url.Values, files FileSubmissions) *Submission {
	return &Submission{
		Values: values,
		Files:  files,
	}
}
