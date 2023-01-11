package internal_errors

import (
	"errors"
	"reservation-api/internal_errors/message_keys"
)

var (
	TypeHasHotelError = errors.New(message_keys.TypeHashotel)
	GradeHasHotel     = errors.New(message_keys.GradeHashotel)
	DuplicatedUser    = errors.New(message_keys.UsernameDuplicated)
)
