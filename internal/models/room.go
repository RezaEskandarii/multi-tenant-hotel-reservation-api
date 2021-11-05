package models

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"reservation-api/internal/message_keys"
)

type CleanStatus int

var (
	Clean      CleanStatus = 1
	Dirty      CleanStatus = 2
	InProgress CleanStatus = 3

	InvalidCleanStatusErr = errors.New(message_keys.InvalidRoomCleanStatus)
)

type Room struct {
	BaseModel
	Name        string      `json:"name" valid:"required"`
	RoomType    RoomType    `json:"room_type" valid:"-"`
	RoomTypeId  uint64      `json:"room_type_id" valid:"required"`
	MaxBeds     uint64      `json:"max_beds" valid:"required"`
	CleanStatus CleanStatus `json:"clean_status" valid:"required"`
	Description string      `json:"description" valid:"maxstringlength(255)"`
}

func (r *Room) Validate() (bool, error) {

	ok, err := govalidator.ValidateStruct(r)

	if err != nil {
		return false, err
	}

	hasValidCleanStatus := false

	if r.CleanStatus == Dirty || r.CleanStatus == InProgress || r.CleanStatus == Clean {
		hasValidCleanStatus = true
	}

	if hasValidCleanStatus == false {

		return false, InvalidCleanStatusErr
	}

	return ok, err
}
