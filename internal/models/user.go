package models

import (
	"errors"
	"fmt"
	"github.com/andskur/argon2-hashing"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
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
	FirstName            string     `json:"first_name" valid:"required"`
	LastName             string     `json:"last_name" valid:"required"`
	Username             string     `json:"username" valid:"required"`
	Email                string     `json:"email" valid:"required,email"`
	PhoneNumber          string     `json:"phone_number" valid:"required"`
	BirthDate            *time.Time `json:"birth_date"`
	PhoneNumberConfirmed bool       `json:"-"`
	EmailConfirmed       bool       `json:"-"`
	Password             string     `json:"password"`
	CityId               int        `json:"city_id"`
	City                 *City      `json:"city" valid:"-"`
	Gender               Gender     `json:"gender" valid:"required"`
	PassportNumber       string     `json:"passport_number"`
	NationalId           string     `json:"national_id" valid:"required"`
	Address              string     `json:"address"`
	IsActive             bool       `json:"is_active"`
}

func (u *User) Validate() (bool, error) {

	ok, err := govalidator.ValidateStruct(u)
	if err != nil {
		return false, err
	}

	return ok, nil
}

// BeforeCreate validates user struct and change user's password to bcrypt.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := u.Validate()

	if err != nil {
		tx.AddError(err)
		return err
	}
	hasWringGender := true
	if u.Gender == Male || u.Gender == Female || u.Gender == Other {
		hasWringGender = false
	}

	if hasWringGender {
		tx.AddError(InvalidGenderError)
		return InvalidGenderError
	}

	hash, err := argon2.GenerateFromPassword([]byte(u.Password), argon2.DefaultParams)

	if err != nil {
		tx.AddError(err)
		return err
	}

	u.Password = fmt.Sprintf("%s", hash)
	return nil
}

// MarshalJSON prevents to show user's password in json serializations.
//func (u User) MarshalJSON() ([]byte, error) {
//	type user User
//	x := user(u)
//	x.Password = ""
//	return json.Marshal(x)
//}
