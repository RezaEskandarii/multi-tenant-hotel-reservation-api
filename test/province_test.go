package test

import (
	"github.com/stretchr/testify/assert"
	"hotel-reservation/internal/models"
	"testing"
)

func init() {
	setRepository()
}

func TestCanCreateNewProvince(t *testing.T) {

	country, err := countryService.Create(&models.Country{Name: "Iran"})

	assert.Nil(t, err)
	assert.NotNil(t, country)

	Province1 := models.Province{
		Name:      "Mazandaran",
		Alias:     "mz",
		CountryId: country.Id,
	}

	assert.Nil(t, err)

	Province2, err := provinceService.Create(&Province1)

	assert.Nil(t, err)
	assert.NotNil(t, Province2)
	assert.Equal(t, Province1.Id, Province2.Id)
	assert.NotEqual(t, Province2, 0)

}

func TestCanFindProvince(t *testing.T) {

	Province1 := models.Province{
		Name:  "Tehran",
		Alias: "th",
	}

	Province2, err := provinceService.Create(&Province1)

	assert.Nil(t, err)
	assert.NotNil(t, Province2)
	assert.NotEqual(t, Province2, 0)

	Province3, err := provinceService.Find(Province2.Id)
	assert.Nil(t, err)
	assert.NotNil(t, Province3)
	assert.NotEqual(t, Province3.Id, 0)
	assert.Equal(t, Province2.Name, Province3.Name)
}

func TestCanUpdateProvince(t *testing.T) {
	Province1 := models.Province{
		Name:  "NewYork",
		Alias: "Ny",
	}
	Province2, err := provinceService.Create(&Province1)

	assert.Nil(t, err)
	assert.NotNil(t, Province2)
	assert.NotEqual(t, Province2, 0)

	Province2.Name = "Texas"
	Province2.Alias = "tx"

	Province3, err := provinceService.Update(Province2)

	assert.Nil(t, err)
	assert.NotNil(t, Province3)
	assert.NotEqual(t, Province3.Id, 0)
	assert.Equal(t, Province2.Name, Province3.Name)
	assert.Equal(t, Province2.Alias, Province3.Alias)
	assert.Equal(t, Province3.Name, "Texas")
	assert.Equal(t, Province3.Alias, "tx")
}

//func TestCanFindAllProvinces(t *testing.T) {
//	list, err := provinceService.FindAll(&pagination)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, list)
//}
