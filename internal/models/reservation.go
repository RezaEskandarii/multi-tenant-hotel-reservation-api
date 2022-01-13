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
	HotelId         uint64          `json:"hotel_id" valid:"required"`
	Hotel           Hotel           `json:"hotel" valid:"-"  gorm:"foreignKey:HotelId;references:id"`
	SupervisorId    uint64          `json:"supervisor_id" valid:"required"`
	Supervisor      Guest           `json:"supervisor" valid:"-"   gorm:"foreignKey:SupervisorId;references:id"`
	CheckinDate     time.Time       `json:"checkin_date" valid:"required"`
	CheckoutDate    time.Time       `json:"checkout_date" valid:"required"`
	RoomId          uint64          `json:"room_id" valid:"required"`
	Room            Room            `json:"room" valid:"-"   gorm:"foreignKey:RoomId;references:id"`
	RateCodeId      uint64          `json:"rate_code_id" valid:"required"`
	RateCode        RateCode        `json:"rate_code" valid:"required"   gorm:"foreignKey:RateCodeId;references:id"`
	GuestCount      uint64          `json:"guest_count" valid:"required"`
	ParentId        uint64          `json:"parent_id" valid:"-"`
	Parent          *Reservation    `json:"parent" gorm:"foreignKey:ParentId;references:id"`
	ReservationType ReservationType `json:"reservation_type"`
	Price           float64         `json:"price"`
	Nights          float64         `json:"nights"`
	RequestKey      string          `json:"request_key" gorm:"-"`

	///Children        []*Reservation  `gorm:"many2many:reservation_children;association_jointable_foreignkey:parent_id"`
}

func (r *Reservation) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

func (r *Reservation) BeforeCreate(tx *gorm.DB) error {

	//if r.ReservationType == "" {
	//	r.ReservationType = Normal
	//}
	//
	//if _, err := r.Validate(); err != nil {
	//	tx.AddError(err)
	//	return err
	//}
	return nil
}
