package gomail

import (
	"testing"
)

func TestEmailAddress(t *testing.T) {
	for _, tc := range []struct {
		email    string
		local    string
		domain   string
		expError bool
	}{
		// Good cases.
		{"a@b.com", "a", "b.com", false},

		// Malformed email addresses.
		{"ab.com", "", "", true},
		{"a@", "", "", true},
		{"@b.com", "", "", true},

		// Emails with invalid spaces.
		{" a@b.com", "", "", true},
		{"a@b.com ", "", "", true},
		{"a @b.com ", "", "", true},
		{"a@ b.com ", "", "", true},
		{"a@ ", "", "", true},
		{" @b.com", "", "", true},
	} {
		e, err := NewEmailAddress(tc.email)
		if tc.expError {
			if err == nil {
				t.Fatalf("Error expected - email: %s", tc.email)
			}
			continue
		}
		if e.local != tc.local {
			t.Fatalf("local parts don't match - expected: %s, actual: %s", tc.local, e.local)
		}
		if e.domain != tc.domain {
			t.Fatalf("domain parts don't match - expected: %s, actual: %s", tc.domain, e.domain)
		}
		if e.String() != tc.email {
			t.Fatalf("reconstituted email doesn't match - expected: %s, actual: %s", tc.email, e.String())
		}
	}
}
