package notification_manager

type SmsSender interface {
	Send(to string, message string) error
}
