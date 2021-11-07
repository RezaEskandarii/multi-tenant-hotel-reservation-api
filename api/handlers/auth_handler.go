package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/dto"
	"reservation-api/internal/services"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/translator"
	"time"
)

type AuthHandler struct {
	Router     *echo.Group
	Service    *services.UserService
	translator *translator.Translator
	logger     applogger.Logger
}

func (handler *AuthHandler) Register(input *dto.HandlerInput, service *services.UserService) {

	handler.Router = input.Router
	handler.Service = service
	handler.translator = input.Translator
	handler.logger = input.Logger

	routeGroup := handler.Router.Group("/auth")
	routeGroup.POST("/signin", handler.signin)
	routeGroup.POST("", handler.refreshToken)
	routeGroup.POST("", handler.logOut)
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type TokenResponse struct {
	ExpireAt    time.Time `json:"expire_at"`
	AccessToken string    `json:"access_token"`
}

func (handler *AuthHandler) signin(c echo.Context) error {

	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(c.Request().Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		return c.JSON(http.StatusBadRequest, nil)
	}

	// Get the expected password from our in memory map
	user, err := handler.Service.FindByUsernameAndPassword(creds.Username, creds.Password)

	if user == nil && err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString("jwtKey")
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return c.JSON(http.StatusInternalServerError, nil)
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself

	return c.JSON(http.StatusOK, TokenResponse{
		ExpireAt:    expirationTime,
		AccessToken: tokenString,
	})
}

func (handler *AuthHandler) refreshToken(c echo.Context) error {

	return nil
}

func (handler *AuthHandler) logOut(c echo.Context) error {

	return nil
}
