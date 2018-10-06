package smtp

import (
	"errors"
	"net"
	"net/smtp"
	"strconv"

	"github.com/sabhiram/gomail"
	"github.com/sabhiram/gomail/types"
)

// SMTP encapsulates all the settings needed to send an email via the `Simple
// Mail Transfer Protocol`: https://tools.ietf.org/html/rfc5321
type SMTP struct {
	addr string              // SMTP server hostname
	host string              // SMTP server address
	port int                 // SMTP server port
	from *types.EmailAddress // Sender email address
	auth smtp.Auth           // SMTP simple auth interface
}

// New returns an instance of the SMTP structure which adheres to the `Mailer`
// interface.
func New(addr, from, pass string) (gomail.Mailer, error) {
	h, _p, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	p, err := strconv.Atoi(_p)
	if err != nil {
		return nil, err
	}

	f, err := types.NewEmailAddress(from)
	if err != nil {
		return nil, err
	}

	return &SMTP{
		addr: addr,
		host: h,
		port: p,
		from: f,
		auth: smtp.PlainAuth("", f.String(), pass, h),
	}, nil
}

// SendEmail sends an email to all the addresses listed in the `to` slice. The
// payload of the message is `msg`.
func (s *SMTP) SendEmail(to []string, msg []byte) error {
	if to == nil || len(to) == 0 {
		return errors.New("no recipients specified")
	}
	if msg == nil {
		return errors.New("invalid message specified")
	}

	var recipients types.EmailAddresses
	for _, ea := range to {
		if e, err := types.NewEmailAddress(ea); err == nil {
			recipients = append(recipients, e)
		}
	}
	if len(recipients) == 0 {
		return errors.New("no valid recipient emails specified")
	}

	return smtp.SendMail(s.addr, s.auth, s.from.String(), recipients.All(), msg)
}
