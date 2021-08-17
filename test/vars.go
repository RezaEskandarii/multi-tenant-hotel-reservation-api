package test

import (
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/repositories"
	"hotel-reservation/internal/services"
	"hotel-reservation/pkg/database"
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
