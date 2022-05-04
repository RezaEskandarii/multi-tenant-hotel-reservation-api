package dto

type CountryDto struct {
	BaseDto
	Name      string         `json:"name" valid:"required"`
	Alias     string         `json:"alias" valid:"required"`
	Provinces []*ProvinceDto `json:"provinces" valid:"-"`
}

type ProvinceDto struct {
	BaseDto
	Name      string      `json:"name" valid:"required" `
	Alias     string      `json:"alias" valid:"required" `
	Cities    []*CityDto  `json:"cities" valid:"-"`
	CountryId uint64      `json:"country_id" valid:"required"`
	Country   *CountryDto `json:"country" valid:"-"`
}

type CityDto struct {
	BaseDto
	Name       string       `json:"name" valid:"required"`
	Alias      string       `json:"alias" valid:"required"  `
	ProvinceId uint64       `json:"province_id" valid:"required"`
	Province   *ProvinceDto `json:"province,omitempty"  valid:"-"`
}
