package dto

type CurrencyDto struct {
	BaseDto
	Name   string `json:"name" valid:"required"`
	Symbol string `json:"symbol" valid:"required"`
}
