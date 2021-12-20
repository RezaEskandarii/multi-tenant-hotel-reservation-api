package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type RateCode struct {
	BaseModel
	Name        string    `json:"name"  valid:"required"`
	Hotel       Hotel     `json:"hotel"  valid:"-"`
	HotelId     uint64    `json:"hotel_id"  valid:"required"`
	Currency    Currency  `json:"currency"  valid:"-"`
	CurrencyId  uint64    `json:"currency_id" valid:"required"`
	RateGroup   RateGroup `json:"rate_group" valid:"-"`
	RateGroupId uint64    `json:"rate_group_id"  valid:"required"`
	Guest       Guest     `json:"guest"  valid:"-"`
	GuestId     uint64    `json:"guest_id"  valid:"required"`
	Status      bool      `json:"status"`
}

func (r *RateCode) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

func (r *RateCode) BeforeCreate(tx *gorm.DB) error {
	if _, err := r.Validate(); err != nil {
		tx.AddError(err)
		return err
	}
	return nil
}
