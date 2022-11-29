package domain_services

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reservation-api/internal/commons"
	"reservation-api/internal/global_variables"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/internal/utils"
	"time"
)

type Credentials struct {
	Password string `json:"password"  valid:"required"`
	Username string `json:"username"  valid:"required"`
}

type Claims struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	TenantID    string `json:"tenant_id"`
	jwt.StandardClaims
}

type AuthService struct {
	UserService *UserService
	Config      *global_variables.Config
}

var (
	InvalidTokenError = errors.New("token is invalid")
)

// NewAuthService returns new instance of auth service.
func NewAuthService(service *UserService, cfg *global_variables.Config) *AuthService {
	return &AuthService{
		UserService: service,
		Config:      cfg,
	}
}

// SignIn finds user with given username and password and
// generates JWT token for user.
func (s AuthService) SignIn(ctx context.Context, username, password string) (error, *commons.JWTTokenResponse) {

	// Get the expected password from our in memory map
	user, err := s.UserService.FindByUsernameAndPassword(ctx, username, password)

	if user == nil || err != nil {
		return err, nil
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(time.Duration(s.Config.Authentication.TokenAliveTime) * time.Minute)
	// SetUp the JWT claims, which includes the username and expiry time
	claims := &Claims{
		TenantID:    utils.Encrypt(fmt.Sprintf("%d", tenant_resolver.GetCurrentTenant(ctx))),
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := s.Config.Authentication.JwtKey

	tokenString, err := token.SignedString([]byte(jwtKey))
	return nil, &commons.JWTTokenResponse{
		ExpireAt:    expirationTime,
		AccessToken: tokenString,
	}
}

func (s *AuthService) RefreshToken(tokenStr string) (error, *commons.JWTTokenResponse) {

	jwtKey := s.Config.Authentication.JwtKey
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return err, nil
	}

	if !tkn.Valid {
		return InvalidTokenError, nil
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > time.Duration(s.Config.Authentication.TokenAliveTime)*time.Second {
		return InvalidTokenError, nil
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(time.Duration(s.Config.Authentication.TokenAliveTime) * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(jwtKey)

	if err != nil {
		return err, nil
	}

	return nil, &commons.JWTTokenResponse{
		ExpireAt:    expirationTime,
		AccessToken: result,
	}
}

func (s *AuthService) VerifyToken(ctx context.Context, jwtToken string, tenantID uint64) (error, *Claims) {

	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("uexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.Config.Authentication.JwtKey), nil
	})

	if token == nil {
		return errors.New("token is null"), nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		username := claims["username"]
		err, user := s.UserService.FindByUsername(ctx, fmt.Sprintf("%s", username))

		if err != nil {
			return err, nil
		}

		if user == nil {
			return InvalidTokenError, nil
		}

		return nil, &Claims{
			Username:    fmt.Sprintf("%s", claims["username"]),
			Email:       fmt.Sprintf("%s", claims["email"]),
			FirstName:   fmt.Sprintf("%s", claims["firstname"]),
			LastName:    fmt.Sprintf("%s", claims["lastname"]),
			Address:     fmt.Sprintf("%s", claims["address"]),
			PhoneNumber: fmt.Sprintf("%s", claims["phonenumber"]),
			TenantID:    fmt.Sprintf("%d", tenantID), //utils.Encrypt(fmt.Sprintf("%d", tenantID)),
		}

	} else {
		return InvalidTokenError, nil
	}

}
