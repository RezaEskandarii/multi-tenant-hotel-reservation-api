package test

import (
	"github.com/stretchr/testify/assert"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
	"hotel-reservation/internal/services"
	"hotel-reservation/pkg/database"
	"testing"
)

var (
	db, _      = database.GetDb(true)
	pagination = dto.PaginationInput{
		Page:    1,
		PerPage: 20,
	}
	countryService = *services.NewCountryService()
)

func init() {
	countryService.Repository = repositories.NewCountryRepository(db)
}

func TestCanCreateNewCountry(t *testing.T) {

	country1 := models.Country{
		Name:  "Iran",
		Alias: "IRI",
	}

	country2, err := countryService.Create(&country1)

	assert.Nil(t, err)
	assert.NotNil(t, country2)
	assert.Equal(t, country1.Id, country2.Id)
	assert.NotEqual(t, country2, 0)

}

func TestCanFindCountry(t *testing.T) {

	country1 := models.Country{
		Name:  "Canada",
		Alias: "Ca",
	}

	country2, err := countryService.Create(&country1)

	assert.Nil(t, err)
	assert.NotNil(t, country2)
	assert.NotEqual(t, country2, 0)

	country3, err := countryService.Find(country2.Id)
	assert.Nil(t, err)
	assert.NotNil(t, country3)
	assert.NotEqual(t, country3.Id, 0)
	assert.Equal(t, country2.Name, country3.Name)
}

func TestCanUpdateCountry(t *testing.T) {
	country1 := models.Country{
		Name:  "Canada",
		Alias: "Ca",
	}
	country2, err := countryService.Create(&country1)

	assert.Nil(t, err)
	assert.NotNil(t, country2)
	assert.NotEqual(t, country2, 0)

	country2.Name = "Iran"
	country2.Alias = "IRI"

	country3, err := countryService.Update(country2)

	assert.Nil(t, err)
	assert.NotNil(t, country3)
	assert.NotEqual(t, country3.Id, 0)
	assert.Equal(t, country2.Name, country3.Name)
	assert.Equal(t, country2.Alias, country3.Alias)
	assert.Equal(t, country3.Name, "Iran")
	assert.Equal(t, country3.Alias, "IRI")
}

func TestCanFindAllCountries(t *testing.T) {
	list, err := countryService.FindAll(&pagination)

	assert.Nil(t, err)
	assert.NotNil(t, list)
}
