package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Province province model
type Province struct {
	BaseModel
	Name      string   `json:"name" valid:"required"`
	Alias     string   `json:"alias" valid:"required"`
	Cities    []*City  `json:"cities" valid:"-"`
	CountryId uint64   `json:"country_id" valid:"required"`
	Country   *Country `json:"country" valid:"-"`
}

func (p *Province) Validate() (bool, error) {

	return govalidator.ValidateStruct(p)
}

func (p *Province) BeforeCreate(tx *gorm.DB) error {
	_, err := p.Validate()

	if err != nil {
		tx.AddError(err)
		return err
	}

	return nil
}
