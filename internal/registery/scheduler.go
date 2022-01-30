package registery

import (
	"context"
	"github.com/procyon-projects/chrono"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/applogger"
)

// schedule remove expired reservation requests job every night.
func scheduleRemoveExpiredReservationRequests(s *domain_services.ReservationService, logger applogger.Logger) error {

	taskScheduler := chrono.NewDefaultTaskScheduler()
	pattern := "0 0 * * *" // “At 00:00.”

	_, err := taskScheduler.ScheduleWithCron(func(ctx context.Context) {

		if err := s.RemoveExpiredReservationRequests(); err != nil {
			logger.LogError(err.Error())
		}

	}, pattern, nil)

	return err
}
