package common_services

import (
	"reservation-api/internal/config"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/message_broker"
)

type EventService struct {
	MessageBrokerManager message_broker.MessageBrokerManager
	EmailSender          EmailSender
	Logger               applogger.AppLogger
}

func NewEventService(broker message_broker.MessageBrokerManager, emailSender EmailSender) *EventService {

	return &EventService{
		MessageBrokerManager: broker,
		EmailSender:          emailSender,
	}
}

func (e *EventService) SendEmailToGuestOnReservation() {

	e.MessageBrokerManager.Consume(config.ReservationQueueName, func(payload []byte) {

		reservation := utils.ConvertWithGenerics(models.Reservation{}, payload)

		if reservation.Supervisor != nil {
			e.EmailSender.Send(&dto.SendEmailDto{
				From:    "reservationapi@test.test",
				To:      reservation.Supervisor.Email,
				Subject: "reservation",
				Body:    "your reservation completed successfully!",
			})
		}
	})
}
