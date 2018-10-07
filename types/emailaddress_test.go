package types

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
				t.Fatalf("Error expected, got valid return\n")
			}
			continue
		}
		if e.local != tc.local {
			t.Fatalf("local mismatch : expected: %s, actual: %s", tc.local, e.local)
		}
		if e.domain != tc.domain {
			t.Fatalf("domain mismatch : expected: %s, actual: %s", tc.domain, e.domain)
		}
		if e.String() != tc.email {
			t.Fatalf("String() mismatch : expected: %s, actual: %s", tc.email, e.String())
		}
	}
}

func TestNewEmailAddresses(t *testing.T) {
	for _, tc := range []struct {
		emails   []string
		expError bool
		expStr   string
		expLen   int
	}{
		// Good cases.
		{[]string{"a@b.com", "a@c.com"}, false, "a@b.com;a@c.com", 2},

		// Malformed email addresses.
		{[]string{"ab.com"}, true, "", 0},
	} {
		es, err := NewEmailAddresses(tc.emails...)
		if tc.expError {
			if err == nil {
				t.Fatalf("Error expected, got valid return\n")
			}
		}
		if es.String() != tc.expStr {
			t.Fatalf("String() mismatch : expected: %s, actual: %s", tc.expStr, es.String())
		}
		if len(es.All()) != tc.expLen {
			t.Fatalf("All() mismatch : expected: %d, actual: %d", tc.expLen, len(es.All()))
		}
	}
}
