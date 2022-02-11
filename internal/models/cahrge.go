package models

type Charge struct {
	BaseModel
	Amount        float64     `json:"amount"`
	ReceiptEmail  string      `json:"receipt_email"`
	Reservation   Reservation `json:"reservation"  gorm:"foreignkey:ReservationId"`
	ReservationId uint64      `json:"reservation_id"`
	CreatedBy     string      `json:"created_by"`
}
