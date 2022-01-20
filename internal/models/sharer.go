package models

type Sharer struct {
	BaseModel
	GuestId       uint64      `json:"guest_id"`
	Guest         Guest       `json:"guest" gorm:"foreignKey:GuestId;references:id"`
	ReservationId uint64      `json:"reservation_id"`
	Reservation   Reservation `json:"-" gorm:"foreignKey:ReservationId;references:id"`
}
