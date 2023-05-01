package models

import (
	"github.com/shopspring/decimal"
)

type Wallet struct {
	BaseModel
	UserId  uint64          `gorm:"not null" json:"user_id"`
	Balance decimal.Decimal `gorm:" type:decimal(10,2);not null;default:0" json:"balance"`
}
