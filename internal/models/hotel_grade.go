package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type HotelGrade struct {
	BaseModel
	Name        string     `json:"name" gorm:"type:varchar(100)" valid:"required"`
	HotelType   *HotelType `json:"hotel_type" valid:"-"`
	HotelTypeId uint64     `json:"hotel_type_id" gorm:"foreignKey:HotelType" valid:"required"`
}

func (r *HotelGrade) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

func (r *HotelGrade) BeforeCreate(tx *gorm.DB) error {
	if _, err := r.Validate(); err != nil {
		tx.AddError(err)
		return err
	}
	return nil
}
