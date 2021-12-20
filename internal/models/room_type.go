package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type RoomType struct {
	BaseModel
	Hotel         *Hotel `json:"hotel" valid:"-"`
	HotelId       uint64 `json:"hotel_id" gorm:"foreiknKey:Hotel" valid:"required"`
	Name          string `json:"name" valid:"required"`
	MaxGuestCount uint64 `json:"max_guest_count" valid:"required"`
	Description   string `json:"description" valid:"maxstringlength(150)"`
}

func (r *RoomType) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

func (r *RoomType) BeforeCreate(tx *gorm.DB) error {
	if _, err := r.Validate(); err != nil {
		tx.AddError(err)
		return err
	}
	return nil
}
