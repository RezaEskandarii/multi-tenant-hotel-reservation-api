package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type ResidenceType struct {
	BaseModel
	Name   string            `json:"name" gorm:"type:varchar(100)" valid:"required"`
	Grades []*ResidenceGrade `json:"grades" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" valid:"-"`
}

func (r *ResidenceType) Validate() (bool, error) {

	return govalidator.ValidateStruct(r)
}

func (r *ResidenceType) BeforeCreate(tx *gorm.DB) error {
	if _, err := r.Validate(); err != nil {
		tx.AddError(err)
		return err
	}
	return nil
}
