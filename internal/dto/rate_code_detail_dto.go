package dto

import "time"

type RateCodeDetailDto struct {
	BaseDto
	RateCode   *RateCodeDto `json:"rate_code"     `
	RateCodeId uint64       `json:"rate_code_id"  valid:"required"`
	MinNights  uint64       `json:"min_nights"    valid:"required"`
	MaxNights  uint64       `json:"max_nights"    valid:"required"`
	DateStart  *time.Time   `json:"date_start"    valid:"required"`
	DateEnd    *time.Time   `json:"date_end"      valid:"required"`
	Room       *Room        `json:"room"`
	RoomId     uint64       `json:"room_id"       valid:"required"`

	RatePrices []*RateCodeDetailPriceDto `json:"rate_prices" valid:"required"`
}

type RateCodeDetailPriceDto struct {
	BaseDto
	GuestCount       uint64             `json:"guest_count"  valid:"required"`
	Price            float64            `json:"price"  valid:"required"`
	RateCodeDetail   *RateCodeDetailDto `json:"rate_code_detail"`
	RateCodeDetailId uint64             `json:"rate_code_detail_id"`
}
