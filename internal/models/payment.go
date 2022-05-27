package models

import "time"

type PaymentType int

const (
	DEBIT PaymentType = iota
	CREDIT
)

type Payment struct {
	BaseModel
	Amount        float64     `json:"amount"`
	PaymentType   PaymentType `json:"payment_type"`
	Payer         Guest       `json:"payer"`
	PayerId       uint64      `json:"payer_id"`
	PaymentDate   *time.Time  `json:"payment_date"`
	Reservation   Reservation `json:"reservation"`
	ReservationId uint64      `json:"reservation_id"`
}
