package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
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

func (c *Currency) BeforeCreate(tx *gorm.DB) error {
	if _, err := c.Validate(); err != nil {
		tx.AddError(err)
		return err
	}
	return nil
}
