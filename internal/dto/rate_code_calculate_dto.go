package dto

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type GetRatePriceDto struct {
	RoomId     uint64     `json:"room_id" valid:"required"`
	NightCount float64    `json:"night_count" valid:"required"`
	GuestCount uint64     `json:"guest_count" valid:"required"`
	DateStart  *time.Time `json:"date_start"  valid:"required"`
	DateEnd    *time.Time `json:"date_end"    valid:"required"`
	RateCodeId uint64     `json:"rate_code_id"`
}

func (d *GetRatePriceDto) Validate() (bool, error) {
	ok, err := govalidator.ValidateStruct(d)
	if err != nil {
		return false, err
	}
	return ok, nil
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
