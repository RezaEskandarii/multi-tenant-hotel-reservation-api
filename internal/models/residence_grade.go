package models

type ResidenceGrade struct {
	BaseModel
	Name            string         `json:"name" gorm:"type:varchar(100)"`
	ResidenceType   *ResidenceType `json:"residence_type"`
	ResidenceTypeId uint64         `json:"residence_type_id" gorm:"foreignKey:ResidenceType"`
}
