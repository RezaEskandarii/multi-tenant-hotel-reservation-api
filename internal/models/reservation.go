package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type ReservationType string

const (
	Normal ReservationType = "Normal" // normal reservation.
	Lock   ReservationType = "Lock"   // block a period of times to prevent reservation.
)

type Reservation struct {
	BaseModel
	Residence       Residence       `json:"residence" valid:"-"`
	ResidenceId     uint64          `json:"residence_id" valid:"required"`
	Supervisor      Guest           `json:"supervisor" valid:"-"`
	SupervisorId    uint64          `json:"supervisor_id" valid:"required"`
	CheckinDate     time.Time       `json:"checkin_date" valid:"required"`
	CheckoutDate    time.Time       `json:"checkout_date" valid:"required"`
	Room            Room            `json:"room" valid:"-"`
	RoomId          uint64          `json:"room_id" valid:"required"`
	RateCode        RateCode        `json:"rate_code" valid:"required"`
	RateCodeId      uint64          `json:"rate_code_id" valid:"required"`
	GuestCount      uint64          `json:"guest_count" valid:"required"`
	Children        []*Reservation  `json:"children" valid:"-"`
	ParentId        uint64          `json:"parent_id" valid:"-"`
	ReservationType ReservationType `json:"reservation_type"`
}

func (r *Reservation) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

func (r *Reservation) BeforeCreate(tx *gorm.DB) error {

	if r.ReservationType == "" {
		r.ReservationType = Normal
	}

	if _, err := r.Validate(); err != nil {
		tx.AddError(err)
		return err
	}
	return nil
}
