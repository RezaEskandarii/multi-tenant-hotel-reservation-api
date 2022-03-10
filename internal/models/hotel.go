package models

import (
	"github.com/asaskevich/govalidator"
	"os"
)

type Hotel struct {
	BaseModel
	Name         string      `json:"name" gorm:"type:varchar(100)" valid:"required"`
	PhoneNumber1 string      `json:"phone_number1" gorm:"type:varchar(100)" valid:"required"`
	PhoneNumber2 string      `json:"phone_number2" gorm:"type:varchar(100)"`
	Province     *Province   `json:"province" valid:"-"`
	ProvinceId   uint64      `json:"province_id" gorm:"foreignKey:Province" valid:"required"`
	City         *City       `json:"city" valid:"-"`
	CityId       uint64      `json:"city_id" gorm:"foreignKey:City" valid:"required"`
	Address      string      `json:"address" valid:"required"`
	PostalCode   string      `json:"postal_code" gorm:"type:varchar(100)" valid:"required"`
	Longitude    float64     `json:"longitude" valid:"required"`
	Latitude     float64     `json:"latitude" valid:"required"`
	FaxNumber    string      `json:"fax_number" gorm:"type:varchar(100)"`
	Website      string      `json:"website" gorm:"type:varchar(100)"`
	EmailAddress string      `json:"email_address" gorm:"type:varchar(100)" valid:"email"`
	Owner        *User       `json:"owner"  valid:"-"`
	OwnerId      uint64      `json:"owner_id" gorm:"foreignKey:Owner" valid:"required"`
	Description  string      `json:"description"`
	HotelType    *HotelType  `json:"hotel_type"  valid:"-"`
	HotelTypeId  uint64      `json:"hotel_type_id" gorm:"foreignKey:HotelType" valid:"required"`
	HotelGrade   *HotelGrade `json:"hotel_grade"  valid:"-"`
	HotelGradeId uint64      `json:"hotel_grade_id" gorm:"foreignKey:HotelGrade" valid:"required"`
	Thumbnails   []*os.File  `json:"thumbnails" gorm:"-"`
}

func (h *Hotel) Validate() (bool, error) {

	return govalidator.ValidateStruct(h)
}

func (h *Hotel) SetAudit(username string) {
	h.CreatedBy = username
	h.UpdatedBy = username
}

func (h *Hotel) SetUpdatedBy(username string) {
	h.UpdatedBy = username
}
