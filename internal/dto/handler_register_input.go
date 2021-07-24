package dto

import (
	"github.com/labstack/echo/v4"
	"hotel-reservation/pkg/translator"
)

type HandlerRegisterInput struct {
	Router     *echo.Group
	Translator *translator.Translator
}
