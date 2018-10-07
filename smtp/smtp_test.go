package smtp

import (
	"testing"
)

func TestNewSMTPWrapper(t *testing.T) {
	for _, tc := range []struct {
		addr     string
		from     string
		pass     string
		expError bool
	}{
		// Good cases.
		{"foo:44", "a@b.com", "password", false},

		// Bad addresses
		{"foo::44", "a@b.com", "password", true},
		{"foo:4o4", "a@b.com", "password", true},
		{"foo", "a@b.com", "password", true},

		// Bad email addr
		{"foo:22", "b.com", "password", true},
	} {
		_, err := New(tc.addr, tc.from, tc.pass)
		if tc.expError {
			if err == nil {
				t.Fatalf("Error expected, got valid return\n")
			}
		}
	}
}
