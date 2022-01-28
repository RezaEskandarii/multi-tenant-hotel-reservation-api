package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Country country struct
type Country struct {
	BaseModel
	Name      string      `json:"name" valid:"required"  gorm:"type:varchar(50)"`
	Alias     string      `json:"alias" valid:"required"  gorm:"type:varchar(50)"`
	Provinces []*Province `json:"provinces" gorm:"foreignKey:CountryId" valid:"-"`
}

func (c *Country) Validate() (bool, error) {

	return govalidator.ValidateStruct(c)
}

func (c *Country) BeforeCreate(tx *gorm.DB) error {
	if _, err := c.Validate(); err != nil {
		tx.AddError(err)
		return err
	}
	return nil
}
