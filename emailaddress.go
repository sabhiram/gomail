package gomail

import (
	"errors"
	"fmt"
	"strings"
)

// EmailAddress represents an email address.
type EmailAddress struct {
	local  string
	domain string
}

// NewEmailAddress returns a new instance of an email address or an error.
func NewEmailAddress(s string) (*EmailAddress, error) {
	if len(s) != len(strings.TrimSpace(s)) {
		return nil, errors.New("email address cannot contain spaces")
	}

	parts := strings.Split(s, "@")
	if len(parts) != 2 {
		return nil, errors.New("email contains more than one \"@\"")
	}

	local, domain := parts[0], parts[1]
	if len(local) == 0 || len(domain) == 0 {
		return nil, errors.New("invalid local / domain specified")
	}

	return &EmailAddress{
		local:  local,
		domain: domain,
	}, nil
}

// String implements the stringer interface.
func (e *EmailAddress) String() string {
	return fmt.Sprintf("%s@%s", e.local, e.domain)
}

// EmailAddresses represents a list of email addresses.
type EmailAddresses []*EmailAddress

// All returns a list of strings of all underlying email addresses.
func (ee EmailAddresses) All() []string {
	r := []string{}
	for _, e := range ee {
		r = append(r, e.String())
	}
	return r
}

// String implements the stringer interface for a list of email addresses.
func (ee EmailAddresses) String() string {
	if ee == nil || len(ee) == 0 {
		return ""
	}

	return strings.Join(ee.All(), ";")
}
