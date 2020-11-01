package domain

// Report comes from the client
type Report struct {
	url    string
	message    string
}

// NewReport returns the Report struct
func NewReport(url string, message string) Report {
	return Report{url: url, message: message}
}

// Url returns the report URL
func (r Report) Url() string {
	return r.url
}

// Message returns the report message
func (r Report) Message() string {
	return r.message
}
