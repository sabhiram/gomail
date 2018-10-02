package gomail

// Mailer describes an interface that allows sending email messages.
type Mailer interface {
	SendEmail(toAddrs []string, msg []byte) error
}
