package tenant_resolver

import "context"

func GetCurrentTenant(ctx context.Context) uint64 {
	return ctx.Value("TenantID").(uint64)
}
