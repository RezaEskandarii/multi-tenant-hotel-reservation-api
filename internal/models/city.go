package models

// City city struct
type City struct {
	BaseModel
	Name      string   `json:"name"`
	Alias     string   `json:"alias"`
	CountryId uint64   `json:"country_id"`
	Country   *Country `json:"country,omitempty" gorm:"foreignkey:CountryId"`
}
