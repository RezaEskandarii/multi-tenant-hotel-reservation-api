package models

import (
	"github.com/asaskevich/govalidator"
)

type HotelType struct {
	BaseModel
	Name   string        `json:"name" gorm:"type:varchar(100)" valid:"required"`
	Grades []*HotelGrade `json:"grades" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" valid:"-"`
}

func (h *HotelType) Validate() (bool, error) {

	return govalidator.ValidateStruct(h)
}

func (h *HotelType) SetAudit(username string) {
	h.CreatedBy = username
	h.UpdatedBy = username
}

func (h *HotelType) SetUpdatedBy(username string) {
	h.UpdatedBy = username
}
