package models

import "time"

type Payment struct {
	BaseModel
	Amount      float64    `json:"amount"`
	Payer       Guest      `json:"payer"`
	PayerId     uint64     `json:"payer_id"`
	PaymentDate *time.Time `json:"payment_date"`
}
