package tenant_resolver

import (
	"context"
	"reservation-api/internal/global_variables"
)

// GetCurrentTenant reads TenantID value from give context
// and returns as an uint64 type
func GetCurrentTenant(ctx context.Context) uint64 {
	return ctx.Value(global_variables.TenantIDKey).(uint64)
}
