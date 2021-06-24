package models

type Currency struct {
	BaseModel
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
