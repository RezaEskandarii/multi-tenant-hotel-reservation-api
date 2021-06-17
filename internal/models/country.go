package models

// Country country struct
type Country struct {
	BaseModel
	Name      string      `json:"name"`
	Alias     string      `json:"alias"`
	Provinces []*Province `json:"provinces" gorm:"foreignKey:CountryId"`
}
