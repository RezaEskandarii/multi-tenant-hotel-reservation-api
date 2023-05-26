package dto

type SendEmailRequest struct {
	From        string
	To          string
	Subject     string
	ContentType string
	Body        string
	Attachment  []byte
}
