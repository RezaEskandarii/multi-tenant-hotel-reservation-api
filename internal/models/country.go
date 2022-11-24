package models

import (
	"github.com/asaskevich/govalidator"
)

// Country struct contains all information about city.
type Country struct {
	BaseModel
	Name      string      `json:"name"  gorm:"type:varchar(50)"`
	Alias     string      `json:"alias" gorm:"type:varchar(50)"`
	Provinces []*Province `json:"provinces" gorm:"foreignKey:CountryId"`
}

// CountryCreateUpdate contains all information about update or create new country.
type CountryCreateUpdate struct {
	BaseModel
	Name  string `json:"name" valid:"required"`
	Alias string `json:"alias" valid:"required"  `
}

// GetCountry
type GetCountry struct {
	BaseDto
	Name      string      `json:"name"`
	Alias     string      `json:"alias"`
	Provinces []*Province `json:"provinces" swagger:ignore`
}

func (c *CountryCreateUpdate) Validate() (bool, error) {

	return govalidator.ValidateStruct(c)
}

func (c *Country) SetAudit(username string) {
	c.CreatedBy = username
	c.UpdatedBy = username
}

func (c *Country) SetUpdatedBy(username string) {
	c.UpdatedBy = username
}
