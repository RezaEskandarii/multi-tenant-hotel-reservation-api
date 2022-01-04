package models

type Audit struct {
	BaseModel
	UserId     uint64 `json:"user_id"`
	User       User   `json:"user"`
	Username   string `json:"username"`
	HttpMethod string `json:"http_method"`
	Url        string `json:"url"`
	Data       string `json:"data"`
}
