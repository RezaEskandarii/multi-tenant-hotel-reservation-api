package models

// Country country struct
type Country struct {
	BaseModel
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Cities []City `json:"cities"`
}
