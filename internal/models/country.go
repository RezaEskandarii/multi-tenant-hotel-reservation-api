package models

import (
	"github.com/asaskevich/govalidator"
)

// Country country struct
type Country struct {
	BaseModel
	Name  string `json:"name" valid:"required"  gorm:"type:varchar(50)"`
	Alias string `json:"alias" valid:"required"  gorm:"type:varchar(50)"`
	// swagger:ignore
	Provinces []*Province `json:"provinces" gorm:"foreignKey:CountryId" valid:"-" swagger:ignore`
}

func (c *Country) Validate() (bool, error) {

	return govalidator.ValidateStruct(c)
}

func (c *Country) SetAudit(username string) {
	c.CreatedBy = username
	c.UpdatedBy = username
}

func (c *Country) SetUpdatedBy(username string) {
	c.UpdatedBy = username
}
