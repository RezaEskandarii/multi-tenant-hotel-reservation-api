package models

// Country country struct
type Country struct {
	BaseModel
	Name      string      `json:"name" valid:"required"`
	Alias     string      `json:"alias" valid:"required"`
	Provinces []*Province `json:"provinces" gorm:"foreignKey:CountryId"`
}
