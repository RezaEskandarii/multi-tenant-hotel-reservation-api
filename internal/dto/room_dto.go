package dto

import (
	"errors"
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
	BaseDto
	Name        string      `json:"name" valid:"required" `
	RoomType    RoomTypeDto `json:"room_type" valid:"-"`
	RoomTypeId  uint64      `json:"room_type_id" valid:"required"`
	MaxBeds     uint64      `json:"max_beds" valid:"required"`
	CleanStatus CleanStatus `json:"clean_status" valid:"required"`
	Description string      `json:"description" valid:"maxstringlength(255)"`
}

type RoomTypeDto struct {
}
