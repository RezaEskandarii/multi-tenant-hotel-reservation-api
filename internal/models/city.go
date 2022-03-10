package models

import (
	"github.com/asaskevich/govalidator"
)

// City city struct
type City struct {
	BaseModel
	Name       string    `json:"name" valid:"required"  gorm:"type:varchar(255)"`
	Alias      string    `json:"alias" valid:"required"  gorm:"type:varchar(255)"`
	ProvinceId uint64    `json:"province_id" valid:"required"`
	Province   *Province `json:"province,omitempty" gorm:"foreignkey:ProvinceId" valid:"-"`
}

func (c *City) Validate() (bool, error) {

	return govalidator.ValidateStruct(c)
}

func (c *City) SetAudit(username string) {
	c.CreatedBy = username
	c.UpdatedBy = username
}
