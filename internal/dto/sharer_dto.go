package dto

type SharerDto struct {
	BaseDto
	GuestId       uint64               `json:"guest_id"`
	Guest         GuestDto             `json:"guest"`
	ReservationId uint64               `json:"reservation_id"`
	Reservation   ReservationCreateDto `json:"reservation"`
}
