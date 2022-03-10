package models

import (
	"github.com/asaskevich/govalidator"
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

func (r *HotelGrade) SetAudit(username string) {
	r.CreatedBy = username
	r.UpdatedBy = username
}
