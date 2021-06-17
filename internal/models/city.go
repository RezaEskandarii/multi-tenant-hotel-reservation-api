package models

// City city struct
type City struct {
	BaseModel
	Name       string    `json:"name"`
	Alias      string    `json:"alias"`
	ProvinceId uint64    `json:"province_id"`
	Province   *Province `json:"province,omitempty" gorm:"foreignkey:ProvinceId"`
}
