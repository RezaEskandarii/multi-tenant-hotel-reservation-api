package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
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

func (p *Province) BeforeCreate(tx *gorm.DB) error {
	if _, err := p.Validate(); err != nil {
		tx.AddError(err)
		return err
	}
	return nil
}
