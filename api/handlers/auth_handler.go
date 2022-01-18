package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"reservation-api/internal/dto"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/translator"
	"strings"
	"time"
)

type AuthHandler struct {
	Router     *echo.Group
	Service    *domain_services.UserService
	translator *translator.Translator
	logger     applogger.Logger
}

func (handler *AuthHandler) Register(input *dto.HandlerInput, service *domain_services.UserService) {

	handler.Router = input.Router
	handler.Service = service
	handler.translator = input.Translator
	handler.logger = input.Logger

	routeGroup := handler.Router.Group("/auth")
	routeGroup.POST("/signin", handler.signin)
	routeGroup.POST("/refresh-token", handler.refreshToken)
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
	user, userErr := handler.Service.FindByUsernameAndPassword(creds.Username, creds.Password)

	if user == nil || userErr != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(30 * time.Minute)
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

	tokenString, err := token.SignedString([]byte(getJwtKey()))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, TokenResponse{
		ExpireAt:    expirationTime,
		AccessToken: tokenString,
	})
}

// refresh token.
func (handler *AuthHandler) refreshToken(c echo.Context) error {

	tknStr := c.Request().Header.Get("Authorization")

	if tknStr == "" {
		return c.JSON(http.StatusBadRequest, "Authorization header is empty.")
	}

	authToken := strings.Split(tknStr, "Bearer ")

	tknStr = authToken[1]
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return getJwtKey(), nil
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if !tkn.Valid {
		return c.JSON(http.StatusUnauthorized, nil)
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return c.JSON(http.StatusBadRequest, nil)
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(30 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJwtKey())

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, TokenResponse{
		ExpireAt:    expirationTime,
		AccessToken: tokenString,
	})
}

func (handler *AuthHandler) logOut(c echo.Context) error {

	return nil
}

func getJwtKey() string {
	jwtKey, _ := os.LookupEnv("JWT_KEY")
	return jwtKey
}
