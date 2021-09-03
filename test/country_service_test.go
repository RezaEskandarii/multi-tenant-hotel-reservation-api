package test

import (
	"github.com/stretchr/testify/assert"
	"hotel-reservation/internal/models"
	"testing"
)

func init() {

	setRepository()
}

func TestCanCreateNewCountry(t *testing.T) {

	country := models.Country{
		Name:  "Iran",
		Alias: "IRI",
	}

	actualCountry, err := countryService.Create(&country)

	assert.Nil(t, err)
	assert.NotNil(t, actualCountry)
	assert.Equal(t, actualCountry.Id, country.Id)
	assert.Equal(t, actualCountry.Name, country.Name)
	assert.Equal(t, actualCountry.Alias, country.Alias)
	assert.NotEqual(t, actualCountry, 0)

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
	country := models.Country{
		Name:  "Canada",
		Alias: "Ca",
	}

	actualCountry, err := countryService.Create(&country)

	assert.Nil(t, err)
	assert.NotNil(t, actualCountry)
	assert.NotEqual(t, actualCountry, 0)

	actualCountry.Name = "Iran"
	actualCountry.Alias = "IRI"

	exceptedCountry, err := countryService.Update(actualCountry)

	assert.Nil(t, err)
	assert.NotNil(t, exceptedCountry)
	assert.NotEqual(t, exceptedCountry.Id, 0)
	assert.Equal(t, exceptedCountry.Name, actualCountry.Name)
	assert.Equal(t, exceptedCountry.Alias, actualCountry.Alias)
	assert.Equal(t, exceptedCountry.Name, "Iran")
	assert.Equal(t, exceptedCountry.Alias, "IRI")
}

func TestCanFindAllCountries(t *testing.T) {
	list, err := countryService.FindAll(&pagination)

	assert.Nil(t, err)
	assert.NotNil(t, list)
}
