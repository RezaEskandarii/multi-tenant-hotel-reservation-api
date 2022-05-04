package dto

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type UserDto struct {
	BaseDto
	FirstName            string     `json:"first_name" valid:"required"`
	LastName             string     `json:"last_name" valid:"required" `
	Username             string     `json:"username" valid:"required" `
	Email                string     `json:"email" valid:"required,email" `
	PhoneNumber          string     `json:"phone_number" valid:"required" `
	BirthDate            *time.Time `json:"birth_date"`
	PhoneNumberConfirmed bool       `json:"-"`
	EmailConfirmed       bool       `json:"-"`
	Password             string     `json:"password"`
	CityId               int        `json:"city_id"`
	City                 *CityDto   `json:"city" valid:"-"`
	Gender               Gender     `json:"gender" valid:"required"`
	Address              string     `json:"address"`
	IsActive             bool       `json:"is_active"`
}

func (u *UserDto) Validate() (bool, error) {

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
