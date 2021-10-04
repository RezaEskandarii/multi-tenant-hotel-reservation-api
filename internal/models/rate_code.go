package models

type RateCode struct {
	BaseModel
	Name        string    `json:"name"`
	Residence   Residence `json:"residence"`
	ResidenceId uint64    `json:"residence_id"`
	Currency    Currency  `json:"currency"`
	CurrencyId  uint64    `json:"currency_id"`
	RateGroup   RateGroup `json:"rate_group"`
	RateGroupId uint64    `json:"rate_group_id"`
	Guest       Guest     `json:"guest"`
	GuestIs     uint64    `json:"guest_is"`
	Status      bool      `json:"status"`
}
