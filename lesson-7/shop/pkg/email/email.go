package email

import (
	"net/smtp"
	"shop/models"
)

type EmailClient interface {
	SendOrderConfirmation(order *models.Order) error
}

type emailClient struct {
	cli *smtp.Client
}

func (email *emailClient) SendOrderConfirmation(order *models.Order) error {
	// ИЩИТЕ В ССЫЛКЕ К ЗАДАНИЮ
	return nil
}
