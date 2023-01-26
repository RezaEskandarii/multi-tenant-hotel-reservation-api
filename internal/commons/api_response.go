package commons

import (
	"time"
)

// ApiResponse api response struct
type ApiResponse struct {
	Data interface{} `json:"data"`
	//	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message"`
	Errors       interface{} `json:"internal_errors"`
	ResponseCode int         `json:"response_code"`
}

func NewApiResponse() ApiResponse {
	return ApiResponse{}
}

func (a ApiResponse) SetData(data interface{}) ApiResponse {
	a.Data = data
	return a
}

func (a ApiResponse) SetResponseCode(statusCode int) ApiResponse {
	//a.ResponseCode = statusCode
	return a
}

func (a ApiResponse) SetMessage(message string) ApiResponse {
	a.Message = message
	return a
}

// =========================================================================

// JWTTokenResponse returns authentication token and expire time.
type JWTTokenResponse struct {
	ExpireAt    time.Time `json:"expire_at"`
	AccessToken string    `json:"access_token"`
}

// =========================================================================

type RequestTrace struct {
	TraceID      string      `json:"trace_id"`
	ResponseCode int         `json:"response_code,omitempty"`
	IPAddress    string      `json:"ip_address"`
	Tag          string      `json:"tag"`
	EndPoint     string      `json:"end_point"`
	Data         interface{} `json:"data"`
}

// =========================================================================
