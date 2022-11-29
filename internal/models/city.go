package models

// City struct contains all information about city.
type City struct {
	BaseModel
	Name       string    `json:"name"  gorm:"type:varchar(255)"`
	Alias      string    `json:"alias" gorm:"type:varchar(255)"`
	ProvinceId uint64    `json:"province_id"`
	Province   *Province `json:"province,omitempty" gorm:"foreignkey:ProvinceId"`
}

// CreateUpdateCity contains all information about update or create new city.
type CreateUpdateCity struct {
	Name       string `json:"name" valid:"required"`
	Alias      string `json:"alias" valid:"required"`
	ProvinceId uint64 `json:"province_id"`
}

func (c *City) SetAudit(username string) {
	c.CreatedBy = username
	c.UpdatedBy = username
}

func (c *City) SetUpdatedBy(username string) {
	c.UpdatedBy = username
}
