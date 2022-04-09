package commons

import "time"

// ApiResponse api response struct
type ApiResponse struct {
	Data         interface{} `json:"data"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message"`
}

// JWTTokenResponse returns authentication token and expire time.
type JWTTokenResponse struct {
	ExpireAt    time.Time `json:"expire_at"`
	AccessToken string    `json:"access_token"`
}
