package common_services

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"reservation-api/internal/dto"
	"strings"
)

type EmailSender interface {
	Send(dto *dto.SendEmailDto) error
}

type EmailService struct {
	Dialer *gomail.Dialer
}

func NewEmailService(host, username, password string, port int) *EmailService {

	return &EmailService{
		// Settings for SMTP server
		Dialer: gomail.NewDialer(host, port, username, password),
	}
}

func (s *EmailService) Send(dto *dto.SendEmailDto) error {

	if strings.TrimSpace(dto.ContentType) == "" {
		dto.ContentType = "text/plain"
	}

	m := gomail.NewMessage()

	m.SetHeader("From", dto.From)

	m.SetHeader("To", dto.To)

	m.SetHeader("Subject", dto.Subject)

	m.SetBody(dto.ContentType, dto.Body)

	dialer := s.Dialer

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
