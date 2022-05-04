package dto

import "time"

type ReservationCheckStatus int

const (
	CheckIn ReservationCheckStatus = iota
	Checkout
	Block
)

type ReservationCreateDto struct {
	BaseDto
	HotelId      uint64                 `json:"hotel_id" valid:"-"`
	Hotel        *HotelDto              `json:"hotel" valid:"-"`
	SupervisorId uint64                 `json:"supervisor_id" valid:"required"`
	Supervisor   *GuestDto              `json:"supervisor" valid:"-"`
	CheckinDate  *time.Time             `json:"checkin_date" valid:"required"`
	CheckoutDate *time.Time             `json:"checkout_date" valid:"required"`
	RoomId       uint64                 `json:"room_id" valid:"required"`
	Room         *Room                  `json:"room" valid:"-"`
	RateCodeId   uint64                 `json:"rate_code_id" valid:"required"`
	RateCode     *RateCodeDto           `json:"rate_code" valid:"-"`
	GuestCount   uint64                 `json:"guest_count"`
	ParentId     uint64                 `json:"parent_id" valid:"-"`
	Parent       *ReservationCreateDto  `json:"parent"`
	Price        float64                `json:"price"`
	Nights       float64                `json:"nights"`
	RequestKey   string                 `json:"request_key"`
	CheckStatus  ReservationCheckStatus `json:"check_status" valid:"required"`
	Sharers      []*SharerDto           `json:"sharers"`
}
