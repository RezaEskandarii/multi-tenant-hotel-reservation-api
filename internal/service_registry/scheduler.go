package service_registry

import (
	"context"
	"github.com/jasonlvhit/gocron"
	"reservation-api/internal/global_variables"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/applogger"
	"time"
)

// schedule remove expired reservation requests job every night.
func scheduleRemoveExpiredReservationRequests(s *domain_services.ReservationService, logger applogger.Logger,
	tenantService *domain_services.TenantService) {

	tenants, err := tenantService.GetAll()

	if err != nil {
		logger.LogError(err)

	} else {

		task := func() {
			for _, tenant := range tenants {

				parentCtx := context.Background()
				ctx := context.WithValue(parentCtx, global_variables.TenantIDKey, tenant.Id)

				if err := s.RemoveExpiredReservationRequests(ctx); err != nil {
					logger.LogError(err.Error())
				}

				time.Sleep(3000)
			}
		}

		err := gocron.Every(1).Hour().Do(task)
		if err != nil {
			logger.LogError(err.Error())
		}
	}

}
