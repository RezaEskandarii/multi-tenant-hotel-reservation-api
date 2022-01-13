package models

import "time"

type RateCodeDetail struct {
	BaseModel
	RateCode   *RateCode  `json:"rate_code"     gorm:"foreignkey:RateCodeId"`
	RateCodeId uint64     `json:"rate_code_id"  valid:"required"`
	MinNights  uint64     `json:"min_nights"    valid:"required"`
	MaxNights  uint64     `json:"max_nights"    valid:"required"`
	DateStart  *time.Time `json:"date_start"    valid:"required"`
	DateEnd    *time.Time `json:"date_end"      valid:"required"`
	Room       *Room      `json:"room"          gorm:"foreignkey:RoomId"`
	RoomId     uint64     `json:"room_id"       valid:"required"`

	RatePrices []*RateCodeDetailPrice `json:"rate_prices" valid:"required"`
}

type RateCodeDetailPrice struct {
	BaseModel
	GuestCount       uint64          `json:"guest_count"  valid:"required"`
	Price            float64         `json:"price"  valid:"required"`
	RateCodeDetail   *RateCodeDetail `json:"rate_code_detail" gorm:"foreignkey:RateCodeDetailId"`
	RateCodeDetailId uint64          `json:"rate_code_detail_id"`
}
