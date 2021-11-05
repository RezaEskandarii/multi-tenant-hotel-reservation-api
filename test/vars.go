package test

import (
	"reservation-api/internal/dto"
	"reservation-api/internal/repositories"
	"reservation-api/internal/services"
	"reservation-api/pkg/database"
)

var (
	db, _      = database.GetDb(true)
	pagination = dto.PaginationInput{
		Page:    1,
		PerPage: 20,
	}

	countryService  = services.NewCountryService()
	provinceService = services.NewProvinceService()
)

func setRepository() {
	countryService.Repository = repositories.NewCountryRepository(db)
	provinceService.Repository = repositories.NewProvinceRepository(db)
}
