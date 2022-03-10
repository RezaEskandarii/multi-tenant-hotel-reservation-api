package models

import (
	"github.com/asaskevich/govalidator"
)

type RoomType struct {
	BaseModel
	Hotel         *Hotel `json:"hotel" valid:"-"`
	HotelId       uint64 `json:"hotel_id" gorm:"foreiknKey:Hotel" valid:"required"`
	Name          string `json:"name" valid:"required"  gorm:"type:varchar(255)"`
	MaxGuestCount uint64 `json:"max_guest_count" valid:"required"`
	Description   string `json:"description" valid:"maxstringlength(255)"  gorm:"type:varchar(255)"`
}

func (r *RoomType) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

func (r *RoomType) SetAudit(username string) {
	r.CreatedBy = username
	r.UpdatedBy = username
}

func (r *RoomType) SetUpdatedBy(username string) {
	r.UpdatedBy = username
}
