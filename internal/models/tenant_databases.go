package models

type Tenant struct {
	BaseModel
	Name        string `json:"name"`
	Hash        string `json:"hash"`
	Description string `json:"description"`
}
