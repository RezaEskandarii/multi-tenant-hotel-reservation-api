package models

type Province struct {
	BaseModel
	Name   string  `json:"name"`
	Alias  string  `json:"alias"`
	Cities []*City `json:"cities"`
}
