package sendmail

import (
	"fmt"
	"net/smtp"
)

type Sendmail struct {
	from string
	host string
	auth smtp.Auth
}

func NewSendmail(from, host string, auth smtp.Auth) *Sendmail {
	return &Sendmail{from, host, auth}
}

func (s Sendmail) Send(to, messages string) error {
	return smtp.SendMail(fmt.Sprintf("%s:%d", s.host, 25), s.auth, s.from, []string{to}, []byte(messages))
}
