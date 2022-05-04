package dto

type RateCodeDto struct {
	BaseDto
	Name        string       `json:"name"  valid:"required"`
	Hotel       *HotelDto    `json:"hotel"  valid:"-"`
	HotelId     uint64       `json:"hotel_id"  valid:"required"`
	Currency    *CurrencyDto `json:"currency"  valid:"-"`
	CurrencyId  uint64       `json:"currency_id" valid:"required"`
	RateGroupId uint64       `json:"rate_group_id"`
	Status      bool         `json:"status"`
}
