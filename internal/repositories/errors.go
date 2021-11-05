package repositories

import (
	"errors"
	"reservation-api/internal/message_keys"
)

var (
	TypeHasResidenceError = errors.New(message_keys.TypeHasResidence)
	GradeHasResidence     = errors.New(message_keys.GradeHasResidence)
)
