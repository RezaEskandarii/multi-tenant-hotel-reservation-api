package models

import "time"

type Gender string

var (
	Female Gender = "Female"
	Male   Gender = "Male"
	Other  Gender = "Other"
)

type User struct {
	BaseModel
	FirstName            string     `json:"first_name"`
	LastName             string     `json:"last_name"`
	Username             string     `json:"username"`
	Email                string     `json:"email"`
	PhoneNumber          string     `json:"phone_number"`
	BirthDate            *time.Time `json:"birth_date"`
	PhoneNumberConfirmed bool       `json:"phone_number_confirmed"`
	EmailConfirmed       bool       `json:"email_confirmed"`
	Password             string     `json:"-"`
	CityId               int        `json:"city_id"`
	City                 *City      `json:"city"`
	Gender               Gender     `json:"gender"`
	PassportNumber       string     `json:"passport_number"`
	NationalId           string     `json:"national_id"`
	Address              string     `json:"address"`
}
