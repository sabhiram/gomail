package gomail

import (
	"errors"
	"fmt"
	"time"

	"github.com/sabhiram/gomail/types"
)

// Mailer describes an interface that allows sending email messages.
type Mailer interface {
	SendEmail(toAddrs []string, msg []byte) error
}

// Messager describes an interface to send email messages.
type Messager interface {
	String() string
}

// Message describes the collection of data that an email can contain and
// adheres to the `Messager` interface.
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

// MultipartSection encapsulates a single section in a multipart email.
type MultipartSection struct {
	headers map[string]string // headers specific to this section
	content []byte            // section content
}

func NewMultipartSection(headers map[string]string, content []byte) *MultipartSection {
	headers["Content-Disposition"] = `inline`
	return &MultipartSection{
		headers: headers,
		content: content,
	}
}

func NewMultipartTextSection(content []byte) *MultipartSection {
	return NewMultipartSection(map[string]string{
		"Content-Type":              `text/plain; charset="utf-8"`,
		"Content-Transfer-Encoding": `quoted-printable`,
	}, content)
}

func NewMultipartHTMLSection(content []byte) *MultipartSection {
	return NewMultipartSection(map[string]string{
		"Content-Type": `text/html; charset="utf-8"`,
	}, content)
}

func (ms *MultipartSection) String() string {
	var s string
	for k, v := range ms.headers {
		s += fmt.Sprintf("%s: %s\n", k, v)
	}
	if len(s) > 0 {
		s += "\n"
	}
	return s + string(ms.content) + "\n"
}

type MultipartMessage struct {
	*Message
	boundary string
	sections []*MultipartSection
}

func (mp *MultipartMessage) SetBoundary(boundary string) {
	mp.boundary = boundary
}

func NewMultipartMessage(to types.EmailAddresses, subject string, sections []*MultipartSection) (Messager, error) {
	if len(sections) <= 0 {
		return nil, errors.New("no multipart sections specified")
	}
	m, err := NewMessage(to, subject, "")
	if err != nil {
		return nil, err
	}

	return &MultipartMessage{
		Message:  m,
		boundary: fmt.Sprintf("--__%s__", time.Now().String()),
		sections: sections,
	}, nil
}

func (mp *MultipartMessage) String() string {
	msg := mp.boundary + "\n"
	for _, section := range mp.sections {
		msg += section.String() + "\n"
		msg += mp.boundary + "\n"
	}
	return msg
}
