package test

import (
	"reservation-api/internal/dto"
	"reservation-api/internal/repositories"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/database"
)

var (
	db, _      = database.GetDb(true)
	pagination = dto.PaginationFilter{
		Page:    1,
		PerPage: 20,
	}

	countryService  = domain_services.NewCountryService()
	provinceService = domain_services.NewProvinceService()
)

func setRepository() {
	countryService.Repository = repositories.NewCountryRepository(db)
	provinceService.Repository = repositories.NewProvinceRepository(db)
}
