package models

import (
	"fmt"
	"gorm.io/gorm"
	"reservation-api/internal/utils"
	"time"
)

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

func (r *ReservationLock) BeforeCreate(tx *gorm.DB) error {
	strToHash := fmt.Sprintf("%d_%d_%d_%s", r.RoomId, r.GuestId, r.RateCodeId, r.StartTime)
	r.LockKey = utils.GenerateSHA256(strToHash)
	return nil
}
