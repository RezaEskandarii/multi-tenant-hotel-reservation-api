package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type Guest struct {
	BaseModel
	Country            Country    `json:"country" valid:"-"`
	CountryId          uint64     `json:"country_id" valid:"required"`
	Gender             Gender     `json:"gender" valid:"required"`
	FirstName          string     `json:"first_name" valid:"required"`
	MiddleName         string     `json:"middle_name"`
	LastName           string     `json:"last_name" valid:"required"`
	NationalId         string     `json:"national_id" valid:"required"`
	CellNumber         string     `json:"cell_number" valid:"required"`
	PhoneNumber        string     `json:"phone_number"`
	PassportNumber     string     `json:"passport_number" valid:"required"`
	PassportIssueDate  string     `json:"passport_date_of_issue" valid:"required"`
	PassportExpireDate string     `json:"passport_expire_date" valid:"required"`
	Email              string     `json:"email" valid:"email"`
	DateOfBirth        *time.Time `json:"date_of_birth" valid:"required"`
	Address            string     `json:"address" valid:"required"`
}

// Validate validates guest model.
func (g *Guest) Validate() (bool, error) {

	return govalidator.ValidateStruct(g)
}

func (g *Guest) BeforeCreate(tx *gorm.DB) error {
	_, err := g.Validate()

	if err != nil {
		tx.AddError(err)
		return err
	}

	return nil
}
