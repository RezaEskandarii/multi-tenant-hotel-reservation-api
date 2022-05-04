package dto

import (
	"time"
)

type Gender string

var (
	Female Gender = "Female"
	Male   Gender = "Male"
	Other  Gender = "Other"
)

type GuestDto struct {
	BaseDto
	Country            *CountryDto `json:"country" valid:"-"`
	CountryId          uint64      `json:"country_id" valid:"required"`
	Gender             Gender      `json:"gender" valid:"required"`
	FirstName          string      `json:"first_name" valid:"required"`
	MiddleName         string      `json:"middle_name"`
	LastName           string      `json:"last_name" valid:"required" `
	NationalId         string      `json:"national_id" valid:"required" `
	CellNumber         string      `json:"cell_number" valid:"required" `
	PhoneNumber        string      `json:"phone_number" `
	PassportNumber     string      `json:"passport_number" `
	PassportIssueDate  string      `json:"passport_date_of_issue"`
	PassportExpireDate string      `json:"passport_expire_date"`
	Email              string      `json:"email" valid:"email" `
	DateOfBirth        *time.Time  `json:"date_of_birth"`
	Address            string      `json:"address" valid:"required"`
}
