package dto

import "time"

type GetRatePriceDto struct {
	RoomId     uint64     `json:"room_id"`
	NightCount uint64     `json:"night_count"`
	GuestCount uint64     `json:"guest_count"`
	DateStart  *time.Time `json:"date_start"`
	DateEnd    *time.Time `json:"date_end"`
}

type RateCodePricesDto struct {
	RateCodeName string     `json:"rate_code_name"`
	RateCodeId   uint64     `json:"rate_code_id"`
	CreatedAt    *time.Time `json:"created_at"`
	RoomId       uint64     `json:"room_id"`
	DetailId     uint64     `json:"detail_id"`
	DateStart    *time.Time `json:"date_start"`
	DateEnd      *time.Time `json:"date_end"`
	Price        float64    `json:"price"`
	GuestCount   uint64     `json:"guest_count"`
	RatePriceId  uint64     `json:"rate_price_id"`
}
