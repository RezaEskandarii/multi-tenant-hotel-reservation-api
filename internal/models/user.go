package models

import (
	"encoding/json"
	"errors"
	"github.com/asaskevich/govalidator"
	"reservation-api/internal/message_keys"
	"time"
)

type Gender string

var (
	Female Gender = "Female"
	Male   Gender = "Male"
	Other  Gender = "Other"

	InvalidGenderError = errors.New(message_keys.GenderInvalid)
)

type User struct {
	BaseModel
	FirstName            string     `json:"first_name" valid:"required"  gorm:"type:varchar(255)"`
	LastName             string     `json:"last_name" valid:"required"  gorm:"type:varchar(255)"`
	Username             string     `json:"username" valid:"required"     gorm:"type:varchar(255)"`
	Email                string     `json:"email" valid:"required,email"  gorm:"type:varchar(255)"`
	PhoneNumber          string     `json:"phone_number" valid:"required"  gorm:"type:varchar(255)"`
	BirthDate            *time.Time `json:"birth_date"`
	PhoneNumberConfirmed bool       `json:"-"`
	EmailConfirmed       bool       `json:"-"`
	Password             string     `json:"password"  gorm:"type:varchar(255)"`
	CityId               int        `json:"city_id"`
	City                 *City      `json:"city" valid:"-"`
	Gender               Gender     `json:"gender" valid:"required"`
	Address              string     `json:"address"`
	IsActive             bool       `json:"is_active"`
}

type GetUser struct {
	BaseDto
	FirstName            string     `json:"first_name" valid:"required"`
	LastName             string     `json:"last_name" valid:"required" `
	Username             string     `json:"username" valid:"required" `
	Email                string     `json:"email" valid:"required,email" `
	PhoneNumber          string     `json:"phone_number" valid:"required" `
	BirthDate            *time.Time `json:"birth_date"`
	PhoneNumberConfirmed bool       `json:"-"`
	EmailConfirmed       bool       `json:"-"`
	CityId               int        `json:"city_id"`
	City                 *GetCity   `json:"city" valid:"-"`
	Gender               Gender     `json:"gender" valid:"required"`
	Address              string     `json:"address"`
	IsActive             bool       `json:"is_active"`
}

type UserCreateUpdate struct {
	BaseDto
	FirstName            string      `json:"first_name" valid:"required"`
	LastName             string      `json:"last_name" valid:"required" `
	Username             string      `json:"username" valid:"required" `
	Email                string      `json:"email" valid:"required,email" `
	PhoneNumber          string      `json:"phone_number" valid:"required" `
	BirthDate            *time.Time  `json:"birth_date"`
	PhoneNumberConfirmed bool        `json:"-"`
	EmailConfirmed       bool        `json:"-"`
	Password             string      `json:"password"`
	CityId               int         `json:"city_id"`
	City                 *GetCountry `json:"city" valid:"-"`
	Gender               Gender      `json:"gender" valid:"required"`
	Address              string      `json:"address"`
	IsActive             bool        `json:"is_active"`
}

func (u *UserCreateUpdate) Validate() (bool, error) {

	ok, err := govalidator.ValidateStruct(u)
	if err != nil {
		return false, err
	}

	hasWringGender := true
	if u.Gender == Male || u.Gender == Female || u.Gender == Other {
		hasWringGender = false
	}

	if hasWringGender {
		return false, InvalidCleanStatusErr
	}

	return ok, nil
}

//MarshalJSON prevents to show user's password in json serializations.
func (u User) MarshalJSON() ([]byte, error) {
	type user User
	x := user(u)
	x.Password = ""
	return json.Marshal(x)
}

func (u *User) SetAudit(username string) {
	u.CreatedBy = username
	u.UpdatedBy = username
}

func (u *User) SetUpdatedBy(username string) {
	u.UpdatedBy = username
}
