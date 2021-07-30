package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Residence struct {
	BaseModel
	Name             string         `json:"name" gorm:"type:varchar(100)" valid:"required"`
	PhoneNumber1     string         `json:"phone_number1" gorm:"type:varchar(100)" valid:"required"`
	PhoneNumber2     string         `json:"phone_number2" gorm:"type:varchar(100)"`
	Province         Province       `json:"province" valid:"-"`
	ProvinceId       uint64         `json:"province_id" gorm:"foreignKey:Province" valid:"required"`
	City             City           `json:"city" valid:"-"`
	CityId           uint64         `json:"city_id" gorm:"foreignKey:City" valid:"required"`
	Address          string         `json:"address" valid:"required"`
	PostalCode       string         `json:"postal_code" gorm:"type:varchar(100)" valid:"required"`
	Longitude        float64        `json:"longitude" valid:"required"`
	Latitude         float64        `json:"latitude" valid:"required"`
	FaxNumber        string         `json:"fax_number" gorm:"type:varchar(100)"`
	Website          string         `json:"website" gorm:"type:varchar(100)"`
	EmailAddress     string         `json:"email_address" gorm:"type:varchar(100)" valid:"email"`
	Owner            User           `json:"owner"  valid:"-"`
	OwnerId          uint64         `json:"owner_id" gorm:"foreignKey:Owner" valid:"required"`
	Description      string         `json:"description"`
	ResidenceType    ResidenceType  `json:"residence_type"  valid:"-"`
	ResidenceTypeId  uint64         `json:"residence_type_id" gorm:"foreignKey:ResidenceType" valid:"required"`
	ResidenceGrade   ResidenceGrade `json:"residence_grade"  valid:"-"`
	ResidenceGradeId uint64         `json:"residence_grade_id" gorm:"foreignKey:ResidenceGrade" valid:"required"`
}

func (r *Residence) Validate() (bool, error) {
	ok, err := govalidator.ValidateStruct(r)

	if err != nil {
		return false, err
	}

	return ok, nil
}

func (r *Residence) BeforeCreate(tx *gorm.DB) error {
	_, err := r.Validate()

	if err != nil {
		tx.AddError(err)
		return err
	}

	return nil
}
