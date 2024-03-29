package models

// Province province model
type Province struct {
	BaseModel
	Name      string   `json:"name" valid:"required"  gorm:"type:varchar(255)"`
	Alias     string   `json:"alias" valid:"required"  gorm:"type:varchar(255)"`
	Cities    []*City  `json:"cities" valid:"-"`
	CountryId uint64   `json:"country_id" valid:"required"`
	Country   *Country `json:"country" valid:"-"`
}

type ProvinceCreateOrUpdate struct {
	BaseModel
	Name      string   `json:"name" valid:"required"`
	Alias     string   `json:"alias" valid:"required"`
	Cities    []*City  `json:"cities" valid:"-"`
	CountryId uint64   `json:"country_id" valid:"required"`
	Country   *Country `json:"country" valid:"-"`
}

type GetCity struct {
}

func (p *Province) SetAudit(username string) {
	p.CreatedBy = username
	p.UpdatedBy = username
}

func (p *Province) SetUpdatedBy(username string) {
	p.UpdatedBy = username
}
