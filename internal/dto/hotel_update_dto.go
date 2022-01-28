package dto

import (
	"os"
	"reservation-api/internal/models"
)

type HotelUpdateDto struct {
	Name         string  `json:"name" valid:"required" gorm:"type:varchar(255)"`
	PhoneNumber1 string  `json:"phone_number1" valid:"required"  gorm:"type:varchar(255)"`
	PhoneNumber2 string  `json:"phone_number2"  gorm:"type:varchar(255)"`
	ProvinceId   uint64  `json:"province_id" valid:"required"`
	CityId       uint64  `json:"city_id" valid:"required"`
	Address      string  `json:"address" valid:"required"`
	PostalCode   string  `json:"postal_code" valid:"required"  gorm:"type:varchar(255)"`
	Longitude    float64 `json:"longitude" valid:"required"`
	Latitude     float64 `json:"latitude" valid:"required"`
	FaxNumber    string  `json:"fax_number"  gorm:"type:varchar(255)"`
	Website      string  `json:"website"  gorm:"type:varchar(255)"`
	EmailAddress string  `json:"email_address" valid:"email"  gorm:"type:varchar(255)"`
	Description  string  `json:"description"  gorm:"type:varchar(255)"`
	HotelTypeId  uint64  `json:"hotel_type_id" valid:"required"`
	HotelGradeId uint64  `json:"hotel_grade_id" valid:"required"`
}

type HotelCreateDto struct {
	Data       models.Hotel `json:"data"`
	Thumbnails os.File      `json:"thumbnails"`
}
