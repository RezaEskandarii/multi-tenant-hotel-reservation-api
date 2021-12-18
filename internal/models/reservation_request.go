package models

import (
	"time"
)

type ReservationRequest struct {
	BaseModel
	Room                    Room      `json:"room"`
	RoomId                  uint64    `json:"room_id"`
	GuestId                 uint64    `json:"guest_id"`
	RateCodeId              uint64    `json:"rate_code_id"`
	ExpireTime              time.Time `json:"expire_time"`
	LockKey                 string    `json:"lock_key"`
	ReservationCheckinDate  time.Time `json:"reservation_checkin_date"`
	ReservationCheckoutDate time.Time `json:"reservation_checkout_date"`
}
