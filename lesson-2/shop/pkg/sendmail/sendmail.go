package sendmail

import (
	"fmt"
	"net/smtp"
	"shop/models"
)

type sendMail struct {
	from     string
	host     string
	password string
}

type SendMail interface {
	SendMailOrder(order *models.Order) error
}

func NewSentMail(from, host, password string) *sendMail {
	return &sendMail{
		from:     from,
		host:     host,
		password: password,
	}
}

func (s *sendMail) SendMailOrder(order *models.Order) error {

	auth := smtp.PlainAuth("", s.from, s.password, s.host)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Подтверждение заказа номер: %d\r\n\r\nЗаказа %d подвержден.\r\n", s.from, order.Email, order.ID, order.ID)

	if err := smtp.SendMail(s.host+":587", auth, s.from, []string{order.Email}, []byte(msg)); err != nil {
		return err
	}
	return nil

}
