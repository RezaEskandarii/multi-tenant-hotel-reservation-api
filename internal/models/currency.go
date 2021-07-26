package models

type Currency struct {
	BaseModel
	Name   string `json:"name" valid:"required"`
	Symbol string `json:"symbol" valid:"required"`
}
