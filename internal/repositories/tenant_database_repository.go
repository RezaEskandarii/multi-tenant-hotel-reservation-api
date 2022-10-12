package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/connection_resolver"
)

type TenantRepository struct {
	DB *gorm.DB
}

// NewTenantDatabaseRepository returns new TenantRepository.
func NewTenantDatabaseRepository(db *gorm.DB) *TenantRepository {

	return &TenantRepository{DB: db}
}

func (r *TenantRepository) Create(model *models.Tenant) (*models.Tenant, error) {

	if tx := r.DB.Create(&model); tx.Error != nil {
		return nil, tx.Error
	}

	resolver := connection_resolver.NewConnectionResolver()
	db1 := resolver.GetDB(0)
	tenantDB := resolver.GetDB(model.Id)

	resolver.CreateDbForTenant(db1, model.Id)
	resolver.Migrate(tenantDB, model.Id)

	return model, nil
}

func (r *TenantRepository) FindByTenantID(tenantID uint64) (*models.Tenant, error) {

	entity := models.Tenant{}
	if tx := r.DB.Where("tenant_id=?", tenantID).Find(&entity); tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
