package models

import "time"

type RateCodeDetail struct {
	BaseModel
	RateCode   RateCode   `json:"rate_code"`
	RateCodeId uint64     `json:"rate_code_id"`
	NightCount uint64     `json:"night_count"`
	DateStart  *time.Time `json:"date_start"`
	DateEnd    *time.Time `json:"date_end"`
	Room       *Room      `json:"room"`
	RoomId     uint64     `json:"room_id"`

	RatePrices []*RateCodeDetailPrice `json:"rate_prices"`
}

type RateCodeDetailPrice struct {
	GuestCount       uint64         `json:"guest_count"`
	Price            uint64         `json:"price"`
	RateCodeDetail   RateCodeDetail `json:"rate_code_detail"`
	RateCodeDetailId uint64         `json:"rate_code_detail_id"`
}
