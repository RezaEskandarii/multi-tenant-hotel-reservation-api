package dto

import "github.com/asaskevich/govalidator"

type HotelGradeDto struct {
	BaseDto
	Name        string        `json:"name" valid:"required"`
	HotelType   *HotelTypeDto `json:"hotel_type" valid:"-"`
	HotelTypeId uint64        `json:"hotel_type_id"  valid:"required"`
}

func (r *HotelGradeDto) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

type HotelTypeDto struct {
	BaseDto
	Name   string           `json:"name"  valid:"required"`
	Grades []*HotelGradeDto `json:"grades"  valid:"-"`
}

func (h *HotelTypeDto) Validate() (bool, error) {

	return govalidator.ValidateStruct(h)
}
