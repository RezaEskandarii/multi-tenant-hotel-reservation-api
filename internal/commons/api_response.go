package commons

// ApiResponse api response struct
type ApiResponse struct {
	Data         interface{} `json:"resources"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message"`
}
