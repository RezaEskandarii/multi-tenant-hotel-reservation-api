package models

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type ReservationCheckStatus int

const (
	CheckIn ReservationCheckStatus = iota
	Checkout
	Block
)

type Reservation struct {
	BaseModel
	HotelId      uint64                 `json:"hotel_id" valid:"-"`
	Hotel        *Hotel                 `json:"hotel" valid:"-"  gorm:"foreignKey:HotelId;references:id"`
	SupervisorId uint64                 `json:"supervisor_id" valid:"required"`
	Supervisor   *Guest                 `json:"supervisor" valid:"-"   gorm:"foreignKey:SupervisorId;references:id"`
	CheckinDate  *time.Time             `json:"checkin_date" valid:"required"`
	CheckoutDate *time.Time             `json:"checkout_date" valid:"required"`
	RoomId       uint64                 `json:"room_id" valid:"required"`
	Room         *Room                  `json:"room" valid:"-"   gorm:"foreignKey:RoomId;references:id"`
	RateCodeId   uint64                 `json:"rate_code_id" valid:"required"`
	RateCode     *RateCode              `json:"rate_code" valid:"-"   gorm:"foreignKey:RateCodeId;references:id"`
	GuestCount   uint64                 `json:"guest_count"`
	ParentId     uint64                 `json:"parent_id" valid:"-"`
	Parent       *Reservation           `json:"parent" gorm:"foreignKey:ParentId;references:id"`
	Price        float64                `json:"price"`
	Nights       float64                `json:"nights"`
	RequestKey   string                 `json:"request_key" gorm:"-" valid:"required"`
	CheckStatus  ReservationCheckStatus `json:"check_status" valid:"required"`
	Sharers      []*Sharer              `json:"sharers"`
}

func (r *Reservation) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

func (r *Reservation) SetAudit(username string) {
	r.CreatedBy = username
	r.UpdatedBy = username
}

func (r *Reservation) SetUpdatedBy(username string) {
	r.UpdatedBy = username
}
