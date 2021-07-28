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
	Cities    []*City  `json:"cities"`
	CountryId uint     `json:"country_id" valid:"required"`
	Country   *Country `json:"country"`
}

func (p *Province) Validate() (bool, error) {

	ok, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (p *Province) BeforeCreate(tx *gorm.DB) error {

	_, err := p.Validate()

	if err != nil {
		tx.AddError(err)
		return err
	}

	return nil
}
