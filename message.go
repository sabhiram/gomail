package gomail

// Message describes the collection of data that an email can contain.
type Message struct {
	to      string
	headers map[string]string
	subject string
	body    string
}
