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

	return govalidator.ValidateStruct(r)
}

func (r *ResidenceGrade) BeforeCreate(tx *gorm.DB) error {
	if _, err := r.Validate(); err != nil {
		tx.AddError(err)
		return err
	}
	return nil
}
