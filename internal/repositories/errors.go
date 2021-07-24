package repositories

import (
	"errors"
	"hotel-reservation/internal/message_keys"
)

var (
	TypeHasResidenceError = errors.New(message_keys.TypeHasResidence)
	GradeHasResidence     = errors.New(message_keys.GradeHasResidence)
)
