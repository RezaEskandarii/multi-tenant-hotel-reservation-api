package models

// Province province model
type Province struct {
	BaseModel
	Name      string   `json:"name"`
	Alias     string   `json:"alias"`
	Cities    []*City  `json:"cities"`
	CountryId uint     `json:"country_id"`
	Country   *Country `json:"country"`
}
