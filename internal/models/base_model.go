package models

import (
	"time"
)

// BaseModel base model to other models
type BaseModel struct {
	Id        uint64     `json:"id" gorm:"primarykey"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	TenantId  uint64     `json:"tenant_id"`
	CreatedBy uint64     `json:"created_by"`
	UpdatedBy uint64     `json:"updated_by"`
}
