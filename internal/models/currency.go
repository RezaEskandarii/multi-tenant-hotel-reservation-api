package models

import (
	"github.com/asaskevich/govalidator"
)

type Currency struct {
	BaseModel
	Name   string `json:"name" valid:"required"    gorm:"type:varchar(50)"`
	Symbol string `json:"symbol" valid:"required"  gorm:"type:varchar(50)"`
}

// Validate validates currency struct
func (c *Currency) Validate() (bool, error) {

	return govalidator.ValidateStruct(c)
}

func (c *Currency) SetAudit(username string) {
	c.CreatedBy = username
	c.UpdatedBy = username
}
