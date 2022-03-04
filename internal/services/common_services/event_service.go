package common_services

import (
	"fmt"
	"reservation-api/internal/config"
	"reservation-api/pkg/message_broker"
)

type EventService struct {
	MessageBrokerManager message_broker.MessageBrokerManager
	EmailService         *EmailService
}

func NewEventService(broker message_broker.MessageBrokerManager, emailService *EmailService) *EventService {

	return &EventService{
		MessageBrokerManager: broker,
		EmailService:         emailService,
	}
}

func (e *EventService) SendEmailToGuestOnReservation() {

	e.MessageBrokerManager.Consume(config.ReservationQueueName, func(payload interface{}) {
		fmt.Println(payload)
	})
}
