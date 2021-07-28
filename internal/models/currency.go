package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Currency struct {
	BaseModel
	Name   string `json:"name" valid:"required"`
	Symbol string `json:"symbol" valid:"required"`
}

func (c *Currency) Validate() (bool, error) {
	ok, err := govalidator.ValidateStruct(c)

	if err != nil {
		return false, err
	}

	return ok, nil
}

func (c *Currency) BeforeCreate(tx *gorm.DB) error {
	_, err := c.Validate()

	if err != nil {
		tx.AddError(err)
		return err
	}

	return nil
}
