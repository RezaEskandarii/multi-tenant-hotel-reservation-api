package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type ResidenceGrade struct {
	BaseModel
	Name            string         `json:"name" gorm:"type:varchar(100)" valid:"required"`
	ResidenceType   *ResidenceType `json:"residence_type" valid:"-"`
	ResidenceTypeId uint64         `json:"residence_type_id" gorm:"foreignKey:ResidenceType" valid:"required"`
}

func (r *ResidenceGrade) Validate() (bool, error) {
	ok, err := govalidator.ValidateStruct(r)

	if err != nil {
		return false, err
	}
	return ok, nil
}

func (r *ResidenceGrade) BeforeCreate(tx *gorm.DB) error {
	_, err := r.Validate()

	if err != nil {
		tx.AddError(err)
		return err
	}

	return nil
}
