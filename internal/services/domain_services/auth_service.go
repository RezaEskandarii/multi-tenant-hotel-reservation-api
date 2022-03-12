package domain_services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"reservation-api/internal/commons"
	"reservation-api/internal/config"
	"time"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthService struct {
	UserService *UserService
	Config      *config.Config
}

func NewAuthService(service *UserService, cfg *config.Config) *AuthService {
	return &AuthService{
		UserService: service,
		Config:      cfg,
	}
}

func (s AuthService) SignIn(username, password string) (error, *commons.TokenResponse) {

	// Get the expected password from our in memory map
	user, err := s.UserService.FindByUsernameAndPassword(username, password)

	if user == nil || err != nil {
		return err, nil
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(time.Duration(s.Config.Authentication.TokenAliveTime) * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := s.Config.Authentication.JwtKey

	tokenString, err := token.SignedString([]byte(jwtKey))
	return nil, &commons.TokenResponse{
		ExpireAt:    expirationTime,
		AccessToken: tokenString,
	}
}

func (s *AuthService) RefreshToken(tokenStr string) (error, *commons.TokenResponse) {

	jwtKey := s.Config.Authentication.JwtKey
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return err, nil
	}

	if !tkn.Valid {
		return errors.New("invalid Token"), nil
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > time.Duration(s.Config.Authentication.TokenAliveTime)*time.Second {
		return errors.New("invalid Token"), nil
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(time.Duration(s.Config.Authentication.TokenAliveTime) * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(jwtKey)

	if err != nil {
		return err, nil
	}

	return nil, &commons.TokenResponse{
		ExpireAt:    expirationTime,
		AccessToken: result,
	}
}
