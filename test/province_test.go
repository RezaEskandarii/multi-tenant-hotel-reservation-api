package test

import (
	"github.com/stretchr/testify/assert"
	"reservation-api/internal/models"
	"testing"
)

func init() {
	setRepository()
}

func TestCanCreateNewProvince(t *testing.T) {

	country, err := countryService.Create(&models.Country{Name: "Iran"})

	assert.Nil(t, err)
	assert.NotNil(t, country)

	province := models.Province{
		Name:      "Mazandaran",
		Alias:     "mz",
		CountryId: country.Id,
	}

	assert.Nil(t, err)

	actualProvince, err := provinceService.Create(&province)

	assert.Nil(t, err)
	assert.NotNil(t, actualProvince)
	assert.Equal(t, province.Id, actualProvince.Id)
	assert.NotEqual(t, actualProvince, 0)

}

func TestCanFindProvince(t *testing.T) {

	province := models.Province{
		Name:  "Tehran",
		Alias: "th",
	}

	actualProvince, err := provinceService.Create(&province)

	assert.Nil(t, err)
	assert.NotNil(t, actualProvince)
	assert.NotEqual(t, actualProvince, 0)

	exceptedProvince, err := provinceService.Find(actualProvince.Id)
	assert.Nil(t, err)
	assert.NotNil(t, exceptedProvince)
	assert.NotEqual(t, exceptedProvince.Id, 0)
	assert.Equal(t, exceptedProvince.Name, actualProvince.Name)
}

func TestCanUpdateProvince(t *testing.T) {
	province := models.Province{
		Name:  "NewYork",
		Alias: "Ny",
	}

	exceptedProvince, err := provinceService.Create(&province)

	assert.Nil(t, err)
	assert.NotNil(t, exceptedProvince)
	assert.NotEqual(t, exceptedProvince, 0)

	exceptedProvince.Name = "Texas"
	exceptedProvince.Alias = "tx"

	actualProvince, err := provinceService.Update(exceptedProvince)

	assert.Nil(t, err)
	assert.NotNil(t, actualProvince)
	assert.NotEqual(t, actualProvince.Id, 0)
	assert.Equal(t, exceptedProvince.Name, actualProvince.Name)
	assert.Equal(t, exceptedProvince.Alias, actualProvince.Alias)
	assert.Equal(t, actualProvince.Name, "Texas")
	assert.Equal(t, actualProvince.Alias, "tx")
}

//func TestCanFindAllProvinces(t *testing.T) {
//	list, err := provinceService.FindAll(&pagination)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, list)
//}
