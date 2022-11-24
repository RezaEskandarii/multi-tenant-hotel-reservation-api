package repositories

import (
	"context"
	"fmt"
	"github.com/andskur/argon2-hashing"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/internal/utils"
	"reservation-api/pkg/database/tenant_database_resolver"
)

type UserRepository struct {
	ConnectionResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewUserRepository(r *tenant_database_resolver.TenantDatabaseResolver) *UserRepository {
	return &UserRepository{
		ConnectionResolver: r,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	user.TenantId = tenant_resolver.GetCurrentTenant(ctx)
	if tx := db.Create(&user); tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) (*models.User, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	user.TenantId = tenant_resolver.GetCurrentTenant(ctx)

	if tx := db.Updates(&user); tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func (r *UserRepository) Find(ctx context.Context, id uint64) (*models.User, error) {

	model := models.User{}
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	model := models.User{Username: username}

	if tx := db.Where(model).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}
func (r *UserRepository) FindByUsernameAndPassword(ctx context.Context, username string, password string) (*models.User, error) {

	model := models.User{}
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where("username=?", username).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	err := argon2.CompareHashAndPassword([]byte(model.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *UserRepository) Deactivate(ctx context.Context, id uint64) (*models.User, error) {

	user := models.User{}
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	query := db.Model(&models.User{}).Where("id=?", id).Find(&user)

	if query.Error != nil {
		return nil, query.Error
	}

	user.IsActive = false

	if tx := db.Model(&models.User{}).Updates(user); tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r *UserRepository) Activate(ctx context.Context, id uint64) (*models.User, error) {

	user := models.User{}
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	query := db.Model(&models.User{}).Where("id=?", id).Find(&user)

	if query.Error != nil {
		return nil, query.Error
	}

	user.IsActive = true

	if tx := db.Model(&models.User{}).Updates(user); tx.Error != nil {

		return nil, tx.Error
	}

	return &user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Model(&models.User{}).Where("id=?", id).Delete(&models.User{}); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *UserRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	return paginatedList(&models.User{}, db, input)
}

func (r *UserRepository) Seed(ctx context.Context, jsonFilePath string) error {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	users := make([]models.User, 0)
	if err := utils.CastJsonFileToStruct(jsonFilePath, &users); err == nil {
		for _, user := range users {
			var count int64 = 0
			if err := db.Model(models.User{}).Where("username", user.Username).Count(&count).Error; err != nil {
				return err
			} else {
				if count == 0 {
					user.TenantId = tenant_resolver.GetCurrentTenant(ctx)

					hash, err := argon2.GenerateFromPassword([]byte(user.Password), argon2.DefaultParams)

					if err != nil {
						return err
					}
					user.Password = fmt.Sprintf("%s", hash)

					if err := db.Create(&user).Error; err != nil {
						return err
					}
				}
			}
		}
	} else {
		return err
	}
	return nil
}
