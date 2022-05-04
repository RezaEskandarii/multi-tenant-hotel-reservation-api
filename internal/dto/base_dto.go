package dto

import "time"

type BaseDto struct {
	Id        uint64     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	TenantId  uint64     `json:"tenant_id"`
	CreatedBy string     `json:"created_by"`
	UpdatedBy string     `json:"updated_by"`
}
