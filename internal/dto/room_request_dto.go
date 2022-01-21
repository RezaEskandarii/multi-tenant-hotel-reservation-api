package dto

import "time"

type ReservationRequestType int

const (
	CreateReservation = iota
	UpdateReservation
)

type RoomRequestDto struct {
	RequestType  ReservationRequestType `json:"request_type"`
	CheckInDate  *time.Time             `json:"check_in_date"`
	CheckOutDate *time.Time             `json:"check_out_date"`
	RoomId       uint64                 `json:"room_id"`
}
