package models

import (
	"time"
)

type ReservationRequest struct {
	BaseModel
	Room       Room      `json:"room"`
	RoomId     uint64    `json:"room_id"`
	ExpireTime time.Time `json:"expire_time"`
	RequestKey string    `json:"request_key"` // lock room and prevent to concurrent reservation of same room.
}
