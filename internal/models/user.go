package models

import (
	"encoding/json"
	"fmt"
	"github.com/andskur/argon2-hashing"
	"gorm.io/gorm"
	"time"
)

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
	PhoneNumberConfirmed bool       `json:"-"`
	EmailConfirmed       bool       `json:"-"`
	Password             string     `json:"password,omitempty"`
	CityId               int        `json:"city_id"`
	City                 *City      `json:"city"`
	Gender               Gender     `json:"gender"`
	PassportNumber       string     `json:"passport_number"`
	NationalId           string     `json:"national_id"`
	Address              string     `json:"address"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	hash, err := argon2.GenerateFromPassword([]byte(u.Password), argon2.DefaultParams)
	if err != nil {
		tx.AddError(err)
		return err
	}

	u.Password = fmt.Sprintf("%s", hash)

	return nil
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User
	x := user(u)
	x.Password = ""
	return json.Marshal(x)
}
