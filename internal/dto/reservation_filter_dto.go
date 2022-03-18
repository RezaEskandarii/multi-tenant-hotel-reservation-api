package dto

import (
	"reservation-api/internal/models"
	"time"
)

type ReservationFilter struct {
	PaginationFilter
	CheckInFrom *time.Time                     `json:"check_in_from"`
	CheckInTo   *time.Time                     `json:"check_in_to"`
	CreatedFrom *time.Time                     `json:"created_from"`
	CreatedTo   *time.Time                     `json:"created_to"`
	GuestName   string                         `json:"guest_name"`
	RoomId      uint64                         `json:"room_id"`
	RoomTypeId  uint64                         `json:"room_type_id"`
	RateCodeId  uint64                         `json:"rate_code_id"`
	CheckStatus *models.ReservationCheckStatus `json:"check_status"`
}

type ReservationReport struct {
	Data    []*models.Reservation `json:"data"`
	Filters *ReservationFilter    `json:"filters"`
}
