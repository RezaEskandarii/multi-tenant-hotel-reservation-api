package repositories

import (
	"errors"
	"reservation-api/internal/message_keys"
)

var (
	TypeHasHotelError = errors.New(message_keys.TypeHashotel)
	GradeHasHotel     = errors.New(message_keys.GradeHashotel)
)
