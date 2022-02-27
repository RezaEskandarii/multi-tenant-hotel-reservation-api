package models

import (
	"encoding/json"
	"time"
)

// BaseModel base model to other models
type BaseModel struct {
	Id        uint64     `json:"id" gorm:"primarykey"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	TenantId  uint64     `json:"tenant_id"`
	CreatedBy string     `json:"created_by"`
	UpdatedBy string     `json:"updated_by"`
}

func (b BaseModel) ToJson() []byte {
	result, err := json.Marshal(&b)

	if err != nil {
		return result
	}
}
