package models

// Province province model
type Province struct {
	BaseModel
	Name      string   `json:"name" valid:"required"`
	Alias     string   `json:"alias" valid:"required"`
	Cities    []*City  `json:"cities"`
	CountryId uint     `json:"country_id" valid:"required"`
	Country   *Country `json:"country"`
}
