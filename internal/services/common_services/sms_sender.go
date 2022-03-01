package common_services

type SmsSender interface {
	Send(to string, message string) error
}
