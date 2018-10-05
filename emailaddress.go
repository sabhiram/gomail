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

type EmailAddresses []*EmailAddress
