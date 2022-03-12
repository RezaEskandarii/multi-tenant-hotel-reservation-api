package commons

import "time"

// ApiResponse api response struct
type ApiResponse struct {
	Data         interface{} `json:"data"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message"`
}

// TokenResponse returns authentication token and expire time.
type TokenResponse struct {
	ExpireAt    time.Time `json:"expire_at"`
	AccessToken string    `json:"access_token"`
}
