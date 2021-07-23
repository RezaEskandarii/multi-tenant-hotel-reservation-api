package models

type Residence struct {
	BaseModel
	Name             string         `json:"name" gorm:"type:varchar(100)"`
	PhoneNumber1     string         `json:"phone_number1" gorm:"type:varchar(100)"`
	PhoneNumber2     string         `json:"phone_number2" gorm:"type:varchar(100)"`
	Province         Province       `json:"province"`
	ProvinceId       uint64         `json:"province_id" gorm:"foreignKey:Province"`
	Address          string         `json:"address"`
	PostalCode       string         `json:"postal_code" gorm:"type:varchar(100)"`
	Longitude        float64        `json:"longitude"`
	Latitude         float64        `json:"latitude"`
	FaxNumber        string         `json:"fax_number" gorm:"type:varchar(100)"`
	Website          string         `json:"website" gorm:"type:varchar(100)"`
	EmailAddress     string         `json:"email_address" gorm:"type:varchar(100)"`
	Owner            User           `json:"owner"`
	OwnerId          uint64         `json:"owner_id" gorm:"foreignKey:Owner"`
	Description      string         `json:"description"`
	ResidenceType    ResidenceType  `json:"residence_type"`
	ResidenceTypeId  uint64         `json:"residence_type_id" gorm:"foreignKey:ResidenceType"`
	ResidenceGrade   ResidenceGrade `json:"residence_grade"`
	ResidenceGradeId uint64         `json:"residence_grade_id" gorm:"foreignKey:ResidenceGrade"`
}
