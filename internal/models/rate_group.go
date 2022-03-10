package models

import (
	"github.com/asaskevich/govalidator"
)

// RateGroup is struct for rating group of reservations rate price.
type RateGroup struct {
	BaseModel
	Name        string `json:"name" valid:"required"  gorm:"type:varchar(255)"`
	Description string `json:"description"`
	HotelId     uint64 `json:"hotel_id" valid:"required"`
	Hotel       Hotel  `json:"hotel"`
}

func (r *RateGroup) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

func (r *RateGroup) SetAudit(username string) {
	r.CreatedBy = username
	r.UpdatedBy = username
}

func (r *RateGroup) SetUpdatedBy(username string) {
	r.UpdatedBy = username
}
