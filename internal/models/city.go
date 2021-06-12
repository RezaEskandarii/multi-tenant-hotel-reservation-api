package models

// City city struct
type City struct {
	BaseModel
	Name      string  `json:"name"`
	Alias     string  `json:"alias"`
	Country   Country `json:"country" gorm:"foreignkey:CountryId"`
	CountryId uint64  `json:"country_id"`
}
