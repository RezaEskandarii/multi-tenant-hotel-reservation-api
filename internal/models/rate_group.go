package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// RateGroup is struct for rating group of reservations rate price.
type RateGroup struct {
	BaseModel
	Name        string `json:"name" valid:"required"`
	Description string `json:"description"`
	HotelId     uint64 `json:"hotel_id" valid:"required"`
	Hotel       Hotel  `json:"hotel"`
}

func (r *RateGroup) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

func (r *RateGroup) BeforeCreate(tx *gorm.DB) error {
	if _, err := r.Validate(); err != nil {
		tx.AddError(err)
		return err
	}
	return nil
}
