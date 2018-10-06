package gomail

import (
	"errors"
	"fmt"

	"github.com/sabhiram/gomail/types"
)

// Mailer describes an interface that allows sending email messages.
type Mailer interface {
	SendEmail(toAddrs []string, msg []byte) error
}

// Message describes the collection of data that an email can contain.
type Message struct {
	to      types.EmailAddresses
	headers map[string]string
	subject string
	body    string
}

// NewMessage encapsulates an email message along with it's subject and
// recipients.
func NewMessage(to types.EmailAddresses, subject, body string) (*Message, error) {
	if to == nil {
		return nil, errors.New("no recipients specified for message")
	}

	return &Message{
		to:      to,
		headers: map[string]string{},
		subject: subject,
		body:    body,
	}, nil
}

// SetHeader sets or overwrites the specified header key, value pair into the
// email headers that will be emitted with the message.
func (m *Message) SetHeader(k, v string) {
	m.headers[k] = v
}

// String returns a string version of the email message. Adheres to the stringer
// interface.
func (m *Message) String() string {
	m.SetHeader("To", m.to.String())
	m.SetHeader("Subject", m.subject)

	s := ""
	for k, v := range m.headers {
		s += fmt.Sprintf("%s:%s\r\n", k, v)
	}
	return s + fmt.Sprintf("\r\n%s\r\n", m.body)
}
