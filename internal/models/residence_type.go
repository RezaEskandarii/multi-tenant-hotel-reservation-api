package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type ResidenceType struct {
	BaseModel
	Name   string            `json:"name" gorm:"type:varchar(100)" valid:"required"`
	Grades []*ResidenceGrade `json:"grades" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (r *ResidenceType) Validate() (bool, error) {

	ok, err := govalidator.ValidateStruct(r)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (r *ResidenceType) BeforeCreate(tx *gorm.DB) error {

	_, err := r.Validate()

	if err != nil {
		tx.AddError(err)
		return err
	}

	return nil
}
