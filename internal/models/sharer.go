package models

type Sharer struct {
	Guest         Guest  `json:"guest"`
	GuestId       uint64 `json:"guest_id"`
	ReservationId uint64 `json:"reservation_id"`
}
