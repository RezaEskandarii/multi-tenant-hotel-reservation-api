package models

import "time"

type Reservation struct {
	BaseModel
	Residence    Residence      `json:"residence" valid:"-"`
	ResidenceId  uint64         `json:"residence_id" valid:"required"`
	Supervisor   Guest          `json:"supervisor" valid:"-"`
	SupervisorId uint64         `json:"supervisor_id" valid:"required"`
	CheckinDate  time.Time      `json:"checkin_date" valid:"required"`
	CheckoutDate time.Time      `json:"checkout_date" valid:"required"`
	Room         Room           `json:"room" valid:"-"`
	RoomId       uint64         `json:"room_id" valid:"required"`
	RateCode     RateCode       `json:"rate_code" valid:"required"`
	RateCodeId   uint64         `json:"rate_code_id" valid:"required"`
	GuestCount   uint64         `json:"guest_count" valid:"required"`
	Children     []*Reservation `json:"children" valid:"-"`
	ParentId     uint64         `json:"parent_id" valid:"-"`
}
