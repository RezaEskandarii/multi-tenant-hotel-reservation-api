package models

import (
	"github.com/asaskevich/govalidator"
)

// Province province model
type Province struct {
	BaseModel
	Name      string   `json:"name" valid:"required"  gorm:"type:varchar(255)"`
	Alias     string   `json:"alias" valid:"required"  gorm:"type:varchar(255)"`
	Cities    []*City  `json:"cities" valid:"-"`
	CountryId uint64   `json:"country_id" valid:"required"`
	Country   *Country `json:"country" valid:"-"`
}

func (p *Province) Validate() (bool, error) {

	return govalidator.ValidateStruct(p)
}
