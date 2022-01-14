package dto

import "time"

type RoomRequestDto struct {
	CheckInDate  *time.Time `json:"check_in_date"`
	CheckOutDate *time.Time `json:"check_out_date"`
	RoomId       uint64     `json:"room_id"`
}
