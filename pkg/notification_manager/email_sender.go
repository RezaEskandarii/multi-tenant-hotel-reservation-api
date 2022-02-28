package notification_manager

type EmailSender interface {
	SendEmail(to, cc string, message, attachment []byte) error
}
