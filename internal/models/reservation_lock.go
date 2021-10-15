package models

import "time"

type ReservationLock struct {
	BaseModel
	Room       Room      `json:"room"`
	RoomId     uint64    `json:"room_id"`
	GuestId    uint64    `json:"guest_id"`
	RateCodeId uint64    `json:"rate_code_id"`
	StartTime  time.Time `json:"start_time"`
	ExpireTime time.Time `json:"expire_time"`
	LockKey    string    `json:"lock_key"`
}
