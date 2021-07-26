package models

type ResidenceType struct {
	BaseModel
	Name   string            `json:"name" gorm:"type:varchar(100)" valid:"required"`
	Grades []*ResidenceGrade `json:"grades" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
