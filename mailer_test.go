package gomail

import (
	"strings"
	"testing"

	"github.com/sabhiram/gomail/types"
)

func TestNewMessage(t *testing.T) {
	m, err := NewMessage(nil, "SUBJECT", "BODY")
	if err == nil {
		t.Fatalf("Expected bad message due to nil recipients\n")
	}

	to, err := types.NewEmailAddresses("a@b.com", "b@b.com")
	if err != nil {
		t.Fatalf("Expected nil error, got %s\n", err.Error())
	}
	m, err = NewMessage(to, "SUBJECT", "BODY")
	if err != nil {
		t.Fatalf("Expected nil error, got %s\n", err.Error())
	}

	m.SetHeader("KEY", "VALUE")
	s := m.String()

	if strings.Index(s, "To:a@b.com;b@b.com\r\n") < 0 {
		t.Fatalf("Message does not have recipient email addresses.\n")
	}
	if strings.Index(s, "Subject:SUBJECT\r\n") < 0 {
		t.Fatalf("Message does not have subject.\n")
	}
	if strings.Index(s, "KEY:VALUE\r\n") < 0 {
		t.Fatalf("Message does not have custom header.\n")
	}
	if strings.Index(s, "\r\nBODY\r\n") < 0 {
		t.Fatalf("Message does not have body content.\n")
	}
}
