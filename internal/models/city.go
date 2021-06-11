package models

type City struct {
	baseModel
	Name      string  `json:"name"`
	Alias     string  `json:"alias"`
	Country   Country `json:"country"`
	CountryId uint64  `json:"country_id"`
}
