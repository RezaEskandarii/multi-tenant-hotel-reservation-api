package models

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Guest struct {
	BaseModel
	Country            *Country   `json:"country" valid:"-"`
	CountryId          uint64     `json:"country_id" valid:"required"`
	Gender             Gender     `json:"gender" valid:"required"`
	FirstName          string     `json:"first_name" valid:"required"  gorm:"type:varchar(255)"`
	MiddleName         string     `json:"middle_name"  gorm:"type:varchar(255)"`
	LastName           string     `json:"last_name" valid:"required"  gorm:"type:varchar(255)"`
	NationalId         string     `json:"national_id" valid:"required"  gorm:"type:varchar(255)"`
	CellNumber         string     `json:"cell_number" valid:"required"  gorm:"type:varchar(20)"`
	PhoneNumber        string     `json:"phone_number"  gorm:"type:varchar(20)"`
	PassportNumber     string     `json:"passport_number"  gorm:"type:varchar(50)"`
	PassportIssueDate  string     `json:"passport_date_of_issue"`
	PassportExpireDate string     `json:"passport_expire_date"`
	Email              string     `json:"email" valid:"email"  gorm:"type:varchar(255)"`
	DateOfBirth        *time.Time `json:"date_of_birth"`
	Address            string     `json:"address" valid:"required"`
}

// Validate validates guest model.
func (g *Guest) Validate() (bool, error) {

	return govalidator.ValidateStruct(g)
}

func (g *Guest) SetAudit(username string) {
	g.CreatedBy = username
	g.UpdatedBy = username
}

func (g *Guest) SetUpdatedBy(username string) {
	g.UpdatedBy = username
}
