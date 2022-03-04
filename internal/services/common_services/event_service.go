package common_services

import (
	"fmt"
	"reservation-api/internal/config"
	"reservation-api/pkg/message_broker"
)

type EventService struct {
	MessageBrokerManager message_broker.MessageBrokerManager
	EmailSender          EmailSender
}

func NewEventService(broker message_broker.MessageBrokerManager, emailSender EmailSender) *EventService {

	return &EventService{
		MessageBrokerManager: broker,
		EmailSender:          emailSender,
	}
}

func (e *EventService) SendEmailToGuestOnReservation() {

	e.MessageBrokerManager.Consume(config.ReservationQueueName, func(payload []byte) {
		fmt.Println(string(payload))
	})
}
