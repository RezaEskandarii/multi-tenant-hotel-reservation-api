package registery

import (
	"github.com/jasonlvhit/gocron"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/applogger"
)

// schedule remove expired reservation requests job every night.
func scheduleRemoveExpiredReservationRequests(s *domain_services.ReservationService, logger applogger.Logger) error {

	task := func() {
		if err := s.RemoveExpiredReservationRequests(); err != nil {
			logger.LogError(err.Error())
		}
	}

	return gocron.Every(1).Hour().Do(task)

}
