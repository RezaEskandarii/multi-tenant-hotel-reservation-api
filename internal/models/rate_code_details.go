package models

import "time"

type RateCodeDetail struct {
	BaseModel
	RateCode   *RateCode  `json:"rate_code"  gorm:"foreignkey:RateCodeId"`
	RateCodeId uint64     `json:"rate_code_id"`
	NightCount uint64     `json:"night_count"`
	DateStart  *time.Time `json:"date_start"`
	DateEnd    *time.Time `json:"date_end"`
	Room       *Room      `json:"room" gorm:"foreignkey:RoomId"`
	RoomId     uint64     `json:"room_id"`

	RatePrices []*RateCodeDetailPrice `json:"rate_prices"`
}

type RateCodeDetailPrice struct {
	BaseModel
	GuestCount       uint64          `json:"guest_count"`
	Price            float64         `json:"price"`
	RateCodeDetail   *RateCodeDetail `json:"rate_code_detail" gorm:"foreignkey:RateCodeDetailId"`
	RateCodeDetailId uint64          `json:"rate_code_detail_id"`
}
