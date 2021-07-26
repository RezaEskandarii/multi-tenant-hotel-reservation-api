package models

// City city struct
type City struct {
	BaseModel
	Name       string    `json:"name" valid:"required"`
	Alias      string    `json:"alias" valid:"required"`
	ProvinceId uint64    `json:"province_id" valid:"required"`
	Province   *Province `json:"province,omitempty" gorm:"foreignkey:ProvinceId"`
}
