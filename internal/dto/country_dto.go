package dto

import "github.com/asaskevich/govalidator"

type BaseDto struct {
}

type CountryDto struct {
	BaseDto
	Name      string         `json:"name" valid:"required"`
	Alias     string         `json:"alias" valid:"required"`
	Provinces []*ProvinceDto `json:"provinces" valid:"-"`
}

func (c *CountryDto) Validate() (bool, error) {

	return govalidator.ValidateStruct(c)
}

type ProvinceDto struct {
	BaseDto
	Name      string      `json:"name" valid:"required" `
	Alias     string      `json:"alias" valid:"required" `
	Cities    []*CityDto  `json:"cities" valid:"-"`
	CountryId uint64      `json:"country_id" valid:"required"`
	Country   *CountryDto `json:"country" valid:"-"`
}

func (c *ProvinceDto) Validate() (bool, error) {

	return govalidator.ValidateStruct(c)
}

type CityDto struct {
	BaseDto
	Name       string       `json:"name" valid:"required"`
	Alias      string       `json:"alias" valid:"required"  `
	ProvinceId uint64       `json:"province_id" valid:"required"`
	Province   *ProvinceDto `json:"province,omitempty"  valid:"-"`
}
