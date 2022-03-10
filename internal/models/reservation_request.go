package models

import (
	"time"
)

type ReservationRequest struct {
	BaseModel
	Room         *Room      `json:"room"`
	RoomId       uint64     `json:"room_id"`
	ExpireTime   time.Time  `json:"expire_time"`
	RequestKey   string     `json:"request_key"` // lock room and prevent to concurrent reservation of same room.
	CheckInDate  *time.Time `json:"check_in_date"`
	CheckOutDate *time.Time `json:"check_out_date"`
}

func (r *ReservationRequest) SetAudit(username string) {
	r.CreatedBy = username
	r.UpdatedBy = username
}
