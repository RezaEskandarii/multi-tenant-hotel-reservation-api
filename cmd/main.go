package main

import (
	"fmt"
	"os"
	"reservation-api/internal/kernel"
	"reservation-api/internal/utils"
	"reservation-api/pkg/applogger"
)

type r struct {
	Name         string `json:"name" gorm:"type:varchar(100)" valid:"required"`
	PhoneNumber1 string `json:"phone_number1" gorm:"type:varchar(100)" valid:"required"`
	PhoneNumber2 string `json:"phone_number2" gorm:"type:varchar(100)"`

	ProvinceId uint64 `json:"province_id" gorm:"foreignKey:Province" valid:"required"`

	CityId       uint64  `json:"city_id" gorm:"foreignKey:City" valid:"required"`
	Address      string  `json:"address" valid:"required"`
	PostalCode   string  `json:"postal_code" gorm:"type:varchar(100)" valid:"required"`
	Longitude    float64 `json:"longitude" valid:"required"`
	Latitude     float64 `json:"latitude" valid:"required"`
	FaxNumber    string  `json:"fax_number" gorm:"type:varchar(100)"`
	Website      string  `json:"website" gorm:"type:varchar(100)"`
	EmailAddress string  `json:"email_address" gorm:"type:varchar(100)" valid:"email"`

	OwnerId     uint64 `json:"owner_id" gorm:"foreignKey:Owner" valid:"required"`
	Description string `json:"description"`

	HotelTypeId uint64 `json:"hotel_type_id" gorm:"foreignKey:HotelType" valid:"required"`

	HotelGradeId uint64 `json:"hotel_grade_id" gorm:"foreignKey:HotelGrade" valid:"required"`
}

func main() {

	logger := applogger.New(nil)

	fmt.Println("======================================")
	fmt.Println(string(utils.ToJson(r{})))
	if err := kernel.Run(); err != nil {
		logger.LogError(err)
		os.Exit(1)
	}
}
