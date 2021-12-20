package repositories

import (
	"errors"
	"reservation-api/internal/message_keys"
)

var (
	TypeHashotelError = errors.New(message_keys.TypeHashotel)
	GradeHashotel     = errors.New(message_keys.GradeHashotel)
)
