package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type RoomType struct {
	BaseModel
	Residence     Residence `json:"residence" valid:"-"`
	ResidenceId   uint64    `json:"residence_id" gorm:"foreiknKey:Residence" valid:"required"`
	Name          string    `json:"name" valid:"required"`
	MaxGuestCount uint64    `json:"max_guest_count" valid:"required"`
	Description   string    `json:"description" valid:"maxstringlength(150)"`
}

func (r *RoomType) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

func (r *RoomType) BeforeCreate(tx *gorm.DB) error {
	_, err := r.Validate()

	if err != nil {
		tx.AddError(err)
		return err
	}
	return nil
}
