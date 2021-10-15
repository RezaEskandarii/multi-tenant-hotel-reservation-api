package models

import "time"

type Reservation struct {
	BaseModel
	Residence    Residence      `json:"residence"`
	ResidenceId  uint64         `json:"residence_id"`
	Supervisor   Guest          `json:"supervisor"`
	SupervisorId uint64         `json:"supervisor_id"`
	CheckinDate  time.Time      `json:"checkin_date"`
	CheckoutDate time.Time      `json:"checkout_date"`
	Room         Room           `json:"room"`
	RoomId       uint64         `json:"room_id"`
	RateCode     RateCode       `json:"rate_code"`
	RateCodeId   uint64         `json:"rate_code_id"`
	GuestCount   uint64         `json:"guest_count"`
	Children     []*Reservation `json:"children"`
	ParentId     uint64         `json:"parent_id"`
}
